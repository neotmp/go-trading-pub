package position

import (
	"github.com/neotmp/go-trading/order"
)

// Create creates order first and position then if no error
func (p *Position) Create() (*Position, error) {

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

	// then we create position
	p, err = p.dbCreate()
	if err != nil {
		return p, err
	}

	return p, nil
}
