package account

// AccountDeposit changes the balance of the account
// and create a corresponding transaction
// It returns pointer to Account
func (a *Account) Deposit() (*Account, error) {

	// changes the balance of the cash account
	a, err := dbDeposit()
	if err != nil {
		return nil, err
	}

	return a, nil
}
