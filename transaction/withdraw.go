package transaction

import (
	"errors"

	"github.com/neotmp/go-trading/account"
)

// Withdraw debits cash account
// params account id uint16, amount float32
// Funds can be withdrawn only within the parameters of equity
// Withdrawal results in equity drop, reduced (negative) free margin, and reduced margin level
// that may give rise to inadequate cover and subsequent rorced closure of positions - margin call
// TO DO Reflect above mentioned logic into the method.
func (w *TWithdrawal) Withdraw() (*TWithdrawal, *account.Account, error) {

	// get account
	a, err := account.Get(w.AccountId)
	if err != nil {
		return nil, nil, err
	}

	// check if it's a cash type account

	if a.Type != 0 {
		return nil, nil, errors.New("you can withdraw funds only from cash account")
	}

	// balance check
	if a.Balance < w.Amount {
		return nil, nil, errors.New("amount to withdraw can not exceed the balance of the debit account")
	}

	a.Balance -= w.Amount

	a, err = dbUpdate(a)
	if err != nil {
		return nil, nil, err
	}

	return w, a, nil

}
