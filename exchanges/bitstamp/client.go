package bitstamp

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/niccoloCastelli/orderbooks/common"
	"github.com/niccoloCastelli/orderbooks/exchanges"
	"github.com/niccoloCastelli/orderbooks/orderbook"
	"github.com/niccoloCastelli/orderbooks/utils"
	"github.com/pkg/errors"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"io/ioutil"
	"net/http"
	"strings"
)

const wsBaseUrl = "wss://Ws.bitstamp.net"
const httpBaseUrl = "https://www.bitstamp.net"
const orderBookHttpUrl = httpBaseUrl + "/api/v2/order_book"
const name = "bitstamp"

var availablePairs = []common.Pair{{"BTC", "USD"}, {"BTC", "EUR"}, {"EUR", "USD"}, {"XRP", "USD"}, {"XRP", "EUR"}, {"XRP", "BTC"}, {"LTC", "USD"}, {"LTC", "EUR"}, {"LTC", "BTC"}, {"ETH", "USD"}, {"ETH", "EUR"}, {"ETH", "BTC"}, {"BCH", "USD"}, {"BCH", "EUR"}, {"BCH", "BTC"}}

func init() {
	exchanges.RegisterExchange(NewBitstampClient())
}

type bitstampClient struct {
	exchanges.Exchange
}

func NewBitstampClient() common.Exchange {
	return &bitstampClient{*exchanges.NewExchange(name, availablePairs)}
}
func (b *bitstampClient) formatPair(pair common.Pair) string {
	return strings.ToLower(pair.Base) + strings.ToLower(pair.Quote)
}
func (b *bitstampClient) getOrderBookSnapshot(pair common.Pair) (*orderBookResponse, error) {
	resp, err := http.Get(orderBookHttpUrl + "/" + b.formatPair(pair))
	if err != nil {
		return nil, errors.WithStack(err)
	}
	rawData := []byte{}
	defer resp.Body.Close()
	if rawData, err = ioutil.ReadAll(resp.Body); err != nil {
		return nil, errors.WithStack(err)
	}
	if resp.StatusCode != http.StatusOK {
		return nil, errors.New(fmt.Sprintf("%d - %s", resp.StatusCode, string(rawData)))
	}
	obResp := orderBookResponse{}
	if err := json.Unmarshal(rawData, &obResp); err != nil {
		return nil, errors.WithStack(err)
	}
	return &obResp, nil
}

func (b *bitstampClient) Init(interface{}) error {
	b.Logger = log.Logger.With().Str("exchange", name).Logger()
	b.Lock.Lock()
	defer b.Lock.Unlock()
	b.Ws = map[string]*websocket.Conn{}
	b.OrderBooks = map[string]*orderbook.OrderBook{}
	resp, err := http.Get(orderBookHttpUrl + "/" + b.formatPair(common.Pair{Base: "BTC", Quote: "USD"}))
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusOK {
		respData := []byte{}
		if respData, err = ioutil.ReadAll(resp.Body); err != nil {
			return errors.WithStack(err)
		}
		return errors.WithStack(errors.New(fmt.Sprintf("%d - %s", resp.StatusCode, string(respData))))
	}
	return nil
}
func (b *bitstampClient) SubscribeEvents(pair common.Pair) (context.Context, error) {
	c, _, err := websocket.DefaultDialer.Dial(wsBaseUrl, nil)
	if err != nil {
		return nil, err
	}
	pairLogger := b.Logger.With().Str("pair", pair.String()).Logger()

	orderBook := b.NewOrderBook(pair, pairLogger)

	wsCtx, cancelWs, err := b.CheckDisconnections(c, pairLogger)
	if err != nil {
		return nil, err
	}

	startSnapshot, err := b.getOrderBookSnapshot(pair)
	if err != nil {
		return nil, err
	}
	startSnapshot.UpdateOrderBook(orderBook, true)
	go func(orderBook *orderbook.OrderBook, logger zerolog.Logger) {
		for {
			msg := orderBookDiffMsg{}
			utils.TimeProfile("message read", func() {
				err = c.ReadJSON(&msg)
			})
			if err != nil {
				pairLogger.Err(err).Msg("read Ws resp error")
				cancelWs()
				return
			}
			utils.TimeProfile("order book update", func() {
				msg.Data.UpdateOrderBook(orderBook, false)
			})
			//fmt.Println(orderBook.GetSnapshot())
			//orderBook.Update()
		}
	}(orderBook, pairLogger)
	b.Lock.Lock()
	b.Ws[pair.String()] = c
	b.OrderBooks[pair.String()] = orderBook
	b.Lock.Unlock()
	if err := c.WriteJSON(subscribeDiffOrderBook(pair)); err != nil {
		return nil, err
	}

	return wsCtx, nil
}
