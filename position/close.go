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

	// When closing position don't calculate profit and credit balance at the same time
	a, err = p.AccountProfit(a)
	if err != nil {
		return nil, err
	}

	fmt.Println(a.Profit, a.Balance, "A profit, balance")

	// TO DO How do we calculate change?
	// account change
	a, err = p.AccountChange(a)
	if err != nil {
		return nil, err
	}

	fmt.Println(a.Change, "A change")

	// calculate account balance
	a, err = p.AccountBalance(a)
	if err != nil {
		return nil, err
	}

	fmt.Println(a.Balance, "A balance")

	// calculate account margin
	a, err = p.AccountMargin(a)
	if err != nil {
		return nil, err
	}

	fmt.Println(a.Margin, "A margin")

	// calculate account equity
	a, err = AccountEquity(a)
	if err != nil {
		return nil, err
	}

	fmt.Println(a.Equity, "A equity")

	// calculate account free margin
	a, err = AccountFreeMargin(a)
	if err != nil {
		return nil, err
	}

	fmt.Println(a.FreeMargin, "A free margin")

	// calculate account margin level
	a, err = AccountMarginLevel(a)
	if err != nil {
		return nil, err
	}

	fmt.Println(a.MarginLevel, "A margin level")

	// account update
	err = AccountUpdate(a)
	if err != nil {
		return nil, err
	}

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

	// TO DO ACCOUNT UPDATE

	return p, nil
}
