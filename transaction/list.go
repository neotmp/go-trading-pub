package transaction

func List(id uint16) ([]*Transaction, error) {

	tr, err := dbList(id)
	if err != nil {
		return nil, err
	}

	return tr, nil
}
