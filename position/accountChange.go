package position

import (
	"github.com/neotmp/go-trading/account"
)

// Change calculates change as profit/balance * 100
func (p *Position) AccountChange(a *account.Account) (*account.Account, error) {

	// before balance changes
	a.Change = p.Profit / a.Balance * 100

	return a, nil
}
