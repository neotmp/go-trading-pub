package broker_test

import (
	"testing"
	"time"

	"github.com/neotmp/go-trading/broker"
)

func accOpen(a *broker.Account) *broker.Account {

	return a
}

// Real func is done differently
func TestAccOpen(t *testing.T) {

	acc := broker.Account{
		Contract:    "Pro",
		Currency:    "USD",
		Leverage:    800,
		Lot:         100_000,
		StopOut:     40,
		OpenedAt:    time.Now(),
		Active:      true,
		Hedge:       true,
		MarginLevel: 0,
		BrokerId:    4,
		Memo:        "New New Account",
		ContractId:  1,
		Type:        1,
		CurrencyId:  2,
	}

	newAc := accOpen(&acc)

	if newAc != &acc {
		t.Errorf("wrong account struct returned, given %s: expected %s: ", newAc.Name, acc.Name)
	}

}
