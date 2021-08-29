package coinbase

import (
	"context"
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/niccoloCastelli/orderbooks/common"
	"github.com/niccoloCastelli/orderbooks/exchanges"
	"github.com/niccoloCastelli/orderbooks/orderbook"
	"github.com/niccoloCastelli/orderbooks/utils"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"net/http"
)

const wsBaseUrl = "wss://Ws-feed.pro.coinbase.com"
const httpBaseUrl = "https://api.pro.coinbase.com"
const orderBookHttpUrl = httpBaseUrl + "/products/%s/book?level=3"
const name = "coinbase"

var availablePairs = []common.Pair{{"BTC", "USD"}, {"BTC", "EUR"}, {"XRP", "USD"}, {"XRP", "EUR"}, {"XRP", "BTC"}, {"LTC", "USD"}, {"LTC", "EUR"}, {"LTC", "BTC"}, {"ETH", "USD"}, {"ETH", "EUR"}, {"ETH", "BTC"}, {"BCH", "USD"}, {"BCH", "EUR"}, {"BCH", "BTC"}}

func init() {
	exchanges.RegisterExchange(NewCoinbaseClient())
}

type coinbaseClient struct {
	exchanges.Exchange
}

func NewCoinbaseClient() common.Exchange {
	return &coinbaseClient{*exchanges.NewExchange(name, availablePairs)}
}
func (b *coinbaseClient) Init(interface{}) error {
	b.Logger = log.Logger.With().Str("exchange", name).Logger()
	b.Lock.Lock()
	defer b.Lock.Unlock()
	b.Ws = map[string]*websocket.Conn{}
	b.OrderBooks = map[string]*orderbook.OrderBook{}
	resp, err := http.Get(fmt.Sprintf(orderBookHttpUrl, b.AvailablePairs()[0].String()))
	if err != nil {
		return err
	}
	return b.CheckResponse(resp)
}
func (b *coinbaseClient) SubscribeEvents(pair common.Pair) (context.Context, error) {
	c, _, err := websocket.DefaultDialer.Dial(wsBaseUrl, nil)
	if err != nil {
		return nil, err
	}
	pairLogger := b.Logger.With().Str("pair", pair.String()).Logger()
	wsCtx, cancelWs, err := b.CheckDisconnections(c, pairLogger)
	if err != nil {
		return nil, err
	}

	orderBook := b.NewOrderBook(pair, pairLogger)

	if err := c.WriteJSON(newSubscribeOrderBookMsg(pair)); err != nil {
		return nil, err
	}
	subscribeResp := map[string]interface{}{}
	if err := c.ReadJSON(&subscribeResp); err != nil {
		return nil, err
	}
	startSnapshot := orderBookResponse{}
	if err := c.ReadJSON(&startSnapshot); err != nil {
		return nil, err
	}
	startSnapshot.UpdateOrderBook(orderBook, true)

	go func(orderBook *orderbook.OrderBook, logger zerolog.Logger) {
		for {
			msg := orderBookUpdateMsg{}
			utils.TimeProfile("message read", func() {
				err = c.ReadJSON(&msg)
			})
			if err != nil {
				pairLogger.Err(err).Msg("read Ws resp error")
				cancelWs()
				return
			}
			utils.TimeProfile("order book update", func() {
				msg.UpdateOrderBook(orderBook, false)
			})
			//fmt.Println(orderBook.GetSnapshot())
			//orderBook.Update()
		}
	}(orderBook, pairLogger)
	b.Lock.Lock()
	b.Ws[pair.String()] = c
	b.OrderBooks[pair.String()] = orderBook
	b.Lock.Unlock()

	return wsCtx, nil
}
