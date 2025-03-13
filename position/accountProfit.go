package position

import (
	"github.com/neotmp/go-trading/account"
)

func (p *Position) AccountProfit(a *account.Account) (*account.Account, error) {

	a.Profit += p.Profit

	return a, nil

}
