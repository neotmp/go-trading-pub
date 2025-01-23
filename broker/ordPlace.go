package broker

import (
	"errors"
	"fmt"

	"github.com/neotmp/go-trading/pairs"
)

func findCurrencyPair(pairs []*pairs.CurrencyPair, pair string) bool {

	fmt.Println(pair, "PAIR to find")

	for _, v := range pairs {
		if v.Symbol == pair {
			return true
		}
	}

	return false
}

func (b *Broker) OrderPlace(o *Order) (*Broker, error) {

	accId, err := b.AccountIndexFind(o.AccountId)
	if err != nil {
		return b, err
	}

	acf, err := b.CalculatePointValue(accId, o.Pair)
	if err != nil {
		return b, err
	}

	o.Margin = b.CalculateMargin(o) * acf

	fm, err := b.CalculateFreeMargin(b.Accounts[accId])
	if err != nil {
		return b, err
	}

	if o.Margin > fm.Accounts[accId].FreeMargin {
		return b, errors.New("insufficient free margin to place new order")

	}

	orderPlaced, err := b.DbOrderPlace(o)
	if err != nil {
		return b, err
	}

	orderPlaced.Orders = append(b.Orders, o)

	return b, nil
}
