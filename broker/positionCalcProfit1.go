package broker

import (
	"github.com/neotmp/go-trading/pairs"
)

func calculatePnL(p *Position, cq pairs.CurrencyQuote) *Position {

	// 1 - Buy
	if p.Type == 1 {
		p.Profit = (cq.Close - p.Price) * (p.Volume * 100_000)
		// 1 - Sell
	} else {
		p.Profit = (p.Price - cq.Close) * (p.Volume * 100_000)

	}

	return p

}

// CalculateProfit calculates P&L of every open position
// uses Point Value to correctly calculate P&L based on account currency
// params account
func (b *Broker) PositionCalculateProfit1(a *Account) (*Broker, error) {

	pos, err := b.PositionsFind(a.Id)
	if err != nil {
		return b, err
	}

	for _, v := range pos {
		cp := pairs.GetLatestPrice(v.Pair[:3] + "_" + v.Pair[3:])
		calculatePnL(v, *cp)
		//fmt.Println(cp.Close, v.Profit, "Profit")

	}

	//fmt.Println(pos)

	return b, nil
}
