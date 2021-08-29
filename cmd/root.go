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
	"github.com/niccoloCastelli/orderbooks/config"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/rs/zerolog/pkgerrors"
	"github.com/spf13/cobra"
	"os"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

var cfgFile string
var logLevel = int8(zerolog.InfoLevel)
var pairs []common.Pair

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "orderbooks",
	Short: "Downloader orderbooks",
	Long:  `Downloader orderbooks`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	//	Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Error().Err(err).Send()
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "config.json", "config file (default is ./config.json)")
	rootCmd.PersistentFlags().Int8Var(&logLevel, "log_level", int8(zerolog.InfoLevel), "-1(trace)... 5(panic)")

}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	viper.SetDefault(config.Pairs, []string{common.Pair{Base: "BTC", Quote: "USD"}.String()})
	viper.SetDefault(config.SnapshotInterval, 30)
	viper.SetDefault(config.StoragePath, "./storage")
	viper.SetDefault(config.Exchanges, []string{})

	if logLevel < -1 || logLevel > 5 {
		log.Fatal().Msgf("invalid log level: %d", logLevel)
	}
	zerolog.SetGlobalLevel(zerolog.Level(logLevel))

	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			os.Exit(1)
		}

		// Search config in home directory with name ".orderbooks" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".orderbooks")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		log.Info().Msgf("Using config file: %s", viper.ConfigFileUsed())
	}

	var err error
	rawPairs := viper.GetStringSlice(config.Pairs)
	pairs, err = common.PairsFromStrings(rawPairs...)

	zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack
	if err != nil {
		log.Fatal().Err(err).Send()
	}
}
