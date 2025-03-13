package position_test

import (
	"testing"

	"github.com/neotmp/go-trading/pairs"
	"github.com/neotmp/go-trading/position"
)

func isJPY(p string) bool {

	switch p {
	case "AUDJPY":
		return true
	case "CADJPY":
		return true
	case "GBPJPY":
		return true
	case "EURJPY":
		return true
	case "USDJPY":
		return true
	case "NZDJPY":
		return true

	}

	return false

}

func find(pairs []*pairs.CurrencyPair, pair string) bool {

	//fmt.Println(pair, "PAIR to find")

	for _, v := range pairs {
		if v.Symbol == pair {
			return true
		}
	}

	return false
}

func getLatestPrice(pair string) float32 {

	type tmpCyrrencyQuote struct {
		Symbol string
		Close  float32
	}

	var close float32

	db := []*tmpCyrrencyQuote{
		{Symbol: "AUDCAD", Close: 0.90123}, //
		{Symbol: "AUDCHF", Close: 0.55123}, //
		{Symbol: "AUDJPY", Close: 92.123},  //
		{Symbol: "AUDNZD", Close: 1.10123}, //
		{Symbol: "AUDUSD", Close: 0.62123}, //
		{Symbol: "EURUSD", Close: 1.09123}, //
		{Symbol: "EURAUD", Close: 1.73123}, //
		{Symbol: "EURNZD", Close: 1.91123}, //
		{Symbol: "EURCAD", Close: 1.57123}, //
		{Symbol: "EURCHF", Close: 0.96123}, //
		{Symbol: "EURJPY", Close: 160.123}, //
		{Symbol: "EURGBP", Close: 0.84123}, //
		{Symbol: "GBPUSD", Close: 1.12123}, //
		{Symbol: "GBPAUD", Close: 2.05123}, //
		{Symbol: "GBPNZD", Close: 2.26123}, //
		{Symbol: "GBPCAD", Close: 1.86123}, //
		{Symbol: "GBPJPY", Close: 190.123}, //
		{Symbol: "USDJPY", Close: 147.123}, //
		{Symbol: "USDCAD", Close: 1.44123}, //
		{Symbol: "USDCHF", Close: 0.88123}, //
		{Symbol: "NZDUSD", Close: 0.57123}, //
		{Symbol: "NZDCAD", Close: 0.82123}, //
		{Symbol: "NZDCHF", Close: 0.50123}, //
		{Symbol: "NZDJPY", Close: 84.123},  //
		{Symbol: "CHFJPY", Close: 167.123}, //
		{Symbol: "CADCHF", Close: 0.61123}, //
		{Symbol: "CADJPY", Close: 102.123}, //
	}

	for _, v := range db {

		if v.Symbol == pair {
			close = v.Close
		}
	}

	return close

}

func calcAFC(p position.Position, accountCurrency, quoteCurrency string) float32 {

	legalPairs := []*pairs.CurrencyPair{{Symbol: "AUDCAD"}, {Symbol: "AUDCHF"}, {Symbol: "AUDJPY"}, {Symbol: "AUDNZD"}, {Symbol: "AUDUSD"}, {Symbol: "EURUSD"},
		{Symbol: "EURAUD"}, {Symbol: "EURNZD"}, {Symbol: "EURCAD"}, {Symbol: "EURCHF"}, {Symbol: "EURJPY"}, {Symbol: "EURGBP"},
		{Symbol: "GBPUSD"}, {Symbol: "GBPAUD"}, {Symbol: "GBPNZD"}, {Symbol: "GBPCAD"}, {Symbol: "GBPJPY"}, {Symbol: "USDJPY"}, {Symbol: "USDCAD"},
		{Symbol: "USDCHF"}, {Symbol: "NZDUSD"}, {Symbol: "NZDCAD"}, {Symbol: "NZDCHF"}, {Symbol: "NZDJPY"}, {Symbol: "CHFJPY"},
		{Symbol: "CADCHF"}, {Symbol: "CADJPY"}}

	var f float32

	if accountCurrency == quoteCurrency {
		f = 1
		//fmt.Println(f, "Acc currency")
	} else if pairs.Find(legalPairs, accountCurrency+quoteCurrency) {
		if isJPY(p.Pair) {
			//
			f = (1 / (getLatestPrice(accountCurrency + quoteCurrency)) / 0.1)
			//fmt.Println(f, "Normal J")
		} else {
			f = (1 / getLatestPrice(accountCurrency+quoteCurrency))
			//fmt.Println(f, "Normal")
		}

	} else {
		if isJPY(p.Pair) {
			f = (1 * (getLatestPrice(accountCurrency + quoteCurrency)) / 0.1)
			//fmt.Println(f, "Reversed J")
		} else {
			f = getLatestPrice(quoteCurrency + accountCurrency)
			//fmt.Println(f, "Reversed")

		}

	}

	return f
}

