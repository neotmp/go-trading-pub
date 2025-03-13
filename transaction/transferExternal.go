package transaction

import (
	"github.com/neotmp/go-trading/account"
)

// Transact withdraws/deposits amount to/from account.
func (t *Transaction) transferExternal() ([]*Transaction, []*account.Account, error) {

	var tr = []*Transaction{}
	var acc = []*account.Account{}

	// 1 - debit, withdrawal
	if t.Direction == 1 {

		t, a, err := t.Withdraw()
		if err != nil {
			return nil, nil, err
		}

		t, err = t.recordTransaction()
		if err != nil {
			return nil, nil, err
		}

		tr = append(tr, t)
		acc = append(acc, a)

		// 2 - credit, deposit
	} else {

		t, a, err := t.Deposit()
		if err != nil {
			return nil, nil, err
		}

		t, err = t.recordTransaction()
		if err != nil {
			return nil, nil, err
		}

		tr = append(tr, t)
		acc = append(acc, a)

	}

	return tr, acc, nil

	// update account

	// if internal transfer we should launch account update: free margin, equity, margin level
}
