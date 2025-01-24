package broker

import "errors"

// AccountClose updates Active to false and returns pointer to Broker
// Account cannot be closed with none-zero  balance
func (b *Broker) AccountClose(a *Account) (*Broker, error) {

	if a.Balance != 0 {
		return b, errors.New("you cannot close an account with none-zero balance")
	}

	a.Active = false

	return b, nil
}
