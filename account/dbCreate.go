package account

import (
	"github.com/neotmp/go-trading/database"
)

func (a *Account) dbCreate(id uint16) (*Account, error) {

	qr := `INSERT INTO accounts (contract, currency, leverage, lot, stopout, 
	opened_at, active, hedge, broker_id, memo, contract_id, type, currency_id, name) VALUES($1, $2, $3, $4, $5, $6, $7,
	$8, $9, $10, $11, $12, $13, $14) RETURNING id`

	if err := database.DB.QueryRow(qr,
		a.Contract,
		a.Currency,
		a.Leverage,
		a.Lot,
		a.StopOut,
		a.OpenedAt,
		a.Active,
		a.Hedge,
		id,
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
