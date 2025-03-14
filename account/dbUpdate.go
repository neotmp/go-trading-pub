package account

import (
	"github.com/neotmp/go-trading/database"
)

// dbUpdateAccount updates table everytime when we:
// open new position,
// modify position,
// close position
// change account info,
func (a *Account) DbUpdate() (*Account, error) {

	q := `UPDATE accounts SET
	balance = $1,
	contract = $2,
	type = $3,
	leverage = $4,
	lot = $5,
	stopout = $6,
	equity = $7,
	free_margin = $8,
	margin = $9,
	opened_at = $10,
	closed_at = $11,
	active = $12,
	hedge = $13,
	margin_level = $14,
	broker_id = $15,
	memo = $16,
	contract_id = $17,
	currency_id = $18,
	profit = $19,
	change = $20
	WHERE id = $21 
	RETURNING id`

	if err := database.DB.QueryRow(q, &a.Balance,
		&a.Contract,
		&a.Type,
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
		&a.CurrencyId,
		&a.Profit,
		&a.Change,
		&a.Id).Scan(&a.Id); err != nil {
		return nil, err
	}

	return a, nil
}
