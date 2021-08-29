package common

import (
	"fmt"
	"time"
)

type OrderSide int
type EventType int

const (
	OrderSideBid = 0
	OrderSideAsk = 1
)
const (
	EventTypeInit EventType = iota
	EventTypeAdd
	EventTypeChange
	EventTypeRemove
)

type Order struct {
	Timestamp time.Time
	Side      OrderSide
	Amount    float64
	Price     float64
}

func (o Order) Equal(order Order) bool {
	return o.Amount == order.Amount && o.Price == order.Price && o.Side == order.Side //o.Timestamp.Equal(order.Timestamp) &&
}

func (o *Order) Headers() []string {
	return []string{"timestamp", "side", "amount", "price"}
}
func (o *Order) Row() []string {
	return []string{
		fmt.Sprintf("%d", o.Timestamp.UnixNano()),
		fmt.Sprintf("%d", o.Side),
		fmt.Sprintf("%f", o.Amount),
		fmt.Sprintf("%f", o.Price),
	}
}

type Event struct {
	Order
	Type EventType
}

func (e *Event) Headers() []string {
	return append(e.Order.Headers(), "type")
}
func (e *Event) Row() []string {
	return append(e.Order.Row(), fmt.Sprintf("%d", e.Type))
}

type Snapshot struct {
	Timestamp time.Time
	Exchange  string
	Pair      Pair
	Orders    []Order
	Events    []Event
	SessionID int64
	Counter   int64
}

func (s Snapshot) Equal(snapshot Snapshot) bool {
	if len(s.Orders) != len(snapshot.Orders) {
		return false
	}
	for i, order := range s.Orders {
		if !order.Equal(snapshot.Orders[i]) {
			return false
		}
	}
	return true
}
