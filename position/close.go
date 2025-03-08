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

	// done
	no, err := o.Create()
	if err != nil {
		return p, err
	}

	p.OrderId = no.Id

	// profit - position and acc - when update proces or close position
	// balance - acc changes only when position is closed
	// Margin - acc - ditto
	//
	// equity - acc // call positions and profit on each , on updating prices
	// free margin - acc // equity - margin // on updating prices
	// margin level - acc //  on updating prices

	p, err = Profit(p)
	if err != nil {
		return nil, err
	}

	a, err := account.Get(p.AccountId)
	if err != nil {
		return nil, err
	}

	// When closing position don't calculate profit and credit balance at the same time
	// a, err = p.AccountProfit(a)
	// if err != nil {
	// 	return nil, err
	// }

	a, err = p.AccountBalance(a)
	if err != nil {
		return nil, err
	}

	a, err = p.AccountMargin(a)
	if err != nil {
		return nil, err
	}

	a, err = AccountEquity(a)
	if err != nil {
		return nil, err
	}

	a, err = AccountFreeMargin(a)
	if err != nil {
		return nil, err
	}

	a, err = AccountMarginLevel(a)
	if err != nil {
		return nil, err
	}

	a, err = AccountChange(a)
	if err != nil {
		return nil, err
	}

	fmt.Println(a.Profit, "A Profit")

	//Equity(p.AccountId)

	//fmt.Println(p.Profit, "P")

	// p, err = p.dbClose()
	// if err != nil {
	// 	return p, err
	// }

	return p, nil
}
