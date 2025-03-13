package position

import (
	"fmt"

	"github.com/neotmp/go-trading/account"
)

func isJPY(p string) bool {

	switch p {
	case "AUDJPY":
		return true
	case "CADJPY":
		return true
	case "GBPJPY":
		return true
	case "EURJPY":
		return true
	case "USDJPY":
		return true
	case "NZDJPY":
		return true
	}

	return false

}

// Profit calculates profit accrued upon closing a position
// Swap is calculated in account currency already, no afc applies here
// Commision is calculated once in account currency, no afc applies here
// Profit is calculated in quote and base curency, thus afc applies
func Profit(p *Position) (*Position, error) {

	// curPrice, err := pairs.GetLatestPrice(p.Pair)
	// if err != nil {
	// 	return nil, err
	// }

	a, err := account.Get(p.AccountId)
	if err != nil {
		return nil, err
	}

	pv, err := PointValue(p.Pair, a.Currency)
	if err != nil {
		return p, err
	}

	// 	// 1 - Close Buy
	if p.Direction == 1 {

		// special case for JPY
		// TO DO Change later to id
		if isJPY(p.Pair) {
			p.Profit = ((p.CurrentPrice-p.Price)*(p.Volume*100_000)*0.1)*pv + p.Swap + (-p.Commission)
			fmt.Println(p.Profit, "Buy close J")
		} else {
			p.Profit = (p.CurrentPrice-p.Price)*(p.Volume*100_000)*pv + p.Swap + (-p.Commission)
			fmt.Println(p.Profit, p.Direction, p.Price, "Buy close")
		}

		// 2 - Close Sell
	} else {
		// special case for JPY
		// TO DO Change later to id
		if isJPY(p.Pair) {
			p.Profit = ((p.Price-p.CurrentPrice)*(p.Volume*100_000)*0.1)*pv + p.Swap + (-p.Commission)
			fmt.Println(p.Profit, "Sale close J")
		} else {
			p.Profit = (p.Price-p.CurrentPrice)*(p.Volume*100_000)*pv + p.Swap + (-p.Commission)
			fmt.Println(p.Profit, p.Price, p.CurrentPrice, "Sale close")

		}

	}

	return p, nil

}
