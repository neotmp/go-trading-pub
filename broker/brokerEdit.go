package broker

func (b *Broker) BrokerEdit() (*Broker, error) {

	eb, err := b.dbBrokerEdit()
	if err != nil {
		return b, err
	}

	return eb, nil
}
