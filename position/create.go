package position

import (
	"fmt"

	"github.com/neotmp/go-trading/order"
)

// Create creates order first and position then if no error
func (p *Position) Create() (*Position, error) {

	// we create position
	p, err := p.dbCreate()
	if err != nil {
		return p, err
	}

	// first we create order
	o := order.Order{
		Direction:  p.Direction,
		Volume:     p.Volume,
		Timestamp:  p.Timestamp,
		Price:      p.Price,
		Memo:       p.Memo,
		PairId:     p.PairId,
		BrokerId:   p.BrokerId,
		AccountId:  p.AccountId,
		Type:       p.Type,
		Commission: p.Commission,
		PositionId: p.Id,
	}

	// create order
	// TO DO err handling
	_, err = o.Create()
	if err != nil {
		fmt.Println(err)
		return p, err
	}

	return p, nil
}
