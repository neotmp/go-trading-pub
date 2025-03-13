package position

import (
	"errors"

	"github.com/neotmp/go-trading/account"
)

// CalculateMarginLevel returns pointer to Broker
// Margin Level =  (Equity/Margin) * 100
func AccountMarginLevel(a *account.Account) (*account.Account, error) {

	// sanity check

	if a.Equity == 0 {
		return nil, errors.New("something is wrong with equity")
	}

	// no open positions check

	if a.Margin == 0 {
		a.MarginLevel = 100
	} else {

		a.MarginLevel = (a.Equity / a.Margin) * 100

	}

	return a, nil
}
