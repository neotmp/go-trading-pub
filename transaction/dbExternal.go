package transaction

import (
	"github.com/neotmp/go-trading/account"
	"github.com/neotmp/go-trading/database"
)

func dbExternal(a *account.Account) (*account.Account, error) {

	q := `UPDATE accounts SET
	balance = $1
	WHERE id = $2 
	RETURNING id`

	if err := database.DB.QueryRow(q, &a.Balance,
		&a.Id).Scan(&a.Id); err != nil {
		return nil, err
	}

	return a, nil
}
