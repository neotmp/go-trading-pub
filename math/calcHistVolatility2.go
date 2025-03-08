package math

import (
	"errors"
	"fmt"

	"github.com/neotmp/go-trading/pairs"
)

func getQuotes2(p string) ([]*pairs.CurrencyQuote, error) {

	q, err := DbGetQuotes(p)
	if err != nil {
		return []*pairs.CurrencyQuote{}, err
	}

	return q, nil

}

// CalcHistVolatility calculates historyc volatility
// and returns values
// it accepts params, slice of Quote and param for rolling
func CalcHistVolatility2(p string) ([]*HistoricVolatility, error) {

	data, err := getQuotes2(p)
	if err != nil {
		return []*HistoricVolatility{}, errors.New("couldn't get data to work with")
	}

	var hv = []*HistoricVolatility{}

	var year int
	var month int
	var high float32
	var low float32 = 1000

	for _, v := range data {
		var m int
		var y int

		var h float32
		var l float32

		m = int(v.Date.Month())
		y = v.Date.Year()
		h = v.High
		l = v.Low
		year = y

		if h > high {
			high = h
			fmt.Println(m, h, "m-h")

		}
		if l < low {
			low = l
			fmt.Println(year, m, l, "m-l")

		}

		//change of month
		if m != month {
			var d HistoricVolatility
			d.Max = high
			d.Min = low
			d.Year = uint16(year)
			d.Month = uint8(month)
			hv = append(hv, &d)

			low = v.Low
			high = v.High

			//fmt.Println("Year:", y, "Month:", m, "H:", high, "L:", low)

		}

		month = int(v.Date.Month())

	}

	return hv, nil
}
