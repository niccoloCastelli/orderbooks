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

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Start order books scraper",
	Long:  `Start order books scraper`,
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
	},
}

func init() {
	rootCmd.AddCommand(runCmd)
}
