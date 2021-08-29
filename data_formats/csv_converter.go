package data_formats

import "github.com/niccoloCastelli/orderbooks/common"

type ObCsvConverter interface {
	ParseRow([]string) error
	ToOrder() common.Order
	ToEvent() common.Event
}
