package uc

import (
	"errors"
	"testing"

	"github.com/tc1io/clean-arch-demo/ent"
)

type Po struct {
	fail bool
}

func (p Po) PlaceOrder(order ent.Order) (string, error) {
	if p.fail {
		return "", errors.New("error")
	} else {
		return "123", nil
	}
}

func TestSubmitOrderFail(t *testing.T) {
	po := Po{fail: true}
	uc := MakeSubmitOrderUseCase(po)

	_, err := uc(ent.Order{})
	if err == nil {
		t.Errorf("Expected error")
	}
}
func TestSubmitOrderOk(t *testing.T) {
	po := Po{fail: false}
	uc := MakeSubmitOrderUseCase(po)

	id, _ := uc(ent.Order{})
	if id != "123" {
		t.Errorf("Expected 123")
	}
}
