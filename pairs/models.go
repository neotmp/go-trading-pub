package pairs

import "time"

type CurrencyPair struct {
	Id     uint32 `json:"id"`
	Symbol string `json:"symbol"`
	Name   string `json:"name"`
}

type CurrencyQuote struct {
	Id    uint32    `json:"id"`
	Date  time.Time `json:"date"`
	Open  float32   `json:"open"`
	High  float32   `json:"high"`
	Low   float32   `json:"low"`
	Close float32   `json:"close"`
}
