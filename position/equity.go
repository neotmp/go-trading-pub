package position

import (
	"fmt"

	"github.com/neotmp/go-trading/account"
)

// CalculateEquity returns pointer to Broker
// If no open positions, E = B
// otherwise, E = Sum of Open positions(Profit1... + Profitn)
func Equity(accId uint16) (*account.Account, error) {

	p, err := ListPositionsByAccount(accId)
	if err != nil {
		return nil, err
	}

	a, err := account.Get(accId)
	if err != nil {
		return nil, err
	}

	for _, v := range p {

		p, err := Profit(v)
		if err != nil {
			return a, err
		}
		v.Profit = p.Profit
		a.Equity += p.Profit

		fmt.Println(a, "E2")
	}

	fmt.Println(a.Equity, "E")

	return a, nil

}
