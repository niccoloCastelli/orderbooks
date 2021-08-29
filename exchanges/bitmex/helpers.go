package bitmex

import (
	"compress/flate"
	"encoding/json"
	"github.com/gorilla/websocket"
	"github.com/niccoloCastelli/orderbooks/common"
	"io/ioutil"
	"strings"
)

func readWsJson(c *websocket.Conn, o interface{}) error {
	_, msgReader, err := c.NextReader()
	if err != nil {
		return err
	}
	rawMsg, err := ioutil.ReadAll(flate.NewReader(msgReader))
	return json.Unmarshal(rawMsg, o)
}

func formatPair(pair common.Pair) string {
	base := pair.Base
	quote := pair.Quote
	if base == "BTC" {
		base = "XBT"
	}
	if quote == "BTC" {
		quote = "XBT"
	}
	return strings.ToUpper(base) + strings.ToUpper(quote)
}

func priceFromID(id int64, idx int64, instrument *InstrumentInfo) float64 {
	if instrument.Symbol == "XBTUSD" {
		return float64(100000000*idx-id) * 0.01
	}
	return float64(100000000*idx-id) * instrument.TickSize
}
