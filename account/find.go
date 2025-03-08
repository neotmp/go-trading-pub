package account

import (
	"errors"
)

// FindAccount finds account in slice and returns pointer to account
// based on AccountId and BrokerId
func AccountFind(id uint16, ac []*Account, accId uint16) (*Account, error) {

	if len(ac) == 0 {
		return &Account{}, errors.New("no accounts to search")
	}

	for _, v := range ac {
		if v.Id == accId && id == v.BrokerId {
			return v, nil
		}
	}

	return &Account{}, errors.New("account not found")

}
