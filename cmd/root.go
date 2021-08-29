package cmd

import (
	"github.com/niccoloCastelli/orderbooks/common"
	"github.com/niccoloCastelli/orderbooks/config"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/rs/zerolog/pkgerrors"
	"github.com/spf13/cobra"
	"os"

	"github.com/spf13/viper"
)

var cfgFile string
var logLevel = int8(zerolog.InfoLevel)
var pairs []common.Pair

var rootCmd = &cobra.Command{
	Use:   "orderbooks",
	Short: "Downloader orderbooks",
	Long:  `Downloader orderbooks`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Error().Err(err).Send()
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Global flags
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
		// Search config in home directory with name ".orderbooks" (without extension).
		viper.AddConfigPath("./")
		viper.SetConfigName("config.json")
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
