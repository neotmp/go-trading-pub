package broker

import (
	"errors"
)

// FindAccount finds account in slice and returns pointer to account
// based on AccountId and BrokerId
func (b *Broker) AccountFind(ac []*Account, accId uint16) (*Account, error) {

	if len(ac) == 0 {
		return &Account{}, errors.New("no accounts to search")
	}

	for _, v := range ac {
		if v.Id == accId && b.Id == v.BrokerId {
			return v, nil
		}
	}

	return &Account{}, errors.New("account not found")

}
