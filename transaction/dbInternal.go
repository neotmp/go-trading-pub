package transaction

import (
	"context"
	"errors"
	"fmt"

	"github.com/neotmp/go-trading/account"
	"github.com/neotmp/go-trading/database"
)

// transactional operation to update accounts
func dbInternal(d, c *account.Account) ([]*account.Account, error) {

	// First You begin a transaction with a call to database.DB.Begin()
	ctx := context.Background()
	tx, err := database.DB.BeginTx(ctx, nil)
	if err != nil {
		return nil, errors.New(err.Error())
	}

	// query to create new Budget

	//fmt.Println(d.Balance, c.Balance)

	q := `UPDATE accounts SET balance = $1 WHERE id = $2 RETURNING id, balance`

	// First part of transaction
	// Update balance of account based on id and amount, return amount
	if err = tx.QueryRowContext(ctx, q, d.Balance, d.Id).Scan(&d.Id, &d.Balance); err != nil {
		tx.Rollback()
		fmt.Println(err, "roll-back 1")
		return nil, err
	}

	q2 := "UPDATE accounts SET balance = $1 WHERE id = $2 RETURNING id, balance"

	// Second part of transaction
	// Update second account's balance
	if err := tx.QueryRowContext(ctx, q2, c.Balance, c.Id).Scan(&d.Id, &d.Balance); err != nil {
		tx.Rollback()
		fmt.Println("roll-back 2")
		return nil, errors.New(err.Error())
	}

	// close the transaction with a Commit() or Rollback() method on the resulting Tx variable.
	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	var data = []*account.Account{}

	data = append(data, c, d)

	return data, nil

}
