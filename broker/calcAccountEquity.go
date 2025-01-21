package broker

import "fmt"

// CalculateEquity returns pointer to Broker
// Equity = Balance + Profit
func (b *Broker) CalculateAccountEquity(a *Account) (*Broker, error) {

	fmt.Println(a, "ACC")

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
	pos, err := b.FindPositions(accInd)
	if err != nil {
		return b, err
	}

	if len(pos) == 0 {
		b.Accounts[accInd].Equity = b.Accounts[accInd].Balance
		return b, nil
	}

	eq := b.Accounts[accInd].Balance + b.Accounts[accInd].Profit
	b.Accounts[accInd].Equity = eq

	return b, nil

}
