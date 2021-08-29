package cmd

import (
	"context"
	"encoding/csv"
	"fmt"
	"github.com/jinzhu/copier"
	"github.com/niccoloCastelli/orderbooks/common"
	"github.com/niccoloCastelli/orderbooks/data_formats/cryptotick"
	"github.com/niccoloCastelli/orderbooks/orderbook"
	"github.com/niccoloCastelli/orderbooks/workers/backends/csv_backend"
	"github.com/pkg/errors"
	"github.com/remeh/sizedwaitgroup"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/afero"
	"github.com/spf13/cobra"
	"io/ioutil"
	"os"
	"path"
	"strings"
	"time"
)

const timeOffsetDefaultLayout = "2006-01-02T15:04:05Z0700"
const dateOnlyLayout = "2006-01-02"

type reconstructionConfig struct {
	TimeOffsetStr    string
	TimeOffset       time.Time
	SnapshotInterval time.Duration
	ExchangeName     string
	TimeOffsetLayout string
	OutPath          string
	NumWorkers       int
	MaxSnapshotSize  int
	PriceTicks       bool
	DateUntil        *time.Time
}

func (c *reconstructionConfig) ParseTimeOffset() (err error) {
	c.TimeOffset = time.Time{}
	if c.TimeOffsetStr != "" {
		if c.TimeOffset, err = time.Parse(c.TimeOffsetLayout, c.TimeOffsetStr); err != nil {
			return err
		}
	}
	return nil
}

func initReconstructionCmd() {
	dateUntilStr := ""
	conf := reconstructionConfig{
		TimeOffsetLayout: timeOffsetDefaultLayout,
		ExchangeName:     "cryptotick",
		OutPath:          "./storage",
		NumWorkers:       1,
		MaxSnapshotSize:  1000,
	}

	// reconstructionCmd represents the reconstruction command
	var reconstructionCmd = &cobra.Command{
		Use:   "generate_snapshot",
		Short: "Generate snapshots from cryptotick event files",
		Long:  `Generate a snapshot from cryptotick CSV event files. generate_snapshot [PATH (data snapshot)] [FILENAME_1 FILENAME_2 ...FILENAME_n]`,
		Run: func(cmd *cobra.Command, args []string) {
			var (
				logger = log.With().Str("cmd", "generate_snapshot").Str("out_path", conf.OutPath).Logger()
			)
			if dateUntilStr != "" {
				if dateUntil, err := time.Parse(dateOnlyLayout, dateUntilStr); err == nil {
					conf.DateUntil = &dateUntil
					logger.Info().Time("process_until", *conf.DateUntil).Send()
				} else {
					logger.Err(err).Msg("time parse error")
				}

			}

			switch len(args) {
			case 1:
				logger.Info().Str("mode", "single").Send()
				if err := conf.ParseTimeOffset(); err != nil {
					logger.Fatal().Err(err).Send()
				}
				if err := snapshotReconstruction(args[0], logger, conf); err != nil {
					logger.Err(err).Send()
				}
			case 2:
				if err := bulkSnapshotReconstruction(args[0], args[1], logger, conf); err != nil {
					logger.Err(err).Send()
				}
			default:
				logger.Fatal().Msg("no input file")
			}

			logger.Info().Msg("done")
		},
	}

	reconstructionCmd.PersistentFlags().StringVar(&dateUntilStr, "date_until", "", "Generate snapshots until this date (leave empty to generate snapshots until end of files)")
	reconstructionCmd.PersistentFlags().BoolVar(&conf.PriceTicks, "price_ticks", false, "Generate snapshots every [interval] price changes")
	reconstructionCmd.PersistentFlags().IntVar(&conf.MaxSnapshotSize, "max_snapshot_size", 1000, "Max orders per snapshot (per side, default: 1000)")
	reconstructionCmd.PersistentFlags().IntVar(&conf.NumWorkers, "num_workers", 8, "Worker processes (default: 8)")
	reconstructionCmd.PersistentFlags().StringVar(&conf.OutPath, "out_path", "./storage", "Generated snapthots path (--out_path /my/storage/path)")
	reconstructionCmd.PersistentFlags().StringVar(&conf.ExchangeName, "exchange_name", "cryptotick", "Exchange name (--exchange_name bitstamp)")
	reconstructionCmd.PersistentFlags().StringVar(&conf.TimeOffsetStr, "time_offset", "", "Offset time from first shapshot (--time_offset 2019-12-30T00:00:00)")
	reconstructionCmd.PersistentFlags().StringVar(&conf.TimeOffsetLayout, "time_offset_layout", timeOffsetDefaultLayout, "Time offset date layout (--time_offset_layout 2006-01-02T15:04:05Z0700)")
	reconstructionCmd.PersistentFlags().DurationVar(&conf.SnapshotInterval, "interval", time.Second*30, "30s for time ticks, ")
	rootCmd.AddCommand(reconstructionCmd)
}

