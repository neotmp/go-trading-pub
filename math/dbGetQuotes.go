package math

import (
	"errors"
	"fmt"

	"github.com/neotmp/go-trading/database"
	"github.com/neotmp/go-trading/pairs"
)

func DbGetQuotes(p string) ([]*pairs.CurrencyQuote, error) {

	q := fmt.Sprintf(`SELECT id,
		date,
		open,
		high,
		low,
		close
		FROM %s`, p)

	rows, err := database.DB.Query(q)
	if err != nil {
		return nil, errors.New(err.Error())
	}
	defer rows.Close()

	var data []*pairs.CurrencyQuote

	for rows.Next() {

		var q pairs.CurrencyQuote

		err := rows.Scan(&q.Id,
			&q.Date,
			&q.Open,
			&q.High,
			&q.Low,
			&q.Close,
		)
		if err != nil {
			//fmt.Println(err)
			return nil, err
		}

		data = append(data, &q)
	}

	// get any error encountered during iteration
	err = rows.Err()
	if err != nil {
		return nil, err //errors.New(err.Error())
	}

	return data, nil
}
