package orderbook

import (
	"fmt"
	"github.com/niccoloCastelli/orderbooks/common"
	"github.com/rs/zerolog"
	"sync"
	"time"
)

type limit struct {
	Amount float64
	Price  float64
	Left   *limit
	Right  *limit
}

type OrderBook struct {
	bids            *nodeTree
	asks            *nodeTree
	obLock          sync.RWMutex
	logger          zerolog.Logger
	events          EventQueue
	eventsLock      sync.RWMutex
	addMode         bool // true: `Amount` inviato come differenza, false: `Amount` inviato come importo assoluto per il livello di prezzo
	lastUpdate      time.Time
	timeOffset      time.Time
	maxSnapshotSize int
	onUpdate        func(initial bool)
}

func (o *OrderBook) SetTimeOffset(t time.Time) *OrderBook {
	o.timeOffset = t //o.timeOffset.Add(time.Duration(ts))
	return o
}
func (o *OrderBook) LimitSnapshotSize(val int) *OrderBook {
	o.maxSnapshotSize = val
	return o
}
func (o *OrderBook) LastUpdate() time.Time {
	return o.lastUpdate
}
func (o *OrderBook) Update(order common.Order, initial bool) {
	//fmt.Println("update", order)
	o.obLock.Lock()
	defer o.obLock.Unlock()
	if order.Side == common.OrderSideBid {
		o.bids.UpdateTree(order.Amount, order.Price)
		if o.bids.head != nil && o.asks.head != nil && o.bids.tail.Price >= o.asks.head.Price {
			fmt.Println("Bad snapshot", order.Timestamp.String(), o.bids.tail.Price, o.asks.head.Price)
		}
	} else {
		o.asks.UpdateTree(order.Amount, order.Price)
		if o.bids.head != nil && o.asks.head != nil && o.bids.tail.Price >= o.asks.head.Price {
			fmt.Println("Bad snapshot", order.Timestamp.String(), o.bids.tail.Price, o.asks.head.Price)
		}
	}
	updateTime := order.Timestamp
	if order.Timestamp.IsZero() {
		updateTime = time.Now()
	}
	if o.onUpdate != nil {
		go func(o *OrderBook) {
			o.onUpdate(initial)
		}(o)
	}
	o.lastUpdate = updateTime
}
func (o *OrderBook) GetSnapshot() common.Snapshot {
	o.obLock.Lock()
	bids := o.bids.GenerateIndex(o.maxSnapshotSize)
	asks := o.asks.GenerateIndex(o.maxSnapshotSize)
	events := o.events.RemoveAll()
	o.obLock.Unlock()

	ret := make([]common.Order, 0, len(bids)+len(asks))
	ts := o.lastUpdate
	if !o.timeOffset.IsZero() {
		ts = ts.AddDate(o.timeOffset.Year(), int(o.timeOffset.Month()-1), o.timeOffset.Day()-1).Add(time.Duration(o.timeOffset.Hour())*time.Hour + time.Duration(o.timeOffset.Minute())*time.Minute + time.Duration(o.timeOffset.Second())*time.Second)
	}
	year := ts.Year()
	if year == 0 {
		o.logger.Error().Msg("year 0")
	}
	for _, bid := range bids {
		ret = append(ret, common.Order{
			Timestamp: ts,
			Side:      common.OrderSideBid,
			Amount:    bid.Amount,
			Price:     bid.Price,
		})
	}
	for _, ask := range asks {
		ret = append(ret, common.Order{
			Timestamp: ts,
			Side:      common.OrderSideAsk,
			Amount:    ask.Amount,
			Price:     ask.Price,
		})
	}
	var (
		bidPrice float64 = 0
		askPrice float64 = 0
	)
	if o.asks.head != nil {
		askPrice = o.asks.head.Price
	}
	if o.bids.tail != nil {
		bidPrice = o.bids.tail.Price
	}
	o.logger.Debug().
		Float64("bid", bidPrice).
		Float64("ask", askPrice).
		Float64("mid_price", (bidPrice+askPrice)/2).
		Int("num_asks", len(asks)).
		Int("num_bids", len(bids)).
		Int("events", len(events)).
		Time("timestamp", ts).
		Msg("snapshot")

	return common.Snapshot{
		Timestamp: ts,
		Orders:    ret,
		Events:    events,
	}
}
func (o *OrderBook) GetEvents() []common.Event {
	o.obLock.Lock()
	defer o.obLock.Unlock()
	return o.events.RemoveAll()
}
func (o *OrderBook) BestPrices() (float64, float64) {
	if o.asks == nil || o.asks.head == nil || o.bids == nil || o.bids.head == nil {
		return -1, -1
	}
	return o.asks.head.Price, o.bids.tail.Price
}
func (o *OrderBook) registerEvent(order common.Order, eventType common.EventType) {
	o.eventsLock.Lock()
	defer o.eventsLock.Unlock()
	o.events.Add(common.Event{
		Order: order,
		Type:  eventType,
	})
}
func (o *OrderBook) registerEventBid(price float64, amount float64, eventType common.EventType) {
	o.registerEvent(common.Order{
		Timestamp: time.Now(),
		Side:      common.OrderSideBid,
		Amount:    amount,
		Price:     price,
	}, eventType)
}
func (o *OrderBook) registerEventAsk(price float64, amount float64, eventType common.EventType) {
	o.registerEvent(common.Order{
		Timestamp: time.Now(),
		Side:      common.OrderSideAsk,
		Amount:    amount,
		Price:     price,
	}, eventType)
}

