package account

import (
	"database/sql"

	"github.com/neotmp/go-trading/database"
)

func dbGet(id uint16) (*Account, error) {

	a := new(Account)

	q := `SELECT id,
	balance,
	contract,
	currency, 
	leverage, 
	lot, 
	stopout,
	equity,
	free_margin,
	margin,
	opened_at,
	closed_at,
	active, 
	hedge, 
	margin_level,
	broker_id, 
	memo, 
	contract_id, 
	type, 
	currency_id,
	profit, 
	name,
	change
	FROM accounts WHERE
	id = $1`

	if err := database.DB.QueryRow(q,
		id).Scan(&a.Id,
		&a.Balance,
		&a.Contract,
		&a.Currency,
		&a.Leverage,
		&a.Lot,
		&a.StopOut,
		&a.Equity,
		&a.FreeMargin,
		&a.Margin,
		&a.OpenedAt,
		&a.ClosedAt,
		&a.Active,
		&a.Hedge,
		&a.MarginLevel,
		&a.BrokerId,
		&a.Memo,
		&a.ContractId,
		&a.Type,
		&a.CurrencyId,
		&a.Profit,
		&a.Name,
		&a.Change,
	); err != nil {

		if err == sql.ErrNoRows {
			return nil, err
		}
		return nil, err
	}

	return a, nil

}
