package account

import "errors"

// AccountClose updates Active to false and returns pointer to Broker
// Account cannot be closed with none-zero  balance
func (a *Account) Close() (*Account, error) {

	if a.Balance != 0 {
		return a, errors.New("you cannot close an account with none-zero balance")
	}

	a.Active = false

	return a, nil
}
