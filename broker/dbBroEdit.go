package broker

import (
	"time"

	"github.com/neotmp/go-trading/database"
)

func (b *Broker) dbBrokerEdit() (*Broker, error) {

	q := `UPDATE brokers SET
	name = $1,
	country = $2,
	phone = $3,
	email = $4,
	memo = $5,
	status = $6,
	edited_at = $7
	WHERE id = $8 
	RETURNING id`

	if err := database.DB.QueryRow(q, &b.Name,
		&b.Country,
		&b.Phone,
		&b.Email,
		&b.Memo,
		&b.Status,
		time.Now(),
		&b.Id).Scan(&b.Id); err != nil {
		return b, err
	}

	return b, nil
}
