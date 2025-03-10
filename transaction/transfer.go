package transaction

import (
	"errors"
	"fmt"

	"github.com/neotmp/go-trading/account"
)

// Transfer debits one account and credits the other by the same amount
// params debit account id uint16, credit account id uint16, amount float32
// It returns transaction, account, and error
// Account cannot be debited for amount that is larger its balance.
func (t *TTransfer) Transfer() (*TTransfer, *account.Account, error) {

	// get account to debit
	d, err := account.Get(t.DebitId)
	if err != nil {
		return nil, nil, err
	}

	fmt.Println(d.Balance, "D balance 1")

	// balance check
	if d.Balance < t.Amount {
		return nil, nil, errors.New("amount to transfer can not exceed the balance of the debit account")
	}

	// get account to credit
	c, err := account.Get(t.CreditId)
	if err != nil {
		return nil, nil, err
	}
	fmt.Println(c.Balance, "C balance 1")

	// subtract amount from debit account and add it to credit account
	d.Balance -= t.Amount
	c.Balance += t.Amount

	fmt.Println(d.Balance, "D balance 2")
	fmt.Println(c.Balance, "C balance 2")

	tr := Transaction{
		Timestamp:  t.Timestamp,
		BrokerId:   t.BrokerId,
		AccountId:  t.DebitId,
		CurrencyId: t.CurrencyId,
		Direction:  1,
		Amount:     t.Amount,
		Memo:       t.Memo,
		Status:     1,
		Fee:        0,
	}

	record, err := tr.Transact()
	if err != nil {
		return nil, nil, err
	}

	record.AccountId = t.CreditId
	record.Direction = 2
	record, err = record.Transact()
	if err != nil {
		return nil, nil, err
	}

	// TO DO in one transaction
	// update debit account
	a, err := dbUpdate(d)
	if err != nil {
		return nil, nil, err
	}

	// update credit account
	a, err = dbUpdate(c)
	if err != nil {
		return nil, nil, err
	}

	return t, a, nil

}
