package server

import (
	"github.com/niccoloCastelli/orderbooks/common"
	orderbooks "github.com/niccoloCastelli/orderbooks/server/proto"
	"github.com/niccoloCastelli/orderbooks/utils"
)

func SnapshotToPb(in *common.Snapshot) *orderbooks.SnapshotMsg {
	if in == nil {
		return nil
	}
	pbOrders := make([]*orderbooks.Event, len(in.Orders))
	for i, order := range in.Orders {
		pbOrders[i] = &orderbooks.Event{
			Timestamp: order.Timestamp.UnixNano(),
			OrderSide: orderbooks.OrderSide(order.Side),
			Amount:    float32(order.Amount),
			Price:     float32(order.Price),
		}
	}
	return &orderbooks.SnapshotMsg{
		Timestamp: utils.TimeToGogoProtoTsPtr(&in.Timestamp),
		Exchange:  in.Exchange,
		Pair:      in.Pair.String(),
		Orders:    pbOrders,
		Events:    nil,
		SessionId: in.SessionID,
		Counter:   in.Counter,
	}
}
