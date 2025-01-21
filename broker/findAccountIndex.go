package broker

import "errors"

func (b *Broker) FindAccountIndex(accId uint16) (uint16, error) {

	for i, v := range b.Accounts {
		if v.Id == accId {
			return uint16(i), nil
		}
	}

	return 65535, errors.New("did not find account index with given id")

}
