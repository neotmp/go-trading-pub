package broker

import (
	"github.com/neotmp/go-trading/pairs"
)

// CalculateRowPNL calculates profit based on buy sell scenario
func CalculateRowPNL(p *Position, cq pairs.CurrencyQuote) *Position {

	// 1 - Buy
	if p.Type == 1 {
		p.Profit = (cq.Close - p.Price) * (p.Volume * 100_000)
		// 1 - Sell
	} else {
		p.Profit = (p.Price - cq.Close) * (p.Volume * 100_000)

	}

	return p

}

// CalcPositionProfit calculates profit fro every open position based on the latest currency quote available.
// Profit is calculated and ACF is applied, so that profit is always expressed in account currency value.
func (b *Broker) PositionCalculateProfit(a *Account) (*Broker, error) {

	pos, err := b.PositionsFind(a.Id)
	if err != nil {
		return b, err
	}

	for _, v := range pos {
		cp := pairs.GetLatestPrice(v.Pair[:3] + "_" + v.Pair[3:])
		v.Profit = CalculateRowPNL(v, *cp).Profit
		acf, err := b.CalculatePointValue(a.Id, v.Pair)
		if err != nil {
			return b, err
		}
		v.Profit *= acf
		updated, err := b.PositionUpdate(v)
		if err != nil {
			return updated, err
		}

	}

	return b, nil
}
