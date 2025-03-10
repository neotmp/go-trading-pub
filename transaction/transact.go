package transaction

func (t *Transaction) Transact() (*Transaction, error) {

	// TO DO we should enter rejected orders with status 3

	t, err := t.dbTransact()
	if err != nil {
		return nil, err
	}

	return t, nil
}
