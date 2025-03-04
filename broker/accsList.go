package broker

import (
	"fmt"

	"github.com/neotmp/go-trading/database"
)

// AccountList returnls a slice of pointers too all accounts of teh broker with given id
func (b *Broker) AccountsList() ([]*Account, error) {

	q := `SELECT id, balance, contract, type, currency, leverage, lot,
	stopout, equity, free_margin, margin, margin_level, opened_at, closed_at, active,
	hedge, broker_id, memo, contract_id, currency_id, profit, name FROM accounts WHERE broker_id = $1 ORDER BY id`

	rows, err := database.DB.Query(q, b.Id)
	if err != nil {
		return []*Account{}, err
	}
	defer rows.Close()

	var data []*Account

	for rows.Next() {

		var a Account

		err := rows.Scan(&a.Id,
			&a.Balance,
			&a.Contract,
			&a.Type,
			&a.Currency,
			&a.Leverage,
			&a.Lot,
			&a.StopOut,
			&a.Equity,
			&a.FreeMargin,
			&a.Margin,
			&a.MarginLevel,
			&a.OpenedAt,
			&a.ClosedAt,
			&a.Active,
			&a.Hedge,
			&a.BrokerId,
			&a.Memo,
			&a.ContractId,
			&a.CurrencyId,
			&a.Profit,
			&a.Name,
		)
		if err != nil {
			fmt.Println(err)
		}

		data = append(data, &a)
	}

	// get any error encountered during iteration
	err = rows.Err()
	if err != nil {
		return []*Account{}, err
	}

	return data, nil

}
