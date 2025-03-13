package order

import (
	"github.com/neotmp/go-trading/database"
)

// DbPlaceOrder returns pointer to Broker or error
// This is a two step transaction using Context
// Insert new Order
// Insert new Position
// If either failed, we rollback and return an error
func (o *Order) dbCreate() (*Order, error) {

	q1 := `INSERT INTO orders (type, volume, pair, pair_id, timestamp, price, sl, ts, tp, memo, status, broker_id, account_id, margin, direction, commission, position_id)
	VALUES($1, $2, (SELECT symbol FROM currency_pairs WHERE id = $3), $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16) RETURNING id`

	if err := database.DB.QueryRow(q1, o.Type, o.Volume, o.PairId, o.Timestamp,
		o.Price, o.SL, o.TS, o.TP, o.Memo, 1, o.BrokerId, o.AccountId, o.Margin, o.Direction, o.Commission, o.PositionId).Scan(&o.Id); err != nil {

		return o, err
	}

	return o, nil

}
