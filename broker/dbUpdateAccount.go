package broker

import (
	"github.com/neotmp/go-trading/database"
)

// dbUpdateAccount updates table everytime when we:
// open new position,
// modify position,
// close position
// change account info,
func (b *Broker) dbUpdateAccount(a *Account) (*Broker, error) {

	q := `UPDATE accounts SET
	balance = $1,
	contract = $2,
	type = $3,
	currency = $4,
	leverage = $5,
	lot = $6,
	stopout = $7,
	equity = $8,
	free_margin = $9,
	margin = $10,
	opened_at = $11,
	closed_at = $12,
	active = $13,
	hedge = $14,
	margin_level = $15,
	broker_id = $16,
	memo = $17,
	contract_id = $18,
	currency_id = $19,
	profit = $20
	WHERE id = $21 
	RETURNING id`

	if err := database.DB.QueryRow(q, &a.Balance,
		&a.Contract,
		&a.Type,
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
		&a.CurrencyId,
		&a.Profit,
		&a.Id).Scan(&a.Id); err != nil {
		return b, err
	}

	return b, nil
}
