package position

import (
	"fmt"

	"github.com/neotmp/go-trading/account"
)

// Change calculates change as profit/balance * 100
func AccountChange(a *account.Account) (*account.Account, error) {

	a.Change = a.Profit / a.Balance * 100

	fmt.Println(a.Change, "A Change")

	return a, nil
}
