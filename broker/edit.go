package broker

func (b *Broker) Edit() (*Broker, error) {

	eb, err := b.dbEdit()
	if err != nil {
		return b, err
	}

	return eb, nil
}
