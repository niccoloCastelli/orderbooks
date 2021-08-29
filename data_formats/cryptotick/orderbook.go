package cryptotick

import (
	"errors"
	"github.com/niccoloCastelli/orderbooks/common"
	"strconv"
	"time"
)

type UpdateType string

const (
	timeLayout                    = "15:04:05.0000000"
	UpdateModeSnapshot UpdateType = "SNAPSHOT"
	UpdateModeAdd      UpdateType = "ADD"
	UpdateModeSub      UpdateType = "SUB"
)

func orderSideFromStr(in string) (common.OrderSide, error) {
	switch in {
	case "0":
		return common.OrderSideAsk, nil
	case "1":
		return common.OrderSideBid, nil
	default:
		return -1, errors.New("invalid order side value")
	}
}
func parseUpdateType(in string) (UpdateType, error) {
	switch in {
	case string(UpdateModeSnapshot), string(UpdateModeAdd), string(UpdateModeSub):
		return UpdateType(in), nil
	default:
		return "", errors.New("invalid update type value")
	}
}

type CryptotyickRow struct {
	TimeExchange time.Time
	TimeCoinapi  time.Time
	UpdateType   UpdateType
	Side         common.OrderSide
	Amount       float64
	Price        float64
}

func ParseCryptotyickRow(row []string) (*CryptotyickRow, error) {
	r := &CryptotyickRow{}
	if err := r.ParseRow(row); err != nil {
		return nil, err
	}
	return r, nil
}
func (c *CryptotyickRow) ParseRow(row []string) (err error) {
	if len(row) != 6 {
		return errors.New("expected length: 6 columns")
	}
	if c.TimeExchange, err = time.Parse(timeLayout, row[0]); err != nil {
		return err
	}
	c.TimeExchange = c.TimeExchange
	if c.TimeCoinapi, err = time.Parse(timeLayout, row[1]); err != nil {
		return err
	}
	if c.UpdateType, err = parseUpdateType(row[2]); err != nil {
		return err
	}
	if c.Side, err = orderSideFromStr(row[3]); err != nil {
		return err
	}
	if c.Price, err = strconv.ParseFloat(row[4], 64); err != nil {
		return err
	}
	if c.Amount, err = strconv.ParseFloat(row[5], 64); err != nil {
		return err
	}
	if c.UpdateType == UpdateModeSub {
		c.Amount = -c.Amount
	}
	return nil
}

func (c *CryptotyickRow) ToOrder() common.Order {
	return common.Order{
		Timestamp: c.TimeExchange,
		Side:      c.Side,
		Amount:    c.Amount,
		Price:     c.Price,
	}
}
