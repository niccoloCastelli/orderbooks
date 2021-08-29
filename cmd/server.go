package cmd

import (
	"github.com/niccoloCastelli/orderbooks/config"
	"github.com/niccoloCastelli/orderbooks/server"
	"github.com/rs/zerolog/log"

	"github.com/spf13/cobra"
)

func initServerCmd() {
	// serverCmd represents the server command
	var cfg = config.ServerConfig{
		Host:        "0.0.0.0:60000",
		StoragePath: "./storage",
	}
	var serverCmd = &cobra.Command{
		Use:   "server",
		Short: "Run gRPC server",
		Long:  `Run gRPC server`,
		RunE: func(cmd *cobra.Command, args []string) error {
			logger := log.With().Str("cmd", "server").Logger()
			if err := server.NewGrpcServer(&logger, cfg); err != nil {
				return err
			}
			return nil
		},
	}
	serverCmd.PersistentFlags().StringVarP(&cfg.Host, "host", "a", "0.0.0.0:60000", "Server url")
	serverCmd.PersistentFlags().StringVar(&cfg.StoragePath, "storage_path", "./storage", "Storage path")
	rootCmd.AddCommand(serverCmd)
}

func init() {
	initServerCmd()
}