func bulkSnapshotReconstruction(filesPath string, fileName string, logger zerolog.Logger, config reconstructionConfig) error {
	dirContent, err := ioutil.ReadDir(filesPath)
	var wg = sizedwaitgroup.New(config.NumWorkers)
	if err != nil {
		return err
	}
	numDirs := len(dirContent)
	logger.Info().Str("mode", "bulk").Int("workers", config.NumWorkers).Int("snapshot_size", config.MaxSnapshotSize).Send()
	startedTasks := 0
	completedTasks := 0

	for i, f := range dirContent {
		var (
			err error
		)
		if !f.IsDir() {
			continue
		}
		absPath := path.Join(filesPath, f.Name(), fileName)
		fLogger := logger.With().Bool("bulk", true).Int("index", i).Int("dir_count", numDirs).Int("batch", i/config.NumWorkers).Logger()
		fileInfo, err := os.Stat(absPath)
		if err != nil {
			fLogger.Debug().Err(err).Msg("skip fileinfo err")
			continue
		}
		if fileInfo.IsDir() {
			fLogger.Error().Msg("is directory")
		}
		newConfig := reconstructionConfig{}
		if err := copier.Copy(&newConfig, &config); err != nil {
			logger.Fatal().Err(err).Send()
		}
		newConfig.TimeOffsetStr = f.Name()
		newConfig.TimeOffset = time.Time{}
		if err = newConfig.ParseTimeOffset(); err != nil {
			logger.Fatal().Err(err).Send()
		}
		if newConfig.DateUntil != nil && !newConfig.TimeOffset.IsZero() && newConfig.TimeOffset.Sub(*newConfig.DateUntil) >= time.Hour*24 {
			logger.Info().Time("time_offset time", newConfig.TimeOffset).Msg("time limit reached")
			break
		}
		wg.Add()
		startedTasks++
		go func(absPath string, fLogger zerolog.Logger, newConfig reconstructionConfig, wg *sizedwaitgroup.SizedWaitGroup) {
			defer wg.Done()
			if err := snapshotReconstruction(absPath, fLogger, newConfig); err != nil {
				fLogger.Fatal().Err(err).Send()
			}
			completedTasks++
		}(absPath, fLogger, newConfig, &wg)
	}
	logger.Info().Int("started", startedTasks).Int("completed", completedTasks).Msg("waiting... ")
	wg.Wait()
	return nil
}
func snapshotReconstruction(eventsFile string, logger zerolog.Logger, config reconstructionConfig) error {
	var (
		snapshotCount  int
		snapshotChan   = make(chan common.Snapshot, 32)
		sw             = csv_backend.NewSnapshotWriter(afero.NewBasePathFs(afero.NewOsFs(), config.OutPath), snapshotChan, true)
		ctx, cancelCtx = context.WithCancel(context.Background())
	)
	_, err := os.Stat(eventsFile)
	if err != nil {
		logger.Fatal().Err(err).Send()
	}
	csvFile, err := os.Open(eventsFile)
	if err != nil {
		logger.Fatal().Err(err).Send()
	}
	if err := sw.Run(ctx); err != nil {
		logger.Fatal().Err(errors.Wrap(errors.Cause(err), "writer start error")).Send()
	}
	logger = logger.With().Str("file", eventsFile).Logger()
	logger.Debug().Msg("open csv file ok")
	logger.Info().Str("interval", fmt.Sprintf("%v", config.SnapshotInterval)).Str("offset", config.TimeOffset.String()).Msg("open csv file ok")
	defer csvFile.Close()
	reader := csv.NewReader(csvFile)
	reader.Comma = ';'
	ob := orderbook.NewAddModeOrderBook(logger, nil).LimitSnapshotSize(config.MaxSnapshotSize).SetTimeOffset(config.TimeOffset)
	previousUpdate := ob.LastUpdate()
	lastUpdateType := cryptotick.UpdateModeSnapshot
	resetPerformed := false
	priceUpdated := false
	lastAsk, lastBid := ob.BestPrices()
	for i := 0; ; i++ {
		resetPerformed = false
		priceUpdated = false
		csvLine, err := reader.Read()
		if err != nil {
			logger.Debug().Err(err)
			break
		}
		if i == 0 {
			logger.Debug().Str("headers", strings.Join(csvLine, ", ")).Send()
			continue
		}
		parsedRow, err := cryptotick.ParseCryptotyickRow(csvLine)
		if err != nil {
			logger.Error().Err(err).Send()
			break
		}
		if parsedRow.UpdateType == cryptotick.UpdateModeSnapshot && lastUpdateType != cryptotick.UpdateModeSnapshot {
			ob = orderbook.NewAddModeOrderBook(logger, nil).LimitSnapshotSize(config.MaxSnapshotSize).SetTimeOffset(config.TimeOffset)
			resetPerformed = true
		}
		ob.Update(parsedRow.ToOrder(), false)
		lastUpdateType = parsedRow.UpdateType
		if previousUpdate.IsZero() {
			previousUpdate = ob.LastUpdate()
			continue
		}
		lastUpdate := ob.LastUpdate()
		if parsedRow.UpdateType != cryptotick.UpdateModeSnapshot {
			bestAsk, bestBid := ob.BestPrices()
			if bestAsk != lastAsk || bestBid != lastBid {
				priceUpdated = true
			}
		}
		if !resetPerformed && ((!config.PriceTicks && lastUpdate.Sub(previousUpdate) > config.SnapshotInterval) || (config.PriceTicks && priceUpdated)) {
			snapshot := ob.GetSnapshot()
			snapshot.Exchange = "cryptotick_bitstamp" //TODO: parametrizzare
			snapshot.Pair = common.Pair{Base: "BTC", Quote: "USD"}
			snapshotChan <- snapshot
			snapshotCount++
			logger.Debug().Int("snapshot_count", snapshotCount).Str("last_update", lastUpdate.String()).Msg("snapshot")
			previousUpdate = ob.LastUpdate()
		}
		lastAsk, lastBid = ob.BestPrices()
	}
	cancelCtx()
	if err := sw.Close(); err != nil {
		return err
	}
	logger.Info().Msg("file OK")
	return nil
}

func init() {
	initReconstructionCmd()

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// reconstructionCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// reconstructionCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
