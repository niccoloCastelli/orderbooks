package cmd

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/google/uuid"
	"github.com/nats-io/nats.go"
	"github.com/niccoloCastelli/orderbooks/common"
	"github.com/niccoloCastelli/orderbooks/server"
	"github.com/niccoloCastelli/orderbooks/utils"
	"github.com/niccoloCastelli/orderbooks/workers/downloaders"
	"github.com/niccoloCastelli/orderbooks/workers/writers"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"time"

	"github.com/spf13/cobra"
)

var (
	saveFeed bool
)

// toElasticsearchCmd represents the toElasticsearch command
var dataFeedCmd = &cobra.Command{
	Use:   "data_feed",
	Short: "Save snapshots on elasticsearch",
	Long:  `data_feed [pair] [exchange]`,
	RunE: func(cmd *cobra.Command, args []string) error {
		var (
			pair          *common.Pair
			exchange      string
			err           error
			now           = time.Now()
			ctx, cancelFn = context.WithCancel(context.Background())
			snapshotSize  = 10
			esWriter      *writers.ElasticSearchWriter
			esChan        chan common.Snapshot
			sessionId     = int64(uuid.New().ID())
			counter       int64
			esUrl         = utils.EnvOrDefault("ELASTICSEARCH_URL", "http://localhost:9200")
			natsUrl       = utils.EnvOrDefault("NATS_URL", nats.DefaultURL)
		)
		defer cancelFn()
		if len(args) != 2 {
			return errors.New("usage: data_feed [pair] [exchange]")
		}
		//http://localhost:9200

		if pair, err = common.PairsFromString(args[0]); err != nil {
			return err
		}
		exchange = args[1]

		logger := log.Logger.With().Int64("session_id", sessionId).Str("command", "data_feed").Bool("save", saveFeed).Time("started", now).Logger()
		snapshotChan := make(chan common.Snapshot, 10)
		liveDownloader := downloaders.NewSnapshotDownloaderTicks([]common.Pair{*pair}, int(interval), snapshotChan, snapshotSize)
		if err := liveDownloader.Init(exchange); err != nil {
			return err
		}
		if err := liveDownloader.Run(ctx); err != nil {
			return err
		}

		esWriterConf := writers.NewElasticSearchWriterConfig("feed", 100, 1)
		esWriter = writers.NewElasticSearchWriter(logger, exchange, *pair, esWriterConf, esUrl)

		if saveFeed {
			esChan = make(chan common.Snapshot, 10)
			esWriter.Init(esChan)
			if err = esWriter.Run(ctx); err != nil {
				return err
			}
			defer esWriter.Close()
		}

		natsConn, err := nats.Connect(natsUrl)
		if err != nil {
			return err
		}

		go func(globalLogger zerolog.Logger) {
			natsSubj := writers.UniqueName("", *pair, exchange)
			started := time.Now()
			logger := globalLogger.With().Str("started", started.String()).Str("channel_name", natsSubj).Logger()
			logger.Info().Msg("feed started")
			for {
				select {
				case snapshot := <-snapshotChan:
					logger.Trace().Int("events", len(snapshot.Events)).Str("snapshot_time", snapshot.Timestamp.String()).Msg("snapshot received")
					snapshot.SessionID = sessionId
					snapshot.Counter = counter
					if esChan != nil {
						esChan <- snapshot
					}
					msg, err := json.Marshal(server.SnapshotToPb(&snapshot))
					if err != nil {
						logger.Err(err).Msg("proto marshal error")
						continue
					}
					if err := natsConn.Publish(natsSubj, msg); err != nil {
						logger.Err(err).Msg("nats publish error")
						return
					}
					if counter%1000 == 0 {
						logger.Info().Int64("counter", counter).Str("now", time.Now().String()).Msg("msg counter")
					}
					counter++
				case <-ctx.Done():
					return
				}
			}
		}(logger)
		exitChan := utils.MakeExitChan()
		<-exitChan
		/**/

		return nil
	},
}

func init() {
	rootCmd.AddCommand(dataFeedCmd)
	dataFeedCmd.PersistentFlags().Int64Var(&interval, "interval", 10, "Snapshot interval")
	dataFeedCmd.PersistentFlags().IntVar(&snapshotMode, "snapshot_mode", 1, "Snapshot mode (0=time 1=ticks)")
	dataFeedCmd.PersistentFlags().IntVar(&depth, "depth", 10, "Order book levels to stream (10 = 10 bids + 10 asks)")
	dataFeedCmd.PersistentFlags().BoolVar(&saveFeed, "save", false, "Save feed on elasticsearch")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// toElasticsearchCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// toElasticsearchCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
