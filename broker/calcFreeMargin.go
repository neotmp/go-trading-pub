package broker

// CalculateFreeMargin returns pointer to Broker
// Free Margin = Equity - Margin
// It is used to check if opening new position is allowed based on available free margin
func (b *Broker) CalculateFreeMargin(a *Account) (*Broker, error) {
	accId, err := b.AccountIndexFind(a.Id)
	if err != nil {
		return b, err
	}

	b.Accounts[accId].FreeMargin = (b.Accounts[accId].Equity - b.Accounts[accId].Margin)

	return b, nil

}
