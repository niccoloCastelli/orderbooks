package bitmex

import (
	"context"
	"errors"
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

const wsBaseUrl = "wss://www.bitmex.com/realtime"
const httpBaseUrl = "https://www.bitmex.com/api/v1"
const orderBookHttpUrl = httpBaseUrl + "/orderBook/L2?symbol=%s"
const instrumentsUrl = httpBaseUrl + "/instrument?columns=symbol,tickSize&start=0&count=500"
const name = "bitmex"

var availablePairs = []common.Pair{{"BTC", "USD"}, {"XRP", "USD"}, {"XRP", "BTC"}, {"LTC", "BTC"}, {"ETH", "USD"}, {"ETH", "BTC"}, {"BCH", "BTC"}}

func init() {
	exchanges.RegisterExchange(NewBitmexClient())
}

type bitmexClient struct {
	exchanges.Exchange
	instruments Instruments
}

func NewBitmexClient() common.Exchange {
	return &bitmexClient{*exchanges.NewExchange(name, availablePairs), Instruments{}}
}
func (b *bitmexClient) Init(interface{}) error {
	b.Logger = log.Logger.With().Str("exchange", name).Logger()
	b.Lock.Lock()
	defer b.Lock.Unlock()
	b.Ws = map[string]*websocket.Conn{}
	b.OrderBooks = map[string]*orderbook.OrderBook{}

	resp, err := http.Get(instrumentsUrl)
	if err != nil {
		return err
	}
	if err := b.CheckResponse(resp); err != nil {
		return err
	}
	if err := utils.UnmarshalResponse(resp, &b.instruments); err != nil {
		return err
	}

	return nil
}
func (b *bitmexClient) SubscribeEvents(pair common.Pair) (context.Context, error) {
	websocketDialer := websocket.Dialer{
		HandshakeTimeout: time.Second * 30,
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
	//SubscribeResp := map[string]interface{}{}
	_, msg, _ := c.ReadMessage()
	pairLogger.Debug().Str("msg", string(msg)).Msg("subscribe")
	sr := SubscribeResp{}
	if err := c.ReadJSON(&sr); err != nil {
		return nil, err
	}

	pairLogger.Debug().Bool("success", sr.Success).Str("channel", sr.Subscribe).Msg("subscribe")
	/*if sr.Success == false {
		return errors.New("subscribe unsuccessful")
	}*/
	instrumentIdx, instrument := b.instruments.findBySymbol(formatPair(pair))
	if instrument == nil {
		return nil, errors.New("invalid pair")
	}
	startSnapshot := orderBookUpdateMsg{}
	if err := c.ReadJSON(&startSnapshot); err != nil {
		return nil, err
	}
	startSnapshot.UpdateOrderBook(orderBook, instrumentIdx, *instrument, true)

	go func(orderBook *orderbook.OrderBook, logger zerolog.Logger, instrument InstrumentInfo) {
		for {
			msg := orderBookUpdateMsg{}
			utils.TimeProfile("message read", func() {
				err = c.ReadJSON(&msg)
			})
			if err != nil {
				if e, ok := err.(*websocket.CloseError); ok {
					pairLogger.Err(err).Int("code", e.Code).Msg("Ws close error")
				} else {
					pairLogger.Err(err).Msg("read Ws resp error")
				}
				cancelWs()
				return
			}
			utils.TimeProfile("order book update", func() {
				msg.UpdateOrderBook(orderBook, instrumentIdx, instrument, false)
			})
			//fmt.Println(orderBook.GetSnapshot())
			//orderBook.Update()
		}
	}(orderBook, pairLogger, *instrument)

	b.Lock.Lock()
	b.Ws[pair.String()] = c
	b.OrderBooks[pair.String()] = orderBook
	b.Lock.Unlock()

	return wsCtx, nil
}
