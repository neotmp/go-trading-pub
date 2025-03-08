package position

import (
	"fmt"

	"github.com/neotmp/go-trading/account"
)

func AccountFreeMargin(a *account.Account) (*account.Account, error) {

	a.FreeMargin = a.Equity - a.Margin

	fmt.Println(a.FreeMargin, "A Free Margin")

	return a, nil

}
