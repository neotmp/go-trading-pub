package transaction

import (
	"errors"

	"github.com/neotmp/go-trading/account"
)

// Deposit adds funds to cash account
// params account id uint16, amount float32
func (t *Transaction) Deposit() (*Transaction, *account.Account, error) {

	// get account
	a, err := account.Get(t.AccountId)
	if err != nil {
		return nil, nil, err
	}

	// check if it's a cash type account

	if t.External && a.Type != 0 {
		return nil, nil, errors.New("you can deposit funds only to cash account")
	}

	a.Balance += t.Amount

	a, err = dbExternal(a)
	if err != nil {
		return nil, nil, err
	}

	return t, a, nil

}
