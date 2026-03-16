package adapter

import "github.com/tc1io/clean-arch-demo/ent"

type OrderQueue struct{}

func (q OrderQueue) PlaceOrder(order ent.Order) (string, error) {
	return "", nil
}
