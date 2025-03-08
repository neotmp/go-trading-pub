package pairs

import "github.com/neotmp/go-trading/database"

// ListCurrencies returns all currency pairs used by brokers
// Currently only majors
func List() ([]*CurrencyPair, error) {

	rows, err := database.DB.Query("SELECT id, symbol, name FROM currency_pairs")
	if err != nil {

		return nil, err
	}
	defer rows.Close()

	var data []*CurrencyPair

	for rows.Next() {

		var cp CurrencyPair

		err := rows.Scan(&cp.Id,
			&cp.Symbol,
			&cp.Name)
		if err != nil {
			return nil, err
		}

		data = append(data, &cp)
	}

	// get any error encountered during iteration
	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return data, nil
}
