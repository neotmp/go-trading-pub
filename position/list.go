package position

import (
	"github.com/neotmp/go-trading/database"
)

func ListAll(id uint16) ([]*Position, error) {

	q := `SELECT id, pair, volume, timestamp, type, price, sl, ts, tp, profit, memo, 
	change, account_id, pair_id, broker_id, account_id, commission, margin, direction, swap
	FROM positions
	WHERE broker_id = $1 ORDER BY id`

	rows, err := database.DB.Query(q, id)
	if err != nil {

		return nil, err
	}
	defer rows.Close()

	var data []*Position

	for rows.Next() {

		var p Position

		err := rows.Scan(&p.Id,
			&p.Pair,
			&p.Volume,
			&p.Timestamp,
			&p.Type,
			&p.Price,
			&p.SL,
			&p.TS,
			&p.TP,
			&p.Profit,
			&p.Memo,
			&p.Change,
			&p.AccountId,
			&p.PairId,
			&p.BrokerId,
			&p.AccountId,
			&p.Commission,
			&p.Margin,
			&p.Direction,
			&p.Swap,
		)
		if err != nil {
			return nil, err
		}

		data = append(data, &p)
	}

	// get any error encountered during iteration
	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return data, nil
}
