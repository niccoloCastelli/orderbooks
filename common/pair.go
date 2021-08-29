package common

import (
	"errors"
	"strings"
)

const (
	ErrInvalidPair = "invalid pair"
)

type Pair struct {
	Base  string
	Quote string
}

func (p Pair) String() string {
	return p.Base + "-" + p.Quote
}

func PairsFromString(pairStr string) (*Pair, error) {
	spl := strings.Split(pairStr, "-")
	if len(spl) != 2 {
		return nil, errors.New(ErrInvalidPair)
	}
	return &Pair{Base: spl[0], Quote: spl[1]}, nil
}
func PairsFromStrings(pairStrs ...string) ([]Pair, error) {
	if pairStrs == nil || len(pairStrs) == 0 {
		return []Pair{}, nil
	}
	ret := make([]Pair, len(pairStrs))
	for i, pairStr := range pairStrs {
		pair, err := PairsFromString(pairStr)
		if err != nil {
			return nil, err
		}
		ret[i] = *pair
	}
	return ret, nil
}
