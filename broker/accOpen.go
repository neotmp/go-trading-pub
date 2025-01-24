package broker

// AccountOpen appends a pointer to new account to Accounts slise and returns Broker
func (b *Broker) AccountOpen(a *Account) (*Broker, error) {

	b.Accounts = append(b.Accounts, a)

	open, err := b.dbAccountOpen(a)
	if err != nil {
		return b, err
	}

	return open, nil
}
