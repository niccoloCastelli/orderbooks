package okex

import (
	"compress/flate"
	"encoding/json"
	"github.com/gorilla/websocket"
	"github.com/niccoloCastelli/orderbooks/common"
	"io/ioutil"
	"strconv"
)

func readWsJson(c *websocket.Conn, o interface{}) error {
	_, msgReader, err := c.NextReader()
	if err != nil {
		return err
	}
	rawMsg, err := ioutil.ReadAll(flate.NewReader(msgReader))
	return json.Unmarshal(rawMsg, o)
}

func rowToOrder(row []string, o *common.Order) error {
	var amount, price float64
	var err error
	if price, err = strconv.ParseFloat(row[0], 64); err != nil {
		return err
	}
	if amount, err = strconv.ParseFloat(row[1], 64); err != nil {
		return err
	}
	o.Price = price
	o.Amount = amount
	return nil
}
