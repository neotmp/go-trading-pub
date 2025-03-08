package account

// Update update account params:
// balance, margin, profit, margin level, free margin,
// equity and is called by other functions
func Update(a *Account) (*Account, error) {

	a, err := a.dbUpdate()
	if err != nil {
		return nil, nil
	}

	return a, nil

}
