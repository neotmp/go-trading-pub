package broker

func (b *Broker) AccountTransfer(from, to uint32, amount float32) (*Broker, error) {

	// only within free margin
	// only from trade to cash account or visa versa

	b.Accounts[from].Balance -= amount
	b.Accounts[to].Balance += amount
	return b, nil
}