func positions() []*position.Position {

	pos := []*position.Position{

		{Pair: "AUDCAD", Direction: 1, Volume: 0.01, Price: 0.91123, CurrentPrice: 0.90123},
		{Pair: "AUDCHF", Direction: 1, Volume: 0.01, Price: 0.56123, CurrentPrice: 0.55123},
		{Pair: "AUDJPY", Direction: 1, Volume: 0.01, Price: 93.123, CurrentPrice: 92.123},
		{Pair: "AUDNZD", Direction: 1, Volume: 0.01, Price: 1.11123, CurrentPrice: 1.10123},
		{Pair: "AUDUSD", Direction: 1, Volume: 0.01, Price: 0.63123, CurrentPrice: 0.62123},
		{Pair: "EURUSD", Direction: 1, Volume: 0.01, Price: 1.10123, CurrentPrice: 1.09123},
		{Pair: "EURAUD", Direction: 1, Volume: 0.01, Price: 1.72123, CurrentPrice: 1.73123},
		{Pair: "EURNZD", Direction: 1, Volume: 0.01, Price: 1.92123, CurrentPrice: 1.91123},
		{Pair: "EURCAD", Direction: 1, Volume: 0.01, Price: 1.58123, CurrentPrice: 1.57123},
		{Pair: "EURCHF", Direction: 1, Volume: 0.01, Price: 0.97123, CurrentPrice: 0.96123},
		{Pair: "EURJPY", Direction: 1, Volume: 0.01, Price: 161.123, CurrentPrice: 160.123},
		{Pair: "EURGBP", Direction: 1, Volume: 0.01, Price: 0.83123, CurrentPrice: 0.84123},
		{Pair: "GBPUSD", Direction: 1, Volume: 0.01, Price: 1.13123, CurrentPrice: 1.12123},
		{Pair: "GBPAUD", Direction: 1, Volume: 0.01, Price: 2.06123, CurrentPrice: 2.05123},
		{Pair: "GBPNZD", Direction: 1, Volume: 0.01, Price: 2.27123, CurrentPrice: 2.26123},
		{Pair: "GBPCAD", Direction: 1, Volume: 0.01, Price: 1.87123, CurrentPrice: 1.86123},
		{Pair: "GBPJPY", Direction: 1, Volume: 0.01, Price: 191.123, CurrentPrice: 190.123},
		{Pair: "USDJPY", Direction: 1, Volume: 0.01, Price: 148.123, CurrentPrice: 147.123},
		{Pair: "USDCAD", Direction: 1, Volume: 0.01, Price: 1.45123, CurrentPrice: 1.44123},
		{Pair: "USDCHF", Direction: 1, Volume: 0.01, Price: 0.89123, CurrentPrice: 0.88123},
		{Pair: "NZDUSD", Direction: 1, Volume: 0.01, Price: 0.58123, CurrentPrice: 0.57123},
		{Pair: "NZDCAD", Direction: 1, Volume: 0.01, Price: 0.83123, CurrentPrice: 0.82123},
		{Pair: "NZDCHF", Direction: 1, Volume: 0.01, Price: 0.51123, CurrentPrice: 0.50123},
		{Pair: "NZDJPY", Direction: 1, Volume: 0.01, Price: 85.123, CurrentPrice: 84.123},
		{Pair: "CHFJPY", Direction: 1, Volume: 0.01, Price: 168.123, CurrentPrice: 167.123},
		{Pair: "CADCHF", Direction: 1, Volume: 0.01, Price: 0.62123, CurrentPrice: 0.61123},
		{Pair: "CADJPY", Direction: 1, Volume: 0.01, Price: 103.123, CurrentPrice: 102.123},
	}

	return pos
}

var accCurrencies = []string{"USD", "EUR", "NZD"}

func TestPointValue(t *testing.T) {

	// conditions

	for _, v := range accCurrencies {
		for _, vv := range positions() {
			calcAFC(*vv, v, vv.Pair[:3])
		}

	}

}
