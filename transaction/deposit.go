package transaction

import (
	"errors"

	"github.com/neotmp/go-trading/account"
)

// Deposit adds funds to cash account
// params account id uint16, amount float32
func (d *TDeposit) Deposit() (*TDeposit, *account.Account, error) {

	// get account
	a, err := account.Get(d.AccountId)
	if err != nil {
		return nil, nil, err
	}

	// check if it's a cash type account

	if a.Type != 0 {
		return nil, nil, errors.New("you can deposit funds only to cash account")
	}

	a.Balance += d.Amount

	tr := Transaction{
		Timestamp:  d.Timestamp,
		BrokerId:   d.BrokerId,
		AccountId:  d.AccountId,
		CurrencyId: d.CurrencyId,
		Direction:  1,
		Amount:     d.Amount,
		Memo:       d.Memo,
		Status:     1,
		Fee:        0,
	}

	_, err = tr.Transact()
	if err != nil {
		return nil, nil, err
	}

	a, err = dbUpdate(a)
	if err != nil {
		return nil, nil, err
	}

	return d, a, nil

}
