package kraken

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/niccoloCastelli/orderbooks/common"
	"github.com/niccoloCastelli/orderbooks/exchanges"
	"github.com/niccoloCastelli/orderbooks/orderbook"
	"github.com/niccoloCastelli/orderbooks/utils"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"net/http"
	"time"
)

const wsBaseUrl = "wss://ws.kraken.com"
const httpBaseUrl = "https://api.kraken.com"
const orderBookHttpUrl = httpBaseUrl + "/0/public/Depth?pair=%s"
const name = "kraken"

var availablePairs = []common.Pair{{"BTC", "USD"}, {"BTC", "EUR"}, {"XRP", "USD"}, {"XRP", "EUR"}, {"XRP", "BTC"}, {"LTC", "USD"}, {"LTC", "EUR"}, {"LTC", "BTC"}, {"ETH", "USD"}, {"ETH", "EUR"}, {"ETH", "BTC"}, {"BCH", "USD"}, {"BCH", "EUR"}, {"BCH", "BTC"}}

//Bitcoin: XBT
func init() {
	exchanges.RegisterExchange(NewKrakenClient())
}

type krakenClient struct {
	exchanges.Exchange
}

func NewKrakenClient() common.Exchange {
	return &krakenClient{*exchanges.NewExchange(name, availablePairs)}
}
func (b *krakenClient) formatPair(pair common.Pair) string {
	return formatPair(pair, "")
}
func (b *krakenClient) Init(interface{}) error {
	b.Logger = log.Logger.With().Str("exchange", name).Logger()
	b.Lock.Lock()
	defer b.Lock.Unlock()
	b.Ws = map[string]*websocket.Conn{}
	b.OrderBooks = map[string]*orderbook.OrderBook{}
	resp, err := http.Get(fmt.Sprintf(orderBookHttpUrl, b.formatPair(b.AvailablePairs()[0])))
	if err != nil {
		return err
	}
	return b.CheckResponse(resp)
}
func (b *krakenClient) SubscribeEvents(pair common.Pair) (context.Context, error) {
	websocketDialer := websocket.Dialer{
		HandshakeTimeout: time.Second * 10,
		ReadBufferSize:   1024 * 16,
		WriteBufferSize:  1024 * 16,
	}
	c, _, err := websocketDialer.Dial(wsBaseUrl, nil)
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
	connectionResp := map[string]interface{}{}
	if err := c.ReadJSON(&connectionResp); err != nil {
		return nil, err
	}
	subscribeResp := subscriptionResponse{}
	if err := c.ReadJSON(&subscribeResp); err != nil {
		return nil, err
	}
	startSnapshot := orderBookUpdateMsg{initSnapshot: true}
	if err := c.ReadJSON(&startSnapshot); err != nil {
		return nil, err
	}
	startSnapshot.UpdateOrderBook(orderBook, true)
	go func(orderBook *orderbook.OrderBook, logger zerolog.Logger) {
		for {
			msg := orderBookUpdateMsg{}
			buf := []byte{}
			utils.TimeProfile("message read", func() {
				_, buf, err = c.ReadMessage()
			})
			if err != nil {
				pairLogger.Err(err).Msg("read Ws resp error")
				cancelWs()
				return
			}
			if err := json.Unmarshal(buf, &msg); err != nil {
				pairLogger.Err(err).Str("raw", string(buf)).Msg("json decode error")
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
