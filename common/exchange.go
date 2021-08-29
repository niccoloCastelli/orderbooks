package common

import "context"

type Exchange interface {
	Init(config interface{}) error
	Name() string
	AvailablePairs() []Pair
	PairAvailable(pair Pair) bool
	GetSnapshot(pair Pair) (Snapshot, error)
	SubscribeEvents(pair Pair) (context.Context, error)

	SubscribeUpdates(pair Pair, chanSize int) chan uint64
	UnsubscribeUpdates(pair Pair) bool
	SetMaxSnapshotSize(val int)

	Close() error
}

type SnapshotMode int

const (
	SnapshotModeTime SnapshotMode = iota
	SnapshotModeTicks
)
