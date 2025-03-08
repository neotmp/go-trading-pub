package account

func (a *Account) Delete() error {

	// Should not be able to delete if any applies:
	// 1. Open position exists
	// 2. Non-zero balance exists
	// 3. Pending order exists - should just delete it as well

	// TO DO
	// func to check for balance
	// func to check for open positions
	// func to check for pending orders

	err := a.dbDelete()
	if err != nil {
		return err
	}

	return nil

}
