package broker

import (
	"errors"
	"fmt"

	"github.com/neotmp/go-trading/database"
)

func (b *Broker) ListBrokers() ([]*Broker, error) {

	q := `SELECT id, name, country, phone, email, memo, opened_at, status FROM brokers ORDER BY id`

	rows, err := database.DB.Query(q)
	if err != nil {

		fmt.Println("Here1")
		return nil, errors.New(err.Error())
	}
	defer rows.Close()

	var data []*Broker

	for rows.Next() {

		var b Broker

		err := rows.Scan(&b.Id,
			&b.Name,
			&b.Country,
			&b.Phone,
			&b.Email,
			&b.Memo,
			&b.OpenedAt,
			&b.Status,
		)
		if err != nil {
			fmt.Println("Here2")
			return nil, errors.New(err.Error())
		}

		data = append(data, &b)
	}

	// get any error encountered during iteration
	err = rows.Err()
	if err != nil {
		fmt.Println("Here3")
		return nil, errors.New(err.Error())
	}

	// Loop over all accounts

	for _, v := range data {
		acc, err := v.ListAccounts()
		if err != nil {
			return nil, errors.New(err.Error())
		}

		for _, vv := range acc {
			v.Accounts = append(v.Accounts, vv)
		}

	}

	// Loop over all orders

	for _, v := range data {
		o, err := v.ListOrders()
		if err != nil {
			return nil, errors.New(err.Error())
		}

		for _, vv := range o {
			v.Orders = append(v.Orders, vv)
		}

	}

	// Loop over all positions

	for _, v := range data {
		o, err := v.ListPositions()
		if err != nil {
			return nil, errors.New(err.Error())
		}

		for _, vv := range o {
			v.Positions = append(v.Positions, vv)
		}

	}

	return data, nil
}
