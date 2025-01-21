package broker

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/neotmp/go-trading/database"
	"github.com/neotmp/go-trading/pairs"
)

func getAccountSpecs(accId uint16, pair string) *AccountSpecs {

	q := `SELECT id,
	account_contract_id,
	symbol,
	pip_size,
	spread_pips,
	spread_per_lot,
	swap_short_pips,
	swap_long_pips,
	pair_id, 
	lot
	FROM account_specs WHERE account_contract_id = (SELECT contract_id FROM accounts WHERE id = $1) 
	AND symbol = $2  `

	as := new(AccountSpecs)

	if err := database.DB.QueryRow(q, accId, pair).Scan(&as.Id,
		&as.AccountContractId,
		&as.Symbol,
		&as.PipSize,
		&as.SpreadPips,
		&as.SpreadPerLot,
		&as.SwapShortPips,
		&as.SwapLongPips,
		&as.PairId,
		&as.Lot,
	); err != nil {

		if err == sql.ErrNoRows {

			fmt.Println("No Account specs with this code")

		}

		fmt.Println(err)

	}

	return as

}

func findCurrencyPair(pairs []*pairs.CurrencyPair, pair string) bool {

	fmt.Println(pair, "PAIR to find")

	for _, v := range pairs {
		if v.Symbol == pair {
			return true
		}
	}

	return false
}

func (b *Broker) PlaceOrder(o *Order) (*Broker, error) {

	accId, err := b.FindAccountIndex(o.AccountId)
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

	orderPlaced, err := b.DbPlaceOrder(o)
	if err != nil {
		return b, err
	}

	orderPlaced.Orders = append(b.Orders, o)

	return b, nil
}
