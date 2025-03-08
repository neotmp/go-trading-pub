package account

func (a *Account) Withdraw(from uint32, amount float32) (*Account, error) {

	// if len(b.Accounts) == 0 {
	// 	return b, errors.New("no available accounts")
	// }

	// if b.Accounts[from].Type != 2 {
	// 	return b, errors.New("withdrawal is available only for Cash account")
	// }

	// if b.Id == b.Accounts[from].BrokerId {
	// 	b.Accounts[from].Balance -= amount
	// }

	return a, nil
}
