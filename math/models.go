package math

type HistoricVolatility struct {
	Id      uint32  `json:"id"`
	Year    uint16  `json:"year"`
	Month   uint8   `json:"month"`
	Rolling uint8   `json:"rolling"` // 1M, 3M
	Max     float32 `json:"max"`
	Min     float32 `json:"min"`
	Abs     float32 `json:"abs"`
	PairId  uint16  `json:"pairId"`
}

type RollingHV struct {
	Id      uint32  `json:"id"`
	Date    uint16  `json:"date"`    // we calculate it for every date
	Rolling uint8   `json:"rolling"` // the size of the window
	Max     float32 `json:"max"`
	Min     float32 `json:"min"`
	Abs     float32 `json:"abs"` // the ABS difference between H & L
	PairId  uint16  `json:"pairId"`
}
