package transaction_test

import (
	"testing"

	"github.com/neotmp/go-trading/account"
)

func withdraw(a *account.Account, amount int) {

	if a.Type != 0 {
		return
	}
	a.Balance -= float32(amount)

}

func TestWithdraw(t *testing.T) {

	accs := []*account.Account{
		{Type: 1, Balance: 100, Name: "Trade"},
		{Type: 0, Balance: 100, Name: "Cash"},
	}

	expectedType := 0
	var expectedBalance int = 0 // use integer here for testing

	for _, v := range accs {
		withdraw(v, expectedBalance)
		if v.Type != 0 && uint(v.Balance) == uint(expectedBalance) {
			t.Errorf("canot deposit to non-cash account, given type %d:, balance %d: expected type %d:, balance %d: ", v.Type, int(v.Balance),
				expectedType, expectedBalance)
		}

	}

}
