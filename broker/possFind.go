package broker

// FindPositions returns slice of pointers to Position struct
// param is account id
// Positions are matched against AccountId and BrokerId
func (b *Broker) PositionsFind(accId uint16) ([]*Position, error) {

	accs := b.Accounts
	var positions []*Position

	account, err := b.AccountFind(accs, accId)
	if err != nil {
		return positions, err
	}

	for _, v := range b.Positions {
		var p Position

		if v.AccountId == account.Id && b.Id == v.BrokerId {
			p = *v
			positions = append(positions, &p)
		}
	}

	return positions, nil

}
