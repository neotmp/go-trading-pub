package transaction

func (t *Transaction) recordTransaction() (*Transaction, error) {

	tr, err := t.dbRecord()
	if err != nil {
		return nil, err
	}

	return tr, nil

}
