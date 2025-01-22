package broker

// CalcAccountFreeMargin returns pointer to Broker and error
// Free Margin is a dynamic property and calculates every time there is a change in positions
// or we update prices
// F.M = Equity - margin
func (b *Broker) CalcAccountFreeMargin(a *Account) (*Broker, error) {

	accIndex, err := b.AccountIndexFind(a.Id)
	if err != nil {
		return b, err
	}

	if len(b.Positions) == 0 {
		b.Accounts[accIndex].FreeMargin = b.Accounts[accIndex].Balance
		return b, nil

	}

	var fm float32

	for _, v := range b.Positions {
		if v.AccountId == a.Id {
			fm += v.Margin
		}
	}

	b.Accounts[accIndex].FreeMargin = b.Accounts[accIndex].Equity - fm

	updated, err := b.dbAccountUpdate(a)
	if err != nil {
		return b, err
	}

	return updated, nil
}
