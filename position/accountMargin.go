package position

import (
	"github.com/neotmp/go-trading/account"
)

func (p *Position) AccountMargin(a *account.Account) (*account.Account, error) {

	a.Margin -= p.Margin

	return a, nil

}
