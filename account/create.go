package account

// Create creates a new account and returns a pointer to Account struct
// Params id uint16 - broker id
func (a *Account) Create() (*Account, error) {

	open, err := a.dbCreate()
	if err != nil {
		return a, err
	}

	return open, nil
}
