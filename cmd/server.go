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
		Short: "Avvia server",
		Long:  `Avvia server HTTP/GRPC`,
		RunE: func(cmd *cobra.Command, args []string) error {
			logger := log.With().Str("cmd", "server").Logger()
			if err := server.NewGrpcServer(&logger, cfg); err != nil {
				return err
			}
			return nil
		},
	}
	serverCmd.PersistentFlags().StringVarP(&cfg.Host, "host", "a", "0.0.0.0:60000", "Indirizzo e porta server")
	serverCmd.PersistentFlags().StringVar(&cfg.StoragePath, "storage_path", "./storage", "Path storage")
	rootCmd.AddCommand(serverCmd)
}

func init() {
	initServerCmd()

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// serverCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// serverCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
