package broker

import (
	"github.com/neotmp/go-trading/database"
)

func (b *Broker) dbBrokerOpen() (*Broker, error) {

	database.Connect()

	qr := `INSERT INTO brokers (name, country, phone, email, memo, opened_at, status) VALUES($1, $2, $3, $4, $5, $6, $7) RETURNING id`

	if err := database.DB.QueryRow(qr, b.Name, b.Country, b.Phone, b.Email, b.Memo, b.OpenedAt, b.Status).Scan(&b.Id); err != nil {
		return b, err
	}

	return b, nil

}
