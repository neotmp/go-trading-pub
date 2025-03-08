package account

func UpdateBalance(a *Account, p float32) (*Account, error) {

	a.Balance += p

	return a, nil

}
