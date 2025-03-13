package account

import (
	"github.com/neotmp/go-trading/database"
)

func (a *Account) dbCreate() (*Account, error) {

	qr := `INSERT INTO accounts (contract, currency, leverage, lot, stopout, 
	opened_at, active, hedge, broker_id, memo, contract_id, type, currency_id, name) VALUES((SELECT name FROM account_contracts WHERE id = $9), 
	(SELECT symbol FROM account_currencies WHERE id = $11), 
	$1, $2, $3, $4, $5, $6, $7,
	$8, $9, $10, $11, $12) RETURNING id`

	if err := database.DB.QueryRow(qr,
		a.Leverage,
		a.Lot,
		a.StopOut,
		a.OpenedAt,
		a.Active,
		a.Hedge,
		a.BrokerId,
		a.Memo,
		a.ContractId,
		a.Type,
		a.CurrencyId,
		a.Name,
	).Scan(&a.Id); err != nil {
		return a, err
	}

	return a, nil

}
