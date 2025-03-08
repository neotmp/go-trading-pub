package position

import (
	"fmt"

	"github.com/neotmp/go-trading/account"
	"github.com/neotmp/go-trading/pairs"
)

// Profit calculates profit accrued upon closing a position
func Profit(p *Position) (*Position, error) {

	curPrice := pairs.GetLatestPrice(p.Pair)

	a, err := account.Get(p.AccountId)
	if err != nil {
		return nil, err
	}

	pv, err := PointValue(p.Pair, a.Currency)
	if err != nil {
		return p, err
	}

	// 	// 1 - Buy
	if p.Direction == 1 {
		p.Profit = (((float32(curPrice.Close)-p.Price)*(p.Volume*100_000) + p.Swap + p.Commission) * pv)
		fmt.Println(p.Profit, p.Direction, p.Price, "Buy close")
		// 2 - Sell
	} else {
		p.Profit = (((p.Price-float32(curPrice.Close))*(p.Volume*100_000) + p.Swap + p.Commission) * pv)
		fmt.Println(p.Profit, "Sale close")

	}

	return p, nil

}
