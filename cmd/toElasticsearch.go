/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

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
	"errors"
	"github.com/niccoloCastelli/orderbooks/common"
	"github.com/niccoloCastelli/orderbooks/utils"
	"github.com/niccoloCastelli/orderbooks/workers/backends/csv_backend"
	"github.com/niccoloCastelli/orderbooks/workers/writers"
	"github.com/rs/zerolog/log"
	"github.com/spf13/afero"
	"time"

	"github.com/spf13/cobra"
)

const cliDateLayout = "2006-01-02"

var (
	storagePath         = "./storage"
	interval      int64 = 10
	snapshotMode        = int(common.SnapshotModeTicks)
	depth               = 10
	esIndexPrefix       = ""
)

// toElasticsearchCmd represents the toElasticsearch command
var toElasticsearchCmd = &cobra.Command{
	Use:   "to_elasticsearch",
	Short: "Upload csv data to elasticsearch",
	Long:  `to_elasticsearch [pair] [exchange] [date_start] [date_end]`,
	RunE: func(cmd *cobra.Command, args []string) error {
		var (
			pair          *common.Pair
			exchange      string
			dateStart     time.Time
			dateEnd       time.Time
			err           error
			now           = time.Now()
			ctx, cancelFn = context.WithCancel(context.Background())
			esUrl         = utils.EnvOrDefault("ELASTICSEARCH_URL", "http://localhost:9200")
		)
		defer cancelFn()
		if len(args) != 4 {
			return errors.New("usage: to_elasticsearch [pair] [exchange] [date_start] [date_end]")
		}

		if pair, err = common.PairsFromString(args[0]); err != nil {
			return err
		}
		exchange = args[1]
		dtStartStr := args[2]

		if dateStart, err = time.Parse(cliDateLayout, dtStartStr); err != nil {
			return err
		}

		dtEndStr := args[3]
		if dateEnd, err = time.Parse(cliDateLayout, dtEndStr); err != nil {
			return err
		}
		readRange := utils.NewTimeRange(dateStart, dateEnd)

		fs := afero.NewBasePathFs(afero.NewOsFs(), storagePath)
		logger := log.Logger.With().Str("command", "elasticsearch").Time("started", now).Logger()
		reader, err := csv_backend.NewSnapshotReader(fs, &logger)
		if err != nil {
			return err
		}
		if snapshotMode == int(common.SnapshotModeTime) {
			interval = int64(time.Second) * interval
		}
		ch, err := reader.Read(exchange, *pair, readRange, interval, common.SnapshotMode(snapshotMode), 10, ctx)
		if err != nil {
			return err
		}
		esWriter := writers.NewElasticSearchWriter(logger.Level(-1), exchange, *pair, writers.NewElasticSearchWriterConfig(esIndexPrefix, 1000, 1000), esUrl).Init(ch)

		if err := esWriter.Run(ctx); err != nil {
			return err
		}
		defer esWriter.Close()
		exitChan := utils.MakeExitChan()
		<-exitChan
		/**/

		return nil
	},
}

func init() {
	rootCmd.AddCommand(toElasticsearchCmd)
	rootCmd.PersistentFlags().StringVar(&storagePath, "storage_path", "./storage", "Path storage")
	toElasticsearchCmd.PersistentFlags().Int64Var(&interval, "interval", 10, "Snapshot interval")
	toElasticsearchCmd.PersistentFlags().IntVar(&snapshotMode, "snapshot_mode", 1, "Snapshot mode (0=time 1=ticks)")
	toElasticsearchCmd.PersistentFlags().IntVar(&depth, "depth", 10, "Numero livelli order book (10 = 10 bids + 10 asks)")
	toElasticsearchCmd.PersistentFlags().StringVar(&esIndexPrefix, "prefix", "", "Prefisso indice elasticsearch")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// toElasticsearchCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// toElasticsearchCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
