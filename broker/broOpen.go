package broker

func (b *Broker) BrokerOpen() (*Broker, error) {

	nb, err := b.dbBrokerOpen()
	if err != nil {
		return b, err
	}

	return nb, nil
}
