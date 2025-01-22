package broker

// CalculateMarginLevel returns pointer to Broker
// Margin Level =  (Equity/Margin) * 100
func (b *Broker) CalculateMarginLevel(a *Account) (*Broker, error) {

	accId, err := b.AccountIndexFind(a.Id)
	if err != nil {
		return b, err
	}

	if b.Accounts[accId].Margin == 0 {
		b.Accounts[accId].MarginLevel = float32(100)
		return b, nil
	}

	b.Accounts[accId].MarginLevel = (b.Accounts[accId].Equity / b.Accounts[accId].Margin) * 100

	return b, nil
}
