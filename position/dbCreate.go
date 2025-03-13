package position

import (
	"errors"

	"github.com/neotmp/go-trading/database"
)

func (p *Position) dbCreate() (*Position, error) {

	q := `INSERT INTO positions (pair, volume, timestamp, type, price, sl, ts, tp, profit, memo, pair_id, change, broker_id,
	account_id, commission, margin, direction) 
	VALUES((SELECT symbol FROM currency_pairs WHERE id =  $10), $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16) RETURNING id`

	if err := database.DB.QueryRow(q, p.Volume, p.Timestamp, p.Type, p.Price, p.SL, p.TS, p.TP, p.Profit, p.Memo, p.PairId,
		p.Change, p.BrokerId, p.AccountId, p.Commission, p.Margin, p.Direction).Scan(&p.Id); err != nil {

		return p, errors.New(err.Error())
	}

	return p, nil
}
