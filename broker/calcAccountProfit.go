package broker

func (b *Broker) CalcAccountProfit(a *Account) (*Broker, error) {

	accIndex, err := b.FindAccountIndex(a.Id)
	if err != nil {
		return b, err
	}

	if len(b.Positions) == 0 {
		b.Accounts[accIndex].Profit = 0
		return b, nil

	}

	for _, v := range b.Positions {
		if v.AccountId == a.Id {
			b.Accounts[accIndex].Profit += v.Profit
		}
	}

	updated, err := b.dbUpdateAccount(a)
	if err != nil {
		return b, err
	}

	return updated, nil

}
