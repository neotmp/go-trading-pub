package broker

func (b *Broker) BrokerCreate() (*Broker, error) {

	nb, err := b.dbBrokerCreate()
	if err != nil {
		return b, err
	}

	return nb, nil
}
