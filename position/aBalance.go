package position

import (
	"github.com/neotmp/go-trading/account"
)

func AccountBalance(a *account.Account, amount float32) (*account.Account, error) {

	a.Balance += amount

	return a, nil

}
