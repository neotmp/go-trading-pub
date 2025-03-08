package position

import (
	"github.com/neotmp/go-trading/account"
)

func (p *Position) AccountBalance(a *account.Account) (*account.Account, error) {

	a.Balance += p.Profit

	return a, nil

}
