package okex

import (
	"fmt"
	"github.com/niccoloCastelli/orderbooks/common"
	"github.com/niccoloCastelli/orderbooks/orderbook"
	"time"
)

const (
	wsChanOrderBook = "spot/depth_l2_tbt:%s"
)

type wsChannel struct {
	Name       string   `json:"name"`
	ProductIds []string `json:"product_ids"`
}

//{"op": "subscribe", "args": ["spot/depth_l2_tbt:BTC-USDT"]}
type subscribeOrderBookMsg struct {
	Op   string   `json:"op"`
	Args []string `json:"args"`
}

type subscribeResp struct {
	Event   string `json:"event"`
	Channel string `json:"channel"`
}

func newSubscribeOrderBookMsg(pairs ...common.Pair) *subscribeOrderBookMsg {
	pairChans := make([]string, 0, len(pairs))
	for _, pair := range pairs {
		pairChans = append(pairChans, fmt.Sprintf(wsChanOrderBook, pair.String()))
	}
	return &subscribeOrderBookMsg{
		Op:   "subscribe",
		Args: pairChans,
	}
}

type orderBookUpdateMsg struct {
	Table  string `json:"table"`
	Action string `json:"action"`
	Data   []struct {
		InstrumentID string     `json:"instrument_id"`
		Asks         [][]string `json:"asks"`
		Bids         [][]string `json:"bids"`
		Timestamp    time.Time  `json:"timestamp"`
		Checksum     int        `json:"checksum"`
	} `json:"data"`
}

func (o orderBookUpdateMsg) UpdateOrderBook(orderBook *orderbook.OrderBook, initial bool) {
	for _, data := range o.Data {
		for _, newOrder := range data.Bids {
			order := common.Order{
				Timestamp: data.Timestamp,
				Side:      common.OrderSideBid,
			}
			if err := rowToOrder(newOrder, &order); err != nil {
				continue
			}
			orderBook.Update(order, initial)
		}
		for _, newOrder := range data.Asks {
			order := common.Order{
				Timestamp: data.Timestamp,
				Side:      common.OrderSideAsk,
			}
			if err := rowToOrder(newOrder, &order); err != nil {
				continue
			}
			orderBook.Update(order, initial)
		}
	}
	//fmt.Println("update from resp", len(o.Asks), len(o.Bids), len(orderBook.GetSnapshot().Orders))
}
