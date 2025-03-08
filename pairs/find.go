package pairs

import "fmt"

// Find returns boolean. Function is used to validate currency pair
// to make sure that we have it among currencies traded in db.
func Find(pairs []*CurrencyPair, pair string) bool {

	fmt.Println(pair, "PAIR to find")

	for _, v := range pairs {
		if v.Symbol == pair {
			return true
		}
	}

	return false
}
