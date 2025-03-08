package position

import (
	"fmt"

	"github.com/neotmp/go-trading/account"
)

func (p *Position) AccountMargin(a *account.Account) (*account.Account, error) {

	a.Margin -= p.Margin

	fmt.Println(a.Margin, "A Margin")

	return a, nil

}
