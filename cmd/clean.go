/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"github.com/niccoloCastelli/orderbooks/config"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/spf13/cobra"
)

const dateLayout = "2006/01/02"

// cleanCmd represents the clean command
var cleanCmd = &cobra.Command{
	Use:   "clean",
	Short: "Pulizia storage",
	Long:  `Pulizia storage eccetto ultimi n giorni (default 5)`,
	Run: func(cmd *cobra.Command, args []string) {
		var err error
		logger := log.With().Str("cmd", "clean").Logger()
		keepDays := 5
		storagePath := viper.GetString(config.StoragePath)
		storageAbsPath, _ := filepath.Abs(storagePath)
		if len(args) > 0 {
			if keepDays, err = strconv.Atoi(args[0]); err != nil {
				logger.Fatal().Err(err).Msg("clean error")
			}
		}
		logger.Info().Int("keep_days", keepDays).Str("storage_path", storagePath).Send()
		today, _ := time.Parse(dateLayout, time.Now().Format(dateLayout))
		keepAfter := today.Add(-time.Hour * 24 * time.Duration(keepDays))
		logger.Info().Str("today", today.String()).Str("remove_before", keepAfter.String()).Send()
		toRemove := []string{}
		err = filepath.Walk(storagePath, func(path string, info os.FileInfo, err error) error {
			logger := logger.With().Str("path", path).Logger()
			if !info.IsDir() {
				logger.Trace().Msg("skip file")
				return nil
			}
			if err != nil {
				logger.Err(err).Send()
				return nil
			}
			absPath, _ := filepath.Abs(path)
			relPath := strings.TrimPrefix(strings.TrimPrefix(absPath, storageAbsPath), "/")
			logger.Trace().Str("rel_path", relPath).Msg("rel path")
			splitPath := strings.Split(relPath, "/")
			if len(splitPath) != 5 {
				logger.Trace().Int("length", len(splitPath)).Msg("wrong path length, skip")
				return nil
			}

			t, err := time.Parse(dateLayout, strings.Join(splitPath[2:len(splitPath)], "/"))
			if err != nil {
				logger.Debug().Err(err).Str("date", strings.Join(splitPath[2:len(splitPath)], "/")).Str("rel_path", relPath).Send()
				return nil
			}
			if t.Before(keepAfter) {
				logger.Debug().Str("abs_path", absPath).Str("date", t.String()).Msg("remove path")
				toRemove = append(toRemove, absPath)
				return nil
			} else {
				logger.Debug().Str("date", t.String()).Msg("keep path")
			}

			return nil
		})
		if err != nil {
			logger.Fatal().Err(err).Msg("filepath walk error")
			return
		}
		for _, fp := range toRemove {
			if err := os.RemoveAll(fp); err != nil {
				logger.Err(err).Send()
			}
			logger.Warn().Str("abs_path", fp).Msg("path removed")

		}
		logger.Info().Msg("clean ok")

	},
}

func init() {
	rootCmd.AddCommand(cleanCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// cleanCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// cleanCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
