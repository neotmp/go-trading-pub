package account

// UpdateMargin Updates margin of the account
func UpdateMargin(a *Account, m float32) (*Account, error) {

	// Remove margin that was connected to position from total account margin
	a.Margin -= m

	return a, nil

}
