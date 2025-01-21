package broker

import (
	"errors"
	"fmt"

	"github.com/neotmp/go-trading/database"
)

func (b *Broker) ListOrders() ([]*Order, error) {

	q := `SELECT id, "type", volume, pair, "timestamp", price, sl, ts, tp, memo, status, pair_id,
	account_id, broker_id  FROM orders
	WHERE broker_id = $1`

	rows, err := database.DB.Query(q, b.Id)
	if err != nil {

		fmt.Println("Here1")
		return nil, errors.New(err.Error())
	}
	defer rows.Close()

	var data []*Order

	for rows.Next() {

		var o Order

		err := rows.Scan(&o.Id,
			&o.Type,
			&o.Volume,
			&o.Pair,
			&o.Timestamp,
			&o.Price,
			&o.SL,
			&o.TS,
			&o.TP,
			&o.Memo,
			&o.Status,
			&o.PairId,
			&o.AccountId,
			&o.BrokerId,
		)
		if err != nil {
			fmt.Println("Here2")
			return nil, errors.New(err.Error())
		}

		data = append(data, &o)
	}

	// get any error encountered during iteration
	err = rows.Err()
	if err != nil {
		fmt.Println("Here3")
		return nil, errors.New(err.Error())
	}

	return data, nil
}
