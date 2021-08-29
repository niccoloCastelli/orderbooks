package bitmex

import (
	"time"
)

type InstrumentInfo struct {
	Symbol    string    `json:"symbol"`
	Timestamp time.Time `json:"timestamp"`
	TickSize  float64   `json:"tickSize"`
}
type Instruments []InstrumentInfo

func (i Instruments) findBySymbol(symbol string) (int, *InstrumentInfo) {
	for idx, instr := range i {
		if instr.Symbol == symbol {
			return idx, &instr
		}
	}
	return -1, nil
}
func (i Instruments) findByIndex(index int) (int, *InstrumentInfo) {
	for idx, instr := range i {
		if idx == index {
			return idx, &instr
		}
	}
	return -1, nil
}

func getInstruments() {

}
