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
	"context"
	"github.com/niccoloCastelli/orderbooks/common"
	"github.com/niccoloCastelli/orderbooks/config"
	"github.com/niccoloCastelli/orderbooks/utils"
	"github.com/niccoloCastelli/orderbooks/workers/backends/csv_backend"
	"github.com/niccoloCastelli/orderbooks/workers/downloaders"
	"github.com/pkg/errors"
	"github.com/spf13/afero"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"log"
	"os"
	"time"
)

// runCmd represents the run command
var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Avvia il download degli order books",
	Long:  `Avvia il download degli order books`,
	Run: func(cmd *cobra.Command, args []string) {
		ctx, cancelCtx := context.WithCancel(context.Background())
		snapshotChan := make(chan common.Snapshot, 1024)
		storagePath := viper.GetString(config.StoragePath)
		saveEvents := viper.GetBool(config.SaveEvents)
		if err := os.MkdirAll(storagePath, 0755); err != nil {
			log.Fatal(err)
		}

		baseFs := afero.NewBasePathFs(afero.NewOsFs(), storagePath)
		sd := downloaders.NewSnapshotDownloader(pairs, time.Second*time.Duration(viper.GetInt(config.SnapshotInterval)), snapshotChan, 0)
		sw := csv_backend.NewSnapshotWriter(baseFs, snapshotChan, saveEvents)
		enabledExchanges := viper.GetStringSlice(config.Exchanges)
		if err := sd.Init(enabledExchanges...); err != nil {
			log.Fatal(errors.Wrap(errors.Cause(err), "init error"))
		}
		if err := sw.Run(ctx); err != nil {
			log.Fatal(errors.Wrap(errors.Cause(err), "writer start error"))
		}
		if err := sd.Run(ctx); err != nil {
			log.Fatal(errors.Wrap(errors.Cause(err), "downloader start error"))
		}
		exitChan := utils.MakeExitChan()
		<-exitChan
		cancelCtx()
		_ = sw.Close()
		//fmt.Println("snapshot",snapshot.Timestamp.String(), len(snapshot.Orders), snapshot.Orders)
	},
}

func init() {
	rootCmd.AddCommand(runCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// runCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// runCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
