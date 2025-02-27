package broker

func (b *Broker) BrokerDelete() error {

	err := b.dbBrokerDelete()
	if err != nil {
		return err
	}

	return nil
}
