package broker

func (b *Broker) AccountEdit(a *Account) (*Broker, error) {

	edit, err := b.dbAccountEdit(a)
	if err != nil {
		return b, err
	}

	return edit, nil
}
