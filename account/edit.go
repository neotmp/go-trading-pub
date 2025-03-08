package account

// TO DO Account active can only be set to false if account balance = 0 and it has open positions.
func (a *Account) Edit(id uint16) (*Account, error) {

	// get account based on id

	// account, err := DbAccountGet(a.Id)
	// if err != nil {
	// 	return nil, err
	// }

	// check if we're trying to close account with balance
	// or open positions

	edit, err := a.dbEdit()
	if err != nil {
		return a, err
	}

	return edit, nil
}
