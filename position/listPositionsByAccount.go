package position

// FindPositions returns slice of pointers to Position struct
// param is account id
// Positions are matched against AccountId and BrokerId
func ListPositionsByAccount(accId uint16) ([]*Position, error) {

	p, err := dbListPositionsByAccount(accId)
	if err != nil {
		return nil, err
	}

	return p, nil

}
