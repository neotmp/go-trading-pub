package broker

import (
	"context"
	"errors"
	"fmt"

	"github.com/neotmp/go-trading/database"
)

// DbPlaceOrder returns pointer to Broker or error
// This is a two step transaction using Context
// Insert new Order
// Insert new Position
// If either failed, we rollback and return an error
func (b *Broker) DbPlaceOrder(o *Order) (*Broker, error) {

	// First You begin a transaction with a call to database.DB.Begin()
	ctx := context.Background()
	tx, err := database.DB.BeginTx(ctx, nil)
	if err != nil {
		return b, errors.New(err.Error())
	}

	q1 := `INSERT INTO orders (type, volume, pair, pair_id, timestamp, price, sl, ts, tp, memo, status, broker_id, account_id, margin) 
	VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14) RETURNING id`

	// First part of transaction
	// Insert New Order into table and return Id
	if err = tx.QueryRowContext(ctx, q1, o.Type, o.Volume, o.Pair, o.PairId, o.Timestamp,
		o.Price, o.SL, o.TS, o.TP, o.Memo, o.Status, o.BrokerId, o.AccountId, o.Margin).Scan(&o.Id); err != nil {
		tx.Rollback()
		fmt.Println("roll-back 1")
		return b, errors.New(err.Error())
	}

	fmt.Println(o, "order")

	p := Position{
		Type:       o.Type,
		OrderId:    o.Id,
		Volume:     o.Volume,
		Pair:       o.Pair,
		PairId:     o.PairId,
		Timestamp:  o.Timestamp,
		Price:      o.Price,
		SL:         o.SL,
		TS:         o.TS,
		TP:         o.TP,
		Change:     0,
		BrokerId:   o.BrokerId,
		AccountId:  o.AccountId,
		Commission: o.Commission,
		Margin:     o.Margin,
	}

	q2 := `INSERT INTO positions (pair, volume, timestamp, type, price, sl, ts, tp, profit, memo, order_id, pair_id, change, broker_id,
	account_id, commission, margin) 
	VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17) RETURNING id`

	// Second part of transaction
	// Insert New Position
	_, err = tx.ExecContext(ctx, q2, p.Pair, p.Volume, p.Timestamp, p.Type, p.Price, p.SL, p.TS, p.TP, p.Profit, p.Memo, p.OrderId, p.PairId,
		p.Change, p.BrokerId, p.AccountId, p.Commission, p.Margin)
	if err != nil {
		tx.Rollback()
		fmt.Println("roll-back 2")
		return b, errors.New(err.Error())
	}

	// close the transaction with a Commit() or Rollback() method on the resulting Tx variable.
	err = tx.Commit()
	if err != nil {
		return b, errors.New(err.Error())
	}

	return b, nil
}
