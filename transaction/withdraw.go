package transaction

import (
	"errors"

	"github.com/neotmp/go-trading/account"
)

func (t *Transaction) Withdraw() (*Transaction, *account.Account, error) {

	// TO DO check availability of funds when transfering from trade account with open positions
	// Equity cannot go negative

	// get account
	a, err := account.Get(t.AccountId)
	if err != nil {
		return nil, nil, err
	}

	// check if it's a cash type account

	if t.External && a.Type != 0 {
		return nil, nil, errors.New("you can withdraw funds only from cash account")
	}

	// balance check
	if a.Balance < t.Amount {
		return nil, nil, errors.New("amount to withdraw can not exceed the balance of the debit account")
	}

	a.Balance -= t.Amount

	a, err = dbExternal(a)
	if err != nil {
		return nil, nil, err
	}

	return t, a, nil

}
