package okex

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

const wsBaseUrl = "wss://real.okex.com:8443/ws/v3"
const httpBaseUrl = "https://www.okex.com/api"
const orderBookHttpUrl = httpBaseUrl + "/spot/v3/instruments/%s/book?size=200"
const name = "okex"

var availablePairs = []common.Pair{{"BTC", "USDT"}, {"XRP", "USDT"}, {"XRP", "BTC"}, {"LTC", "BTC"}, {"ETH", "USDT"}, {"ETH", "BTC"}, {"BCH", "BTC"}}

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
	url := fmt.Sprintf(orderBookHttpUrl, b.AvailablePairs()[0].String())
	b.Logger.Debug().Str("url", url).Msg("order book url")
	resp, err := http.Get(url)
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
	c.EnableWriteCompression(true)
	pairLogger := b.Logger.With().Str("pair", pair.String()).Logger()
	wsCtx, cancelWs, err := b.CheckDisconnections(c, pairLogger)
	if err != nil {
		return nil, err
	}

	orderBook := b.NewOrderBook(pair, pairLogger)

	if err := c.WriteJSON(newSubscribeOrderBookMsg(pair)); err != nil {
		return nil, err
	}
	//subscribeResp := map[string]interface{}{}
	sr := subscribeResp{}
	if err := readWsJson(c, &sr); err != nil {
		return nil, err
	}

	pairLogger.Debug().Str("event", sr.Event).Str("channel", sr.Channel).Msg("subscribe")
	startSnapshot := orderBookUpdateMsg{}
	if err := readWsJson(c, &startSnapshot); err != nil {
		return nil, err
	}
	startSnapshot.UpdateOrderBook(orderBook, true)
	go func(orderBook *orderbook.OrderBook, logger zerolog.Logger) {
		for {
			msg := orderBookUpdateMsg{}
			utils.TimeProfile("message read", func() {
				err = readWsJson(c, &msg)
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
