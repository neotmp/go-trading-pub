package broker

import (
	"fmt"

	"github.com/neotmp/go-trading/account"
	"github.com/neotmp/go-trading/database"
	"github.com/neotmp/go-trading/order"
	"github.com/neotmp/go-trading/position"
	"github.com/neotmp/go-trading/transaction"
)

// BrokerLists returns the list of all brokers w/ nested Accounts, Orders and Positions as slices
func List() ([]*Broker, error) {

	q := `SELECT id, name, country, phone, email, memo, opened_at, active, edited_at FROM brokers ORDER BY id`

	rows, err := database.DB.Query(q)
	if err != nil {

		fmt.Println("Here1")
		return nil, err
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
			&b.Active,
			&b.EditedAt,
		)
		if err != nil {
			fmt.Println("Here2")
			return nil, err
		}

		data = append(data, &b)
	}

	// get any error encountered during iteration
	err = rows.Err()
	if err != nil {
		fmt.Println("Here3")
		return nil, err
	}

	// Loop over all accounts

	for _, v := range data {
		acc, err := account.List(v.Id)
		if err != nil {
			return nil, err
		}
		v.Accounts = append(v.Accounts, acc...)

	}

	// Loop over all orders

	for _, v := range data {
		o, err := order.List(v.Id)
		if err != nil {
			return nil, err
		}

		v.Orders = append(v.Orders, o...)

	}

	// Loop over all positions

	for _, v := range data {

		pos, err := position.ListAll(v.Id)
		if err != nil {
			return nil, err
		}

		v.Positions = append(v.Positions, pos...)

	}

	// Loop over all transactions

	for _, v := range data {

		tr, err := transaction.List(v.Id)
		if err != nil {
			return nil, err
		}

		v.Transactions = append(v.Transactions, tr...)

	}

	return data, nil
}
