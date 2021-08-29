package ob_format

import (
	"errors"
	"github.com/niccoloCastelli/orderbooks/common"
	"github.com/niccoloCastelli/orderbooks/data_formats"
	"strconv"
	"time"
)

const (
	ErrWrongColumnNumber = "expected length: 4 or 5 columns"
)

func orderSideFromStr(in string) (common.OrderSide, error) {
	switch in {
	case "0":
		return common.OrderSideBid, nil
	case "1":
		return common.OrderSideAsk, nil
	default:
		return -1, errors.New("invalid order side value")
	}
}
func parseUpdateType(in string) (common.EventType, error) {
	switch in {
	case "0":
		return common.EventTypeInit, nil
	case "1":
		return common.EventTypeAdd, nil
	case "2":
		return common.EventTypeChange, nil
	case "3":
		return common.EventTypeRemove, nil
	default:
		return 0, errors.New("invalid event type value")
	}
}

type ObRow struct {
	Timestamp  time.Time
	UpdateType common.EventType
	Side       common.OrderSide
	Amount     float64
	Price      float64
}

func ParseObRow(row []string) (data_formats.ObCsvConverter, error) {
	r := &ObRow{}
	if err := r.ParseRow(row); err != nil {
		return nil, err
	}
	return r, nil
}
func (c *ObRow) ParseRow(row []string) (err error) {
	if len(row) != 4 && len(row) != 5 {
		return errors.New(ErrWrongColumnNumber)
	}
	msInt, err := strconv.ParseInt(row[0], 10, 64)
	if err != nil {
		return err
	}
	c.Timestamp = time.Unix(0, msInt)
	/*if c.UpdateType, err = parseUpdateType(row[2]); err != nil {
		return err
	}*/
	if c.Side, err = orderSideFromStr(row[1]); err != nil {
		return err
	}
	if c.Amount, err = strconv.ParseFloat(row[2], 64); err != nil {
		return err
	}
	if c.Price, err = strconv.ParseFloat(row[3], 64); err != nil {
		return err
	}
	if len(row) == 5 {
		if c.UpdateType, err = parseUpdateType(row[4]); err != nil {
			return err
		}
	}
	//timestamp,side,amount,price,type
	/*if c.UpdateType == UpdateModeSub {
		c.Amount = -c.Amount
	}*/
	return nil
}

func (c *ObRow) ToOrder() common.Order {
	return common.Order{
		Timestamp: c.Timestamp,
		Side:      c.Side,
		Amount:    c.Amount,
		Price:     c.Price,
	}
}
func (c *ObRow) ToEvent() common.Event {
	order := c.ToOrder()
	switch c.UpdateType {
	case common.EventTypeRemove:
		order.Amount = 0
	}
	return common.Event{
		Order: order,
		Type:  c.UpdateType,
	}
}
