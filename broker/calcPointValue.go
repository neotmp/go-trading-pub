package broker

import (
	"fmt"

	"github.com/neotmp/go-trading/pairs"
)

// CalculatePointValue returns float32 or error
// Parameters: accepts Pair as string
func (b *Broker) CalculatePointValue(accId uint16, p string) (float32, error) {

	aId, err := b.FindAccountIndex(accId)
	if err != nil {
		return 0.0, err
	}

	accountCurrency := b.Accounts[aId].Currency
	baseCurrency := p[:3]
	quoteCurrency := p[3:]

	fmt.Println(accountCurrency, baseCurrency, quoteCurrency)

	legalPairs, err := pairs.ListCurrencies()
	if err != nil {
		return 0, err
	}

	var f float32

	if accountCurrency == quoteCurrency {
		f = 1
		fmt.Println(f, "Acc currency")
	} else if findCurrencyPair(legalPairs, accountCurrency+quoteCurrency) {
		f = (1 / pairs.GetLatestPrice(accountCurrency+"_"+quoteCurrency).Close)
		fmt.Println(f, "Normal")

	} else {
		f = pairs.GetLatestPrice(quoteCurrency + "_" + accountCurrency).Close
		fmt.Println(f, "Reversed")

	}

	return f, nil

}
