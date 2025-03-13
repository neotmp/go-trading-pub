package transaction

import (
	"github.com/neotmp/go-trading/database"
)

func (t *Transaction) dbRecord() (*Transaction, error) {

	qr := `INSERT INTO transactions (timestamp, broker_id, account_id, currency_id, direction, amount, memo, status, fee) VALUES($1, $2, $3, $4, $5, $6, $7,
	$8, $9) RETURNING id`

	if err := database.DB.QueryRow(qr,
		t.Timestamp,
		t.BrokerId,
		t.AccountId,
		t.CurrencyId,
		t.Direction,
		t.Amount,
		t.Memo,
		t.Status,
		t.Fee,
	).Scan(&t.Id); err != nil {
		return nil, err
	}

	return t, nil
}
