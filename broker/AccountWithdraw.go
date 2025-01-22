package broker

import (
	"errors"
)

func (b *Broker) AccountWithdraw(from uint32, amount float32) (*Broker, error) {

	// only within free margin
	// only from trade to cash account or visa versa
	// TO DO Calculate available margin

	if len(b.Accounts) == 0 {
		return b, errors.New("no available accounts")
	}

	if b.Accounts[from].Type != 2 {
		return b, errors.New("withdrawal is available only for Cash account")
	}

	if b.Id == b.Accounts[from].BrokerId {
		b.Accounts[from].Balance -= amount
	}

	return b, nil
}
