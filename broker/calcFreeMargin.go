package broker

// CalculateFreeMargin returns pointer to Broker
// Free Margin = Equity - Margin
func (b *Broker) CalculateFreeMargin(a *Account) (*Broker, error) {
	accId, err := b.FindAccountIndex(a.Id)
	if err != nil {
		return b, err
	}

	b.Accounts[accId].FreeMargin = (b.Accounts[accId].Equity - b.Accounts[accId].Margin)

	return b, nil

}
