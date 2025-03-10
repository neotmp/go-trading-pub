package position

import "github.com/neotmp/go-trading/account"

func AccountUpdate(a *account.Account) error {

	a, err := a.DbUpdate()
	if err != nil {
		return err
	}

	return nil
}
