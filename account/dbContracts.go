package account

import (
	"fmt"

	"github.com/neotmp/go-trading/database"
)

func dbContracts(id uint16) ([]*AccountContract, error) {

	q := `SELECT id, name, broker_id, memo FROM account_contracts
	WHERE broker_id = $1`

	rows, err := database.DB.Query(q, id)
	if err != nil {

		fmt.Println("Here1")
		return nil, err
	}
	defer rows.Close()

	var data []*AccountContract

	for rows.Next() {

		var ac AccountContract

		err := rows.Scan(&ac.Id,
			&ac.Name,
			&ac.BrokerId,
			&ac.Memo,
		)
		if err != nil {
			fmt.Println("Here2")
			return nil, err
		}

		data = append(data, &ac)
	}

	// get any error encountered during iteration
	err = rows.Err()
	if err != nil {
		fmt.Println("Here3")
		return nil, err
	}

	return data, nil

}
