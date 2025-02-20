package broker

import "errors"

func BrokerFind(br []*Broker, id uint16) (*Broker, error) {

	for _, v := range br {
		if v.Id == id {
			return v, nil
		}
	}

	return &Broker{}, errors.New("not found")

}
