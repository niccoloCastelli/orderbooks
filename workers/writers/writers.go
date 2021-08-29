package writers

import (
	"github.com/niccoloCastelli/orderbooks/common"
	"strings"
)

func UniqueName(prefix string, pair common.Pair, exchange string) string {
	return strings.ToLower(strings.TrimPrefix(strings.Join([]string{prefix, pair.String(), exchange}, "_"), "_"))
}
