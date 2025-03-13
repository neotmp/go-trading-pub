package broker

import (
	"errors"
	"fmt"

	"github.com/neotmp/go-trading/account"
)

// BrokerClose updates status from active to closed.
// Broker can only be closed when all accounts total ZERO
func (b *Broker) BrokerClose() (*Broker, error) {

	var total float32 = 0

	accs, err := account.List(b.Id)
	if err != nil {
		return b, err
	}

	for _, v := range accs {
		fmt.Println(v.Balance, "balance")
		total += v.Balance
		fmt.Println(total, "total")
	}

	if total != 0 {
		return b, errors.New("broker with none-zero account(s) cannot be closed.")
	} else {
		b.Active = false
	}

	eb, err := b.dbEdit()
	if err != nil {
		return b, err
	}

	return eb, nil

}
