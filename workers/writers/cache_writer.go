package writers

import (
	"context"
	"fmt"
	"github.com/niccoloCastelli/orderbooks/common"
	"github.com/niccoloCastelli/orderbooks/workers"
	"github.com/olivere/elastic"
	"github.com/rs/zerolog"
	"strings"
	"sync"
	"time"
)

type CacheWriter struct {
	inputChan   chan common.Snapshot
	client      *elastic.Client
	logger      zerolog.Logger
	exchange    string
	pair        common.Pair
	ctx         context.Context
	cancelFn    context.CancelFunc
	snapshots   snapshotList
	lock        sync.RWMutex
	updateChans map[string]chan *common.Snapshot
}

func NewCacheWriterWriter(logger zerolog.Logger, exchange string, pair common.Pair) *CacheWriter {
	return &CacheWriter{exchange: exchange, pair: pair, logger: logger.With().Str("component", "writer").Logger(), updateChans: map[string]chan *common.Snapshot{}}
}
func (e *CacheWriter) IndexName() string {
	return strings.ToLower(fmt.Sprintf("%s_%s", e.pair.String(), e.exchange))
}
func (e *CacheWriter) Init(ch chan common.Snapshot) workers.SnapshotWriter {
	e.ctx, e.cancelFn = context.WithCancel(context.Background())
	e.inputChan = ch
	e.snapshots = snapshotList{maxLen: 200}
	return e
}

func (e *CacheWriter) Run(ctx context.Context) error {
	go func() {
		for {
			select {
			case snapshot, ok := <-e.inputChan:
				sn := &snapshot
				if !ok {
					return
				}
				e.snapshots.Add(snapshot)
				e.lock.Lock()
				for _, snapshotCh := range e.updateChans {
					snapshotCh <- sn
				}
				e.logger.Info().Str("now", time.Now().String()).Str("snapshot_time", snapshot.Timestamp.String()).Int("events", len(snapshot.Events)).Msg("add snapshot")
				e.lock.Unlock()
			case <-e.ctx.Done():
				return
			case <-ctx.Done():
				return
			}
		}
	}()
	return nil
}

func (e *CacheWriter) Close() error {
	for _, ch := range e.updateChans {
		close(ch)
	}
	e.cancelFn()
	return nil
}
func (e *CacheWriter) GetSnapshots() []common.Snapshot {
	ret := make([]common.Snapshot, e.snapshots.Len())

	e.snapshots.Iter(func(i int, s common.Snapshot) {
		ret[i] = s
	})
	return ret
}
func (e *CacheWriter) SubscribeUpdates(chanId string, chanSize int) chan *common.Snapshot {
	var (
		ch chan *common.Snapshot
		ok bool
	)
	e.lock.Lock()
	defer e.lock.Unlock()
	if ch, ok = e.updateChans[chanId]; !ok {
		ch = make(chan *common.Snapshot, chanSize)
		e.updateChans[chanId] = ch
	}
	return ch
}
func (e *CacheWriter) UnsubscribeUpdates(chanId string) {
	e.lock.Lock()
	defer e.lock.Unlock()
	ch, ok := e.updateChans[chanId]
	if ok {
		delete(e.updateChans, chanId)
		close(ch)
	}
}

type snapshotList struct {
	first  *snapshotListItem
	last   *snapshotListItem
	lock   sync.RWMutex
	maxLen int
	len    int
}

func (l *snapshotList) Add(s common.Snapshot) {
	l.lock.Lock()
	defer l.lock.Unlock()
	li := snapshotListItem{
		Snapshot: s,
		next:     nil,
		previous: nil,
	}
	if l.first == nil {
		l.first = &li
		l.last = &li
		l.len++
		return
	}
	l.last.next = &li
	li.previous = l.last
	l.last = &li
	l.len++
	if l.len > l.maxLen {
		l.pop()
	}
}
func (l *snapshotList) Pop() *common.Snapshot {
	l.lock.Lock()
	defer l.lock.Unlock()
	return l.pop()
}
func (l *snapshotList) pop() *common.Snapshot {
	if l.first == nil {
		return nil
	}
	ret := l.first
	l.first = l.first.next
	l.first.previous = nil
	l.len--
	return &ret.Snapshot
}
func (l *snapshotList) Len() int {
	l.lock.RLock()
	defer l.lock.RUnlock()
	return l.len
}
func (l *snapshotList) Iter(fn func(i int, s common.Snapshot)) {
	l.lock.RLock()
	defer l.lock.RUnlock()
	i := 0
	nextItem := l.first
	for nextItem != nil {
		fn(i, nextItem.Snapshot)
		i += 1
		nextItem = nextItem.next
	}
}
func (l *snapshotList) Copy() *snapshotList {
	ret := &snapshotList{
		maxLen: l.maxLen,
	}
	l.Iter(func(i int, s common.Snapshot) {
		ret.Add(s)
	})
	return ret
}

type snapshotListItem struct {
	common.Snapshot
	next     *snapshotListItem
	previous *snapshotListItem
}
