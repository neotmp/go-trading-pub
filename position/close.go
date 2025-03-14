package position

import (
	"fmt"

	"github.com/neotmp/go-trading/account"
	"github.com/neotmp/go-trading/order"
)

// Close closes open position
func (p *Position) Close() (*Position, error) {

	// create new order
	// Calculate for position: profit, change
	// Calculate for account: balance, equity, margin level, margin, free margin, profit, position, change

	// profit - position and acc - when update proces or close position
	// balance - acc changes only when position is closed
	// Margin - acc - ditto
	//
	// equity - acc // call positions and profit on each , on updating prices
	// free margin - acc // equity - margin // on updating prices
	// margin level - acc //  on updating prices

	p, err := Profit(p)
	if err != nil {
		return nil, err
	}

	fmt.Println(p.Profit, "P profit")

	a, err := account.Get(p.AccountId)
	if err != nil {
		return nil, err
	}

	// calculate account balance
	a, err = AccountBalance(a, p.Profit)
	if err != nil {
		return nil, err
	}

	fmt.Println(a.Balance, "A balance")

	// TO DO ACCOUNT UPDATE, err handling, what should we do here?
	AccountUpdate(a)

	p, err = p.dbClose()
	if err != nil {
		return p, err
	}

	// finally we create order
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
		Status:     1,
		PositionId: p.Id,
	}

	// TO DO err handling
	// done
	_, err = o.Create()
	if err != nil {
		return p, err
	}

	return p, nil
}
