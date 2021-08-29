package config

import "github.com/niccoloCastelli/orderbooks/common"

const (
	Pairs            = "pairs"
	StoragePath      = "storage_path"
	SnapshotInterval = "snapshot_interval"
	Events           = "events"
	Exchanges        = "exchanges"
	SaveEvents       = "save_events"
)

type Config struct {
	Pairs []common.Pair
}

func NewConfig() *Config {
	return &Config{
		Pairs: []common.Pair{{"BTC", "USD"}},
	}
}

type ServerConfig struct {
	Host        string
	StoragePath string
}
