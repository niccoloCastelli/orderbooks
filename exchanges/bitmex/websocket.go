package bitmex

import (
	"fmt"
	"github.com/niccoloCastelli/orderbooks/common"
	"github.com/niccoloCastelli/orderbooks/orderbook"
	"time"
)

const (
	wsChanOrderBook = "orderBookL2:%s"
	orderBookDepth  = 1000
)

//{"op": "subscribe", "args": ["spot/depth_l2_tbt:BTC-USDT"]}
type subscribeOrderBookMsg struct {
	Op   string   `json:"op"`
	Args []string `json:"args"`
}

type SubscribeResp struct {
	Success   bool   `json:"success"`
	Subscribe string `json:"subscribe"`
	Request   struct {
		Op   string   `json:"op"`
		Args []string `json:"args"`
	} `json:"request"`
}

func newSubscribeOrderBookMsg(pairs ...common.Pair) *subscribeOrderBookMsg {
	pairChans := make([]string, 0, len(pairs))
	for _, pair := range pairs {
		pairChans = append(pairChans, fmt.Sprintf(wsChanOrderBook, formatPair(pair)))
	}
	return &subscribeOrderBookMsg{
		Op:   "subscribe",
		Args: pairChans,
	}
}

type orderBookUpdateMsg struct {
	Table string   `json:"table"`
	Keys  []string `json:"keys"`
	Types struct {
		ID     string `json:"id"`
		Price  string `json:"price"`
		Side   string `json:"side"`
		Size   string `json:"size"`
		Symbol string `json:"symbol"`
	} `json:"types"`
	ForeignKeys struct {
		Side   string `json:"side"`
		Symbol string `json:"symbol"`
	} `json:"foreignKeys"`
	Attributes struct {
		ID     string `json:"id"`
		Symbol string `json:"symbol"`
	} `json:"attributes"`
	Action string `json:"action"`
	Data   []struct {
		Symbol string  `json:"symbol"`
		ID     int64   `json:"id"`
		Side   string  `json:"side"`
		Size   float64 `json:"size"`
		Price  float64 `json:"price"`
	} `json:"data"`
}

func (o orderBookUpdateMsg) UpdateOrderBook(orderBook *orderbook.OrderBook, instrumentIdx int, instrument InstrumentInfo, initial bool) {
	for _, newOrder := range o.Data {
		order := common.Order{
			Timestamp: time.Now(),
			Amount:    newOrder.Size,
			Price:     newOrder.Price,
		}
		if newOrder.Side == "Sell" {
			order.Side = common.OrderSideAsk
		} else {
			order.Side = common.OrderSideBid
		}
		if o.Action != "partial" {
			order.Price = priceFromID(newOrder.ID, int64(instrumentIdx), &instrument)
		}
		orderBook.Update(order, initial)
	}
	//fmt.Println("update from resp", len(o.Asks), len(o.Bids), len(orderBook.GetSnapshot().Orders))
}