func newOrderBook(logger zerolog.Logger, addMode bool, onUpdate func(bool)) *OrderBook {
	askLogger := logger.With().Str("side", "ask").Logger()
	bidLogger := logger.With().Str("side", "bid").Logger()
	ob := &OrderBook{
		addMode:  addMode,
		asks:     &nodeTree{logger: askLogger, addMode: addMode},
		bids:     &nodeTree{reverse: true, logger: bidLogger, addMode: addMode},
		logger:   logger.With().Bool("orderbook", true).Logger(),
		events:   EventQueue{},
		onUpdate: onUpdate,
	}
	ob.asks.registerEvent = ob.registerEventAsk
	ob.bids.registerEvent = ob.registerEventBid
	return ob
}
func NewOrderBook(logger zerolog.Logger, onUpdate func(bool)) *OrderBook {
	return newOrderBook(logger, false, onUpdate)
}
func NewAddModeOrderBook(logger zerolog.Logger, onUpdate func(bool)) *OrderBook {
	return newOrderBook(logger, true, onUpdate)
}

type EventQueue struct {
	head  *eventQueueItem
	tail  *eventQueueItem
	count int
}

func (q *EventQueue) Count() int {
	return q.count
}
func (q *EventQueue) Add(e common.Event) {
	if q.head == nil {
		q.head = &eventQueueItem{event: e}
		q.tail = q.head
		return
	}
	oldTail := q.tail
	newTail := &eventQueueItem{event: e, left: oldTail}
	oldTail.right = newTail
	q.tail = newTail
	q.count++
}
func (q *EventQueue) Remove() *common.Event {
	if q.head == nil {
		return nil
	}
	oldHead := q.head
	if q.head == q.tail {
		q.tail = nil
		q.head = nil
	} else {
		q.head = q.head.right
	}
	q.count--
	return &oldHead.event
}
func (q *EventQueue) RemoveAll() []common.Event {
	retCap := q.count
	if retCap < 1 {
		retCap = 1
	}
	ret := make([]common.Event, 0, retCap)
	if q.head == nil {
		return ret
	}
	for {
		evt := q.Remove()
		if evt == nil {
			break
		}
		ret = append(ret, *evt)
	}
	return ret
}

type eventQueueItem struct {
	event common.Event
	left  *eventQueueItem
	right *eventQueueItem
}
