package position

import (
	"fmt"

	"github.com/neotmp/go-trading/account"
)

func AccountEquity(a *account.Account) (*account.Account, error) {

	a.Equity = a.Balance + a.Profit

	fmt.Println(a.Equity, "A Equity")

	return a, nil

}
