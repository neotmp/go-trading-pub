package position

import (
	"github.com/neotmp/go-trading/account"
)

// Change calculates change as profit/balance * 100
func AccountChange(a *account.Account) (*account.Account, error) {

	// before balance changes

	// sanity check

	if a.Profit == 0 {
		return a, nil
	} else {
		a.Change = a.Profit / a.Balance * 100
	}

	return a, nil
}
