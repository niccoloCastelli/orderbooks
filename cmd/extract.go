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
	"github.com/niccoloCastelli/orderbooks/common"
	"github.com/niccoloCastelli/orderbooks/common/constants"
	"github.com/niccoloCastelli/orderbooks/config"
	"github.com/niccoloCastelli/orderbooks/exchanges"
	"github.com/niccoloCastelli/orderbooks/utils"
	"github.com/rs/zerolog/log"
	"github.com/spf13/afero"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
	"path/filepath"
	"strings"
	"time"
)

var (
	extractEvents bool
)

// extractCmd represents the extract command
var extractCmd = &cobra.Command{
	Use:   "extract",
	Short: "Extract compressed order books",
	Long:  `Extract order books compressed in gzip format. orderbooks extract [EXCHANGE] [PAIR]`,
	Run: func(cmd *cobra.Command, args []string) {
		var err error
		logger := log.With().Str("cmd", "extract").Logger()
		storagePath := viper.GetString(config.StoragePath)
		if len(args) < 2 {
			logger.Error().Msg("usage: orderbooks extract [EXCHANGE] [PAIR]")
			return
		}
		exchangeStr := args[0]
		pairStr := args[1]
		outPath := ""
		if len(args) > 2 {
			outPath = args[2]
		}

		logger = logger.With().Str("pair", pairStr).Str("exchange", exchangeStr).Logger()
		logger.Info().Send()
		exchange, ok := exchanges.GetExchange(exchangeStr)
		if !ok {
			logger.Error().Msg("exchange not found")
			return
		}
		if !ok {
			logger.Error().Msg("exchange not found")
			return
		}
		pair, err := common.PairsFromString(pairStr)
		if err != nil {
			logger.Error().Msg("invalid pair")
			return
		}
		if !exchange.PairAvailable(*pair) {
			logger.Error().Msg("pair not found")
			return
		}
		osFs := afero.NewOsFs()
		fs := afero.NewBasePathFs(osFs, storagePath)
		fileType := constants.FileTypeSnapshots
		if extractEvents {
			fileType = constants.FileTypeEvents
		}
		err = afero.Walk(fs, "/"+string(fileType), func(filePath string, info os.FileInfo, err error) error {
			logger := logger.With().Str("filePath", filePath).Logger()
			if info.IsDir() {
				logger.Trace().Msg("skip dir")
				return nil
			}
			if err != nil {
				logger.Err(err).Send()
				return nil
			}
			logger.Trace().Str("rel_path", filePath).Msg("rel filePath")
			splitPath := strings.Split(filePath, "/")
			if len(splitPath) < 6 {
				logger.Trace().Int("length", len(splitPath)).Msg("wrong filePath length, skip")
				return nil
			}
			dateIdxStart := len(splitPath) - 4
			dateIdxEnd := len(splitPath) - 1
			fileNameIdx := dateIdxEnd
			exchangeIdx := len(splitPath) - 5
			t, err := time.Parse(dateLayout, strings.Join(splitPath[dateIdxStart:dateIdxEnd], "/"))
			if err != nil {
				logger.Debug().Err(err).Str("date", strings.Join(splitPath[dateIdxStart:len(splitPath)], "/")).Str("rel_path", filePath).Send()
				return nil
			}
			if splitPath[exchangeIdx] != exchange.Name() {
				logger.Trace().Int("length", len(splitPath)).Msg("skip exchange")
				return nil
			}
			if ext := filepath.Ext(splitPath[fileNameIdx]); ext != ".gz" {
				logger.Debug().Str("filePath", filePath).Msg("skip file")
				return nil
			}
			if strings.TrimSuffix(splitPath[fileNameIdx], ".csv.gz") != pair.String() {
				logger.Debug().Str("filePath", filePath).Msg("skip pair")
				return nil
			}
			logger.Debug().Str("filePath", filePath).Str("out", outPath).Msg("exctract file")
			outFile, err := utils.ExtractFile(fs, filePath, osFs, outPath, exchange.Name(), fileType, t)
			if err != nil {
				return err
			}
			logger.Info().Str("filePath", filePath).Str("out", outFile).Msg("file extracted")

			return nil
		})
		if err != nil {
			logger.Fatal().Err(err).Msg("filepath walk error")
			return
		}
		logger.Info().Msg("extract ok")

	},
}

func init() {
	rootCmd.AddCommand(extractCmd)

	extractCmd.PersistentFlags().BoolVar(&extractEvents, "extract_events", false, "-1(trace)... 5(panic)")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// extractCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// extractCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
