package account

func contracts(id uint16) ([]*AccountContract, error) {

	a, err := dbContracts(id)
	if err != nil {
		return nil, err
	}

	return a, nil

}
