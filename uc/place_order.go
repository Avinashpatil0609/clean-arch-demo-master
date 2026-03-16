package uc

import (
	"github.com/tc1io/clean-arch-demo/ent"
)

type OrderPlacer interface {
	PlaceOrder(order ent.Order) (string, error)
}

type SubmitOrderUseCase = func(order ent.Order) (string, error)

func MakeSubmitOrderUseCase(orderPlacer OrderPlacer) SubmitOrderUseCase {
	return func(order ent.Order) (string, error) {
		return orderPlacer.PlaceOrder(order)
	}
}
