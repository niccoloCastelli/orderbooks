package exchanges

import "github.com/niccoloCastelli/orderbooks/common"

var registeredExchanges = map[string]common.Exchange{}

func RegisterExchange(exchange common.Exchange) {
	registeredExchanges[exchange.Name()] = exchange
}

func GetExchanges(enabledExchanges ...string) []common.Exchange {
	addAll := enabledExchanges == nil || len(enabledExchanges) == 0
	ret := make([]common.Exchange, 0, len(registeredExchanges))
	for key, _ := range registeredExchanges {
		if addAll {
			ret = append(ret, registeredExchanges[key])
		} else {
			for _, ex := range enabledExchanges {
				if key == ex {
					ret = append(ret, registeredExchanges[key])
					break
				}
			}
		}

	}
	return ret
}

func GetExchange(name string) (common.Exchange, bool) {
	ex, ok := registeredExchanges[name]
	return ex, ok
}
