package okex

import (
	"github.com/niccoloCastelli/orderbooks/common"
	"github.com/niccoloCastelli/orderbooks/orderbook"
	"time"
)

type orderBookResponse struct {
	Timestamp int        `json:"sequence"`
	Pair      string     `json:"product_id"`
	Type      string     `json:"type"`
	Bids      [][]string `json:"bids"`
	Asks      [][]string `json:"asks"`
}

func (o orderBookResponse) UpdateOrderBook(orderBook *orderbook.OrderBook, initial bool) {
	for _, ask := range o.Asks {
		order := common.Order{
			Timestamp: time.Time{},
			Side:      common.OrderSideAsk,
		}
		if err := rowToOrder(ask, &order); err != nil {
			continue
		}
		orderBook.Update(order, initial)
	}
	for _, bid := range o.Bids {
		order := common.Order{
			Timestamp: time.Time{},
			Side:      common.OrderSideBid,
		}
		if err := rowToOrder(bid, &order); err != nil {
			continue
		}
		orderBook.Update(order, initial)
	}
	//fmt.Println("update from resp", len(o.Asks), len(o.Bids), len(orderBook.GetSnapshot().Orders))
}
