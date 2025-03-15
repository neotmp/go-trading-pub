package transaction

import (
	"errors"

	"github.com/neotmp/go-trading/position"
)

// Transfer transfers funds externally (deposit, withdrawal) or internally between two accounts.
func (t *Transaction) Transfer() (*Data, error) {

	d := new(Data)

	if !t.External && t.DebitId != 0 && t.CreditId != 0 {

		a, tr, err := t.transferInternal()
		if err != nil {
			return nil, err
		}

		d.Accounts = append(d.Accounts, a...)
		d.Transactions = append(d.Transactions, tr...)

		for _, v := range a {

			//fmt.Println(v.Balance, v.Type, "balance")

			if v.Type == 2 {

				_, err := position.AccountUpdate(v)
				if err != nil {
					return nil, err
				}

			}

		}

		return d, nil

	} else if t.External && t.AccountId != 0 {
		tr, a, err := t.transferExternal()
		if err != nil {
			return nil, err
		}

		d.Accounts = append(d.Accounts, a...)
		d.Transactions = append(d.Transactions, tr...)

		return d, nil
	} else {
		return nil, errors.New("wrong data")
	}

}
