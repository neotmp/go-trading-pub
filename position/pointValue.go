package position

import (
	"fmt"

	"github.com/neotmp/go-trading/pairs"
)

// PointValue returns float32 or error
// Pars: p string, pair, accCurrency, string
func PointValue(p, accCur string) (float32, error) {

	accountCurrency := accCur
	baseCurrency := p[:3]
	quoteCurrency := p[3:]

	fmt.Println(accountCurrency, baseCurrency, quoteCurrency)

	// get the list fo available currency pairs
	legalPairs, err := pairs.List()
	if err != nil {
		return 0, err
	}

	var f float32

	if accountCurrency == quoteCurrency {
		f = 1
		fmt.Println(f, "Acc currency")
	} else if pairs.Find(legalPairs, accountCurrency+quoteCurrency) {
		f = (1 / pairs.GetLatestPrice(accountCurrency+"_"+quoteCurrency).Close)
		fmt.Println(f, "Normal")

	} else {
		f = pairs.GetLatestPrice(quoteCurrency + "_" + accountCurrency).Close
		fmt.Println(f, "Reversed")

	}

	return f, nil

}
