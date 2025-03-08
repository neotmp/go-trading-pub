package order

import (
	"fmt"

	"github.com/neotmp/go-trading/database"
)

func dbList(id uint16) ([]*Order, error) {

	q := `SELECT id, "type", volume, pair, "timestamp", price, sl, ts, tp, memo, status, pair_id,
	account_id, broker_id, direction, commission, margin  FROM orders
	WHERE broker_id = $1`

	rows, err := database.DB.Query(q, id)
	if err != nil {

		fmt.Println("Here1")
		return nil, err
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
			&o.Direction,
			&o.Commission,
			&o.Margin,
		)
		if err != nil {
			fmt.Println("Here2")
			return nil, err
		}

		data = append(data, &o)
	}

	// get any error encountered during iteration
	err = rows.Err()
	if err != nil {
		fmt.Println("Here3")
		return nil, err
	}

	return data, nil
}
