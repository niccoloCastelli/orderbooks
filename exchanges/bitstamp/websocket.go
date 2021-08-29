package bitstamp

import (
	"fmt"
	"github.com/niccoloCastelli/orderbooks/common"
	"strings"
)

type wsEvent string

const (
	wsEventSubscribe = "bts:subscribe"

	wsChanDiffOrderBook = "diff_order_book"
)

type wsChannelData struct {
	Channel string `json:"channel"`
}

//{"event":"bts:subscribe","data":{"channel":"diff_order_book_btcusd"}}
type websocketSubscription struct {
	Event wsEvent       `json:"event"`
	Data  wsChannelData `json:"data"`
}

func subscribeDiffOrderBook(pair common.Pair) *websocketSubscription {
	pairName := strings.ToLower(fmt.Sprintf("%s%s", pair.Base, pair.Quote))
	return &websocketSubscription{
		Event: wsEventSubscribe,
		Data:  wsChannelData{fmt.Sprintf("%s_%s", wsChanDiffOrderBook, pairName)},
	}
}

type orderBookDiffMsg struct {
	Data    orderBookResponse `json:"data"`
	Event   string            `json:"event"`
	Channel string            `json:"channel"`
}
