package kraken

import (
	"encoding/json"
	"errors"
	"github.com/niccoloCastelli/orderbooks/common"
	"github.com/niccoloCastelli/orderbooks/orderbook"
	"github.com/rs/zerolog/log"
	"time"
)

const (
	wsChanOrderBook = "book"
	orderBookDepth  = 1000
)

type wsChannel struct {
	Name  string `json:"name"`
	Depth int    `json:"depth"`
}
type wsEvent struct {
	Event string `json:"event"`
}

func newSubscribeOrderBookMsg(pairs ...common.Pair) *subscribeOrderBookMsg {
	formattedPairs := make([]string, 0, len(pairs))
	for _, pair := range pairs {
		formattedPairs = append(formattedPairs, formatPair(pair, "/"))
	}
	channel := wsChannel{
		Name:  wsChanOrderBook,
		Depth: orderBookDepth,
	}
	return &subscribeOrderBookMsg{
		Event:        "subscribe",
		Pair:         formattedPairs,
		Subscription: channel,
	}
}

type subscribeOrderBookMsg struct {
	Event        string    `json:"event"`
	Pair         []string  `json:"pair"`
	Subscription wsChannel `json:"subscription"`
}
type subscriptionResponse struct {
	ChannelID    int    `json:"channelID"`
	ChannelName  string `json:"channelName"`
	Event        string `json:"event"`
	Pair         string `json:"pair"`
	Status       string `json:"status"`
	Subscription struct {
		Depth int    `json:"depth"`
		Name  string `json:"name"`
	} `json:"subscription"`
}
type asksBidsMsg struct {
	As [][]string `json:"as"`
	Bs [][]string `json:"bs"`
}
type asksBidsUpdateMsg struct {
	Asks [][]string `json:"a"`
	Bids [][]string `json:"b"`
}
type orderBookUpdateMsg struct {
	initSnapshot bool
	ChannelID    int
	ChannelName  string
	ChannelPair  string
	Time         time.Time
	Asks         [][]string
	Bids         [][]string
}

func (o *orderBookUpdateMsg) UnmarshalJSON(bytes []byte) error {
	var respArr []json.RawMessage

	if err := json.Unmarshal(bytes, &respArr); err != nil {
		event := wsEvent{}
		if err := json.Unmarshal(bytes, &event); err != nil {
			log.Err(err).Str("event", event.Event).Str("raw", string(bytes)).Send()
			return err
		}
		if event.Event != "heartbeat" {
			log.Err(err).Str("event", event.Event).Send()
			return err
		}
		//log.Debug().Err(err).Str("event", event.Event).Send()
		o.Asks = [][]string{}
		o.Bids = [][]string{}
		o.Time = time.Now()
		return nil
	}
	if len(respArr) < 4 {
		return errors.New("min expected length: 4")
	}
	if err := json.Unmarshal(respArr[0], &o.ChannelID); err != nil {
		return err
	}
	if o.initSnapshot {
		asksBids := asksBidsMsg{}
		if err := json.Unmarshal(respArr[1], &asksBids); err != nil {
			return err
		}
		o.Asks = asksBids.As
		o.Bids = asksBids.Bs
	} else {
		asksBids := asksBidsUpdateMsg{}
		if err := json.Unmarshal(respArr[1], &asksBids); err != nil {
			return err
		}
		if len(respArr) > 4 {
			if err := json.Unmarshal(respArr[2], &asksBids); err != nil {
				return err
			}
		}

		o.Asks = asksBids.Asks
		o.Bids = asksBids.Bids
	}

	if err := json.Unmarshal(respArr[len(respArr)-2], &o.ChannelName); err != nil {
		return err
	}
	if err := json.Unmarshal(respArr[len(respArr)-1], &o.ChannelPair); err != nil {
		return err
	}

	o.Time = time.Now()
	return nil
}
func (o orderBookUpdateMsg) UpdateOrderBook(orderBook *orderbook.OrderBook, initial bool) {
	for _, newOrder := range o.Bids {
		order := common.Order{
			Timestamp: time.Time{},
			Side:      common.OrderSideBid,
		}
		if err := rowToOrder(newOrder, &order); err != nil {
			continue
		}
		orderBook.Update(order, initial)
	}
	for _, newOrder := range o.Asks {
		order := common.Order{
			Timestamp: time.Time{},
			Side:      common.OrderSideAsk,
		}
		if err := rowToOrder(newOrder, &order); err != nil {
			continue
		}
		orderBook.Update(order, initial)
	}
	//fmt.Println("update from resp", len(o.Asks), len(o.Bids), len(orderBook.GetSnapshot().Orders))
}
