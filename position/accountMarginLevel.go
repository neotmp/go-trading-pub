package position

import (
	"fmt"

	"github.com/neotmp/go-trading/account"
)

// CalculateMarginLevel returns pointer to Broker
// Margin Level =  (Equity/Margin) * 100
func AccountMarginLevel(a *account.Account) (*account.Account, error) {

	a.MarginLevel = (a.Equity / a.Margin) * 100

	fmt.Println(a.MarginLevel, "A Margin Level")

	return a, nil
}
