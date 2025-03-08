package account

import (
	"github.com/neotmp/go-trading/database"
)

func (a *Account) dbEdit() (*Account, error) {

	q := `UPDATE accounts SET
	memo = $1,
	name = $2,
	active = $3
	WHERE id = $4 
	RETURNING id`

	if err := database.DB.QueryRow(q,
		&a.Memo,
		&a.Name,
		&a.Active,
		&a.Id).Scan(&a.Id); err != nil {
		return a, err
	}

	return a, nil
}
