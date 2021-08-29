package orderbook

import (
	"github.com/niccoloCastelli/orderbooks/common"
	"github.com/niccoloCastelli/orderbooks/utils"
	"github.com/rs/zerolog"
)

func reverse(a []*limit) []*limit {
	for i := len(a)/2 - 1; i >= 0; i-- {
		opp := len(a) - 1 - i
		a[i], a[opp] = a[opp], a[i]
	}
	return a
}

type nodeTree struct {
	head          *limit
	tail          *limit
	reverse       bool
	addMode       bool
	logger        zerolog.Logger
	registerEvent func(price float64, amount float64, eventType common.EventType)
}

func (t *nodeTree) addBefore(node *limit, newNode *limit) {
	if node.Left != nil {
		node.Left.Right = newNode
	}
	newNode.Right = node
	newNode.Left = node.Left
	node.Left = newNode
	if node == t.head {
		t.logger.Trace().Msgf("head changed: %v -> %v", t.head, newNode)
		t.head = newNode
	}
	if t.registerEvent != nil {
		t.registerEvent(node.Price, node.Amount, common.EventTypeAdd)
	}
}
func (t *nodeTree) addAfter(node *limit, newNode *limit) {
	newNode.Left = node
	newNode.Right = node.Right
	if node.Right != nil {
		node.Right.Left = newNode
	}
	node.Right = newNode
	if node == t.tail {
		t.logger.Trace().Msgf("tail changed: %v -> %v", t.head, newNode)
		t.tail = newNode
	}
	if t.registerEvent != nil {
		t.registerEvent(node.Price, node.Amount, common.EventTypeAdd)
	}
}
func (t *nodeTree) removeNode(node *limit) {
	if t.head == nil {
		return
	}
	if node.Left != nil {
		node.Left.Right = node.Right
	}
	if node.Right != nil {
		node.Right.Left = node.Left
	}
	if node == t.head {
		t.logger.Trace().Msgf("head changed: %v -> %v", t.head, t.head.Right)
		t.head = t.head.Right
		if t.head != nil {
			if utils.FloatEquals(t.head.Amount, 0) {
				t.removeNode(t.head)
			}
		}
	}
	if node == t.tail {
		t.logger.Trace().Msgf("tail changed: %v -> %v", t.tail, t.tail.Left)
		t.tail = t.tail.Left
		if t.tail != nil {
			if utils.FloatEquals(t.head.Amount, 0) {
				t.removeNode(t.tail)
			}
		}
	}
	if t.registerEvent != nil {
		t.registerEvent(node.Price, node.Amount, common.EventTypeRemove)
	}
}
func (t *nodeTree) UpdateTree(amount float64, price float64) {
	if t.head == nil && !utils.FloatEquals(amount, 0) {
		t.head = &limit{
			Amount: amount,
			Price:  price,
		}
		t.tail = t.head
		if t.registerEvent != nil {
			t.registerEvent(0, 0, common.EventTypeInit)
			t.registerEvent(price, amount, common.EventTypeAdd)
		}
		return
	}
	if t.head == nil && utils.FloatEquals(amount, 0) {
		return
	}
	l := t.head
	if t.reverse {
		l = t.tail
	}
loop:
	for {
		if utils.FloatEquals(price, l.Price) {
			if utils.FloatEquals(amount, 0) {
				t.removeNode(l)
			} else {
				if t.addMode {
					l.Amount += amount
					if l.Amount <= 0 {
						t.removeNode(l)
					}
				} else {
					l.Amount = amount
				}
				t.registerEvent(price, l.Amount, common.EventTypeChange)
			}
			break loop
		} else if price > l.Price {
			if l.Right == nil || price < l.Right.Price {
				if utils.FloatEquals(amount, 0) {
					break loop
				}
				t.addAfter(l, &limit{Amount: amount, Price: price})
				break loop
			} else {
				l = l.Right
				continue loop
			}
		} else if price < l.Price {
			if l.Left == nil || price > l.Left.Price {
				if utils.FloatEquals(amount, 0) {
					break loop
				}
				t.addBefore(l, &limit{Amount: amount, Price: price})
				break loop
			} else {
				l = l.Left
				continue loop
			}
		}
	}
}
func (t *nodeTree) GenerateIndex(sizeLimit int) []*limit {
	if t.head == nil || t.tail == nil {
		return []*limit{}
	}
	var ret []*limit
	sizeLimit--
	if t.reverse {
		ret = []*limit{t.tail}
		node := t.tail
		for i := 0; ; i++ {
			if node.Left == nil {
				break
			}
			if sizeLimit > 0 && i >= sizeLimit {
				break
			}
			ret = append(ret, node.Left)
			node = node.Left
		}
		reverse(ret)
	} else {
		ret = []*limit{t.head}
		node := t.head
		for i := 0; ; i++ {
			if node.Right == nil {
				break
			}
			if sizeLimit > 0 && i >= sizeLimit {
				break
			}
			ret = append(ret, node.Right)
			node = node.Right
		}
	}

	return ret
}
