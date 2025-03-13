package transaction

import (
	"fmt"

	"github.com/neotmp/go-trading/database"
)

func dbList(id uint16) ([]*Transaction, error) {

	q := `SELECT id, timestamp, broker_id, account_id, currency_id, direction, amount, memo, status, fee FROM transactions
	WHERE broker_id = $1`

	rows, err := database.DB.Query(q, id)
	if err != nil {

		fmt.Println("Here1")
		return nil, err
	}
	defer rows.Close()

	var data []*Transaction

	for rows.Next() {

		var t Transaction

		err := rows.Scan(&t.Id,
			&t.Timestamp,
			&t.BrokerId,
			&t.AccountId,
			&t.CurrencyId,
			&t.Direction,
			&t.Amount,
			&t.Memo,
			&t.Status,
			&t.Fee,
		)
		if err != nil {
			fmt.Println("Here2")
			return nil, err
		}

		data = append(data, &t)
	}

	// get any error encountered during iteration
	err = rows.Err()
	if err != nil {
		fmt.Println("Here3")
		return nil, err
	}

	return data, nil
}
