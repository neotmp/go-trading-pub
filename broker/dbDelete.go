package broker

import (
	"errors"

	"github.com/neotmp/go-trading/database"
)

func (b *Broker) dbDelete() error {

	q := `DELETE FROM brokers WHERE id = $1`

	res, err := database.DB.Exec(q, b.Id)
	if err != nil {
		return err
	}

	count, err := res.RowsAffected()
	if err != nil {
		return err
	}

	// make sure that we deleted the row
	// else return an error
	if count == int64(0) {
		return errors.New("could not delete broker")
	}

	return nil

}
