package broker

// CalAccountMargin returns account margin as a sum of margins of all positions
// Account Margin = Sum of {Position[1].Margin ... + Position[n].Margin}
func (b *Broker) CalcAccountMargin(a *Account) (*Broker, error) {

	accIndex, err := b.FindAccountIndex(a.Id)
	if err != nil {
		return b, err
	}

	if len(b.Positions) == 0 {
		b.Accounts[accIndex].Margin = 0
		return b, nil

	}

	for _, v := range b.Positions {
		if v.AccountId == a.Id {
			b.Accounts[accIndex].Margin += v.Margin
		}
	}

	updated, err := b.dbUpdateAccount(a)
	if err != nil {
		return b, err
	}

	return updated, nil

}
