package position

import (
	"github.com/neotmp/go-trading/pairs"
)

// PointValue calculates account currency factor (ACF) returns float32 and error
// Params: p string, pair, accCurrency, string
func PointValue(p, accCur string) (float32, error) {

	accountCurrency := accCur
	//baseCurrency := p[:3]
	quoteCurrency := p[3:]

	//fmt.Println(accountCurrency, baseCurrency, quoteCurrency, "here")

	// get the list fo available currency pairs
	legalPairs, err := pairs.List()
	if err != nil {
		return 0, err
	}

	var f float32

	if accountCurrency == quoteCurrency {
		f = 1
		//fmt.Println(f, "Acc currency")
	} else if pairs.Find(legalPairs, accountCurrency+quoteCurrency) {
		if isJPY(p) {

			value, err := pairs.GetLatestPrice(accountCurrency + quoteCurrency)
			if err != nil {
				return f, err
			}

			f = (1 / (value.Close) / 0.1)
			//fmt.Println(f, "Normal J")
		} else {

			value, err := pairs.GetLatestPrice(accountCurrency + quoteCurrency)
			if err != nil {
				return f, err
			}

			f = (1 / value.Close)
			//fmt.Println(f, "Normal")
		}

	} else {
		if isJPY(p) {

			value, err := pairs.GetLatestPrice(accountCurrency + quoteCurrency)
			if err != nil {
				return f, err
			}

			f = (1 * (value.Close) / 0.1)
			//fmt.Println(f, "Reversed J")
		} else {

			value, err := pairs.GetLatestPrice(quoteCurrency + accountCurrency)
			if err != nil {
				return f, err
			}

			f = value.Close
			//fmt.Println(f, "Reversed")

		}

	}

	return f, nil

}
