package position

import (
	"github.com/neotmp/go-trading/account"
)

func AccountEquity(a *account.Account) (*account.Account, error) {

	a.Equity = a.Balance + a.Profit

	return a, nil

}
