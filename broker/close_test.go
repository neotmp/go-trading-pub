package broker_test

import (
	"testing"

	"github.com/neotmp/go-trading/account"
	"github.com/neotmp/go-trading/broker"
)

func TestClose(t *testing.T) {

	noneZeroAcc := account.Account{
		Balance: 100.00,
	}

	b := broker.Broker{}

	b.Accounts = append(b.Accounts, &noneZeroAcc)

	var total float32 = 0

	for _, v := range b.Accounts {
		total += v.Balance
	}

	if total != 0 {
		t.Errorf("broker with none-zero account(s) cannot be closed. Total given %f:", total)
	}

}
