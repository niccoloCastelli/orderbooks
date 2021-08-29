package workers

import (
	"context"
	"github.com/niccoloCastelli/orderbooks/common"
	"github.com/niccoloCastelli/orderbooks/utils"
)

type SnapshotWriter interface {
	Init(ch chan common.Snapshot) SnapshotWriter
	Run(ctx context.Context) error
	Close() error
}
type SnapshotReader interface {
	Init() (SnapshotReader, error)
	Close() error
	Exchanges() []string
	Pairs(exchangeName string) []common.Pair
	TimeRange(exchangeName string, pair common.Pair) *utils.TimeRange
	Read(exchangeName string, pair common.Pair, timeRange utils.TimeRange, interval int64, mode common.SnapshotMode, size int, ctx context.Context) (chan common.Snapshot, error)
}
