package position

import "github.com/neotmp/go-trading/account"

// AccountProfitAll adds profit of all open positions profit fields
func AccountProfitAll(a *account.Account, p []*Position) (*account.Account, error) {

	var prof float32 = 0

	for _, v := range p {
		prof += v.Profit
	}

	a.Profit = prof

	return a, nil
}
