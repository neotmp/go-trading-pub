package broker

import "fmt"

// CalculateEquity returns pointer to Broker
// If no open positions, E = B
// otherwise, E = Sum of Open positions(Profit1... + Profitn)
func (b *Broker) CalculateAccountEquity(a *Account) (*Broker, error) {

	// index of account
	accInd, err := b.FindAccountIndex(a.Id)
	if err != nil {
		return b, err
	}

	fmt.Println(accInd, "Index")

	// if account.TYPE is not 1 (TRADE)
	// do nothing
	if b.Accounts[accInd].Type != 1 {
		return b, nil
	}

	// if no open positions equity = balance
	pos, err := b.FindPositions(a.Id)
	if err != nil {
		return b, err
	}

	if len(pos) == 0 {
		b.Accounts[accInd].Equity = b.Accounts[accInd].Balance
		return b, nil
	}

	var equity float32

	for _, v := range pos {
		equity += v.Profit
	}

	b.Accounts[accInd].Equity = b.Accounts[accInd].Balance + equity

	updated, err := b.dbUpdateAccount(a)
	if err != nil {
		return b, err
	}

	return updated, nil

}
