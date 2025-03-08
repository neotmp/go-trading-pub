package account

// AccountGet returns account based on account id or error
func Get(id uint16) (*Account, error) {

	a, err := dbGet(id)
	if err != nil {
		return nil, err
	}
	return a, nil

}
