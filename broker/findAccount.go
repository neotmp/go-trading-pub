package broker

import (
	"errors"
	"fmt"
)

// FindAccount finds account in slice and returns pointer to account
// based on AccountId and BrokerId
func (b *Broker) FindAccount(ac []*Account, accId uint16) (*Account, error) {

	fmt.Println(accId, "find acc")

	if len(ac) == 0 {
		return &Account{}, errors.New("no accounts to search")
	}

	for _, v := range ac {
		fmt.Println(v.Id, accId, "loop")

		if v.Id == accId && b.Id == v.BrokerId {

			return v, nil
		}
	}

	return &Account{}, errors.New("account not found")

}
