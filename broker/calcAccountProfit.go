package broker

func (b *Broker) CalculateAccountProfit(a *Account) (*Broker, error) {

	accIndex, err := b.FindAccountIndex(a.Id)
	if err != nil {
		return b, err
	}

	if len(b.Positions) == 0 {
		b.Accounts[accIndex].Profit = 0
		return b, nil

	}

	var prof float32

	for _, v := range b.Positions {
		if v.AccountId == a.Id {
			prof += v.Profit
		}
	}

	b.Accounts[accIndex].Profit = prof

	updated, err := b.dbUpdateAccount(a)
	if err != nil {
		return b, err
	}

	return updated, nil

}
