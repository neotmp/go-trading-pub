package account_test

import (
	"testing"
	"time"

	"github.com/neotmp/go-trading/account"
)

func accCreate(a *account.Account) *account.Account {

	return a
}

// Real func is done differently
func TestAccountCreate(t *testing.T) {

	acc := account.Account{
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

	newAc := accCreate(&acc)

	if newAc != &acc {
		t.Errorf("wrong account struct returned, given %s: expected %s: ", newAc.Name, acc.Name)
	}

}
