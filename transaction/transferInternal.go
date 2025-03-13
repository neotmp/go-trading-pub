package transaction

import (
	"errors"

	"github.com/neotmp/go-trading/account"
)

func (t *Transaction) transferInternal() ([]*account.Account, []*Transaction, error) {

	// get account to debit
	d, err := account.Get(t.DebitId)
	if err != nil {
		return nil, nil, err
	}

	// balance check
	if d.Balance < t.Amount {
		return nil, nil, errors.New("amount to withdraw can not exceed the balance of the debit account")
	}

	var tr = []*Transaction{}

	d.Balance -= t.Amount
	t.Direction = 1

	// copy value of transaction
	trOut := *t

	// get account to debit and credit
	c, err := account.Get(t.CreditId)
	if err != nil {
		return nil, nil, err
	}

	c.Balance += t.Amount
	t.Direction = 2

	// copy value of transaction
	trIn := *t

	tr = append(tr, &trIn, &trOut)

	res, err := dbInternal(d, c)
	if err != nil {
		return nil, nil, err
	}

	var data = []*Transaction{}

	for _, v := range tr {
		d, err := v.dbRecord()
		if err != nil {
			return nil, nil, err
		}

		data = append(data, d)

	}

	// return a slice of edited accounts
	return res, data, nil

}
