package coinbase

import (
	"github.com/niccoloCastelli/orderbooks/common"
	"strconv"
)

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

func sidedRowToOrder(row []string, o *common.Order) error {
	var amount, price float64
	var err error
	if row[0] == "buy" {
		o.Side = common.OrderSideBid
	} else {
		o.Side = common.OrderSideAsk
	}
	if price, err = strconv.ParseFloat(row[1], 64); err != nil {
		return err
	}
	if amount, err = strconv.ParseFloat(row[2], 64); err != nil {
		return err
	}
	o.Price = price
	o.Amount = amount
	return nil
}
