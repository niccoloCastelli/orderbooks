package coinbase

import (
	"github.com/niccoloCastelli/orderbooks/common"
	"github.com/niccoloCastelli/orderbooks/orderbook"
	"time"
)

const (
	wsChanOrderBook = "level2"
)

type wsChannel struct {
	Name       string   `json:"name"`
	ProductIds []string `json:"product_ids"`
}

type subscribeOrderBookMsg struct {
	Type     string      `json:"type"`
	Channels []wsChannel `json:"channels"`
}

func newSubscribeOrderBookMsg(pairs ...common.Pair) *subscribeOrderBookMsg {
	procuctIds := make([]string, 0, len(pairs))
	for _, pair := range pairs {
		procuctIds = append(procuctIds, pair.String())
	}
	channel := wsChannel{
		Name:       wsChanOrderBook,
		ProductIds: procuctIds,
	}
	return &subscribeOrderBookMsg{
		Type:     "subscribe",
		Channels: []wsChannel{channel},
	}
}

type orderBookUpdateMsg struct {
	Type      string     `json:"type"`
	ProductID string     `json:"product_id"`
	Time      time.Time  `json:"time"`
	Changes   [][]string `json:"changes"`
}

func (o orderBookUpdateMsg) UpdateOrderBook(orderBook *orderbook.OrderBook, initial bool) {
	for _, newOrder := range o.Changes {
		order := common.Order{
			Timestamp: time.Time{},
		}
		if err := sidedRowToOrder(newOrder, &order); err != nil {
			continue
		}
		orderBook.Update(order, initial)
	}
	//fmt.Println("update from resp", len(o.Asks), len(o.Bids), len(orderBook.GetSnapshot().Orders))
}
