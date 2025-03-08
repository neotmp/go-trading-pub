package broker

func (b *Broker) Create() (*Broker, error) {

	nb, err := b.dbCreate()
	if err != nil {
		return b, err
	}

	return nb, nil
}
