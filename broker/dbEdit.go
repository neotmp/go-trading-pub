package broker

import (
	"time"

	"github.com/neotmp/go-trading/database"
)

func (b *Broker) dbEdit() (*Broker, error) {

	q := `UPDATE brokers SET
	name = $1,
	country = $2,
	phone = $3,
	email = $4,
	memo = $5,
	active = $6,
	edited_at = $7
	WHERE id = $8 
	RETURNING id`

	if err := database.DB.QueryRow(q, &b.Name,
		&b.Country,
		&b.Phone,
		&b.Email,
		&b.Memo,
		&b.Active,
		time.Now(),
		&b.Id).Scan(&b.Id); err != nil {
		return b, err
	}

	return b, nil
}
