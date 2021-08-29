package exchanges

import (
	"context"
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/niccoloCastelli/orderbooks/common"
	"github.com/niccoloCastelli/orderbooks/orderbook"
	"github.com/niccoloCastelli/orderbooks/utils"
	"github.com/pkg/errors"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"io/ioutil"
	"net/http"
	"sync"
	"sync/atomic"
	"time"
)

const eventTicksCount = 10

func NewExchange(name string, availablePairs []common.Pair) *Exchange {
	return &Exchange{
		name:           name,
		availablePairs: availablePairs,
		Ws:             map[string]*websocket.Conn{},
		OrderBooks:     map[string]*orderbook.OrderBook{},
		tickCounters:   map[string]*uint64{},
		eventChans:     map[string]chan uint64{},
		Lock:           sync.Mutex{},
		eventLock:      sync.Mutex{},
		Logger:         log.Logger.With().Str("exchange", name).Logger(),
	}
}

type Exchange struct {
	name              string
	availablePairs    []common.Pair
	Ws                map[string]*websocket.Conn
	OrderBooks        map[string]*orderbook.OrderBook
	tickCounters      map[string]*uint64
	eventChans        map[string]chan uint64
	Lock              sync.Mutex
	eventLock         sync.Mutex
	Logger            zerolog.Logger
	snapshotSizeLimit int
}

func (e *Exchange) Name() string {
	return e.name
}
func (e *Exchange) SetMaxSnapshotSize(val int) {
	e.snapshotSizeLimit = val
}
func (e *Exchange) Close() error {
	errs := []string{}
	for key, ws := range e.Ws {
		if err := ws.Close(); err != nil {
			errs = append(errs, fmt.Sprintf("%s: %v", key, err))
		}
	}
	return utils.JoinErrsS(errs...)
}
func (e *Exchange) AvailablePairs() []common.Pair {
	dst := make([]common.Pair, len(e.availablePairs))
	copy(dst, e.availablePairs)
	return dst
}
func (e *Exchange) PairAvailable(pair common.Pair) bool {
	for _, p := range e.availablePairs {
		if p.String() == pair.String() {
			return true
		}
	}
	return false
}
func (e *Exchange) GetSnapshot(pair common.Pair) (common.Snapshot, error) {
	ob, ok := e.OrderBooks[pair.String()]
	if !ok {
		return common.Snapshot{}, errors.New(fmt.Sprintf("pair %s not initialized", pair.String()))
	}
	var snapshot common.Snapshot
	utils.TimeProfile("snapshot", func() {
		snapshot = ob.GetSnapshot()
		snapshot.Pair = pair
		snapshot.Exchange = e.Name()
	})
	return snapshot, nil
}
func (e *Exchange) CheckDisconnections(c *websocket.Conn, logger zerolog.Logger) (context.Context, context.CancelFunc, error) {
	return CheckDisconnections(c, logger)
}
func (e *Exchange) CheckResponse(resp *http.Response) error {
	var err error
	if resp.StatusCode != http.StatusOK {
		respData := []byte{}
		if respData, err = ioutil.ReadAll(resp.Body); err != nil {
			return errors.WithStack(err)
		}
		return errors.WithStack(errors.New(fmt.Sprintf("%d - %s", resp.StatusCode, string(respData))))
	}
	return nil
}
func (e *Exchange) SubscribeUpdates(pair common.Pair, chanSize int) chan uint64 {
	if ch, ok := e.eventChans[pair.String()]; !ok {
		ch = make(chan uint64, chanSize)
		e.eventChans[pair.String()] = ch
		return ch
	}
	return nil
}
func (e *Exchange) UnsubscribeUpdates(pair common.Pair) bool {
	_, ok := e.eventChans[pair.String()]
	if ok {
		delete(e.eventChans, pair.String())
	}
	return ok
}
func (e *Exchange) NewOrderBook(pair common.Pair, pairLogger zerolog.Logger) *orderbook.OrderBook {
	orderBook := orderbook.NewOrderBook(pairLogger, func(initial bool) {
		e.OnOrderBookUpdate(initial, pair)
	})
	if e.snapshotSizeLimit != 0 {
		orderBook.LimitSnapshotSize(e.snapshotSizeLimit)
	}
	return orderBook
}
func (e *Exchange) GetOrderBook(pair common.Pair) *orderbook.OrderBook {
	return e.OrderBooks[pair.String()]
}
func (e *Exchange) OnOrderBookUpdate(initial bool, pair common.Pair) {
	var (
		counter *uint64
		ok      bool
	)
	e.eventLock.Lock()
	defer e.eventLock.Unlock()
	if counter, ok = e.tickCounters[pair.String()]; !ok {
		var c uint64 = 0
		counter = &c
		e.tickCounters[pair.String()] = counter
	}
	atomic.AddUint64(counter, 1)
	if *counter < eventTicksCount {
		return
	}
	if initial {
		return
	}
	if ch, ok := e.eventChans[pair.String()]; ok {
		ch <- *counter
	}

	// c := *counter
	// e.Logger.Info().Str("pair", pair.String()).Uint64("counter", c).Int("now", time.Now().Nanosecond()).Msg("ob update")

	atomic.StoreUint64(counter, 0)
}

func CheckDisconnections(c *websocket.Conn, logger zerolog.Logger) (context.Context, context.CancelFunc, error) {
	wsCtx, cancelWs := context.WithCancel(context.Background())
	if err := c.SetReadDeadline(time.Now().Add(time.Second * 5)); err != nil {
		return wsCtx, cancelWs, err
	}
	go func() {
		ticker := time.NewTicker(time.Second * 2)
		for {
			select {
			case <-ticker.C:
				logger.Trace().Msg("ping")
				if err := c.WriteMessage(websocket.PingMessage, []byte{}); err != nil {
					return
				}
			case <-wsCtx.Done():
				return
			}
		}
	}()
	c.SetPongHandler(func(string) error {
		logger.Trace().Msg("pong")
		return c.SetReadDeadline(time.Now().Add(time.Second * 5))
	})
	return wsCtx, cancelWs, nil
}
