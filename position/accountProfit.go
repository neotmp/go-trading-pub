package position

import (
	"fmt"

	"github.com/neotmp/go-trading/account"
)

func (p *Position) AccountProfit(a *account.Account) (*account.Account, error) {

	a.Profit += p.Profit

	fmt.Println(a.Profit, "A Profit")

	return a, nil

}
