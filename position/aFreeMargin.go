package position

import (
	"github.com/neotmp/go-trading/account"
)

func AccountFreeMargin(a *account.Account) (*account.Account, error) {

	a.FreeMargin = a.Equity - a.Margin

	return a, nil

}
