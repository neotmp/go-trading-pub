package math

import (
	"errors"
	"fmt"

	"github.com/neotmp/go-trading/pairs"
)

func getQuotes(p string) ([]*pairs.CurrencyQuote, error) {

	q, err := DbGetQuotes(p)
	if err != nil {
		return []*pairs.CurrencyQuote{}, err
	}

	return q, nil

}

// CalcHistVolatility calculates historyc volatility
// and returns values
// it accepts params, slice of Quote and param for rolling
func CalcHistVolatility(p string) ([]*HistoricVolatility, error) {

	data, err := getQuotes(p)
	if err != nil {
		return []*HistoricVolatility{}, errors.New("couldn't get data to work with")
	}

	var hv = []*HistoricVolatility{}

	var year int
	var month int
	var high float32
	var low float32 = 1000

	for _, v := range data {
		var y int
		var m int
		var h float32
		var l float32
		var d HistoricVolatility

		//
		y = v.Date.Year()
		m = int(v.Date.Month())
		h = v.High
		l = v.Low

		if l < low {
			low = l
		}
		if h > high {
			high = h
		}

		d.Year = uint16(year)
		d.Month = uint8(month)
		d.Min = l
		d.Max = h
		d.Rolling = 30
		d.Abs = d.Max - d.Min

		//change year
		if year != y {
			year = y
			//fmt.Println(y, "Year")
		}
		//change month
		if month != m {
			fmt.Println(m, &d, "M, data")
			hv = append(hv, &d)
			// reset vars
			month = m
			high = h
			low = l
			//fmt.Println("Year:", year, "Month:", month, "High:", high, "Low:", low)

		}

	}

	return hv, nil
}
