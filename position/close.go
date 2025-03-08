package position

import "github.com/neotmp/go-trading/order"

// Close closes open position
func (p *Position) Close() (*Position, error) {

	// create new order
	// Calculate for position: profit, change
	// Calculate for account: balance, equity, margin level, margin, free margin, profit, position, change

	// first we create order
	o := order.Order{
		Direction:  p.Direction,
		Volume:     p.Volume,
		Pair:       p.Pair,
		Timestamp:  p.Timestamp,
		Price:      p.Price,
		Memo:       p.Memo,
		PairId:     p.PairId,
		BrokerId:   p.BrokerId,
		AccountId:  p.AccountId,
		Type:       p.Type,
		Commission: p.Commission,
	}

	no, err := o.Create()
	if err != nil {
		return p, err
	}

	p.OrderId = no.Id

	p, err = p.dbClose()
	if err != nil {
		return p, err
	}

	return p, nil
}
