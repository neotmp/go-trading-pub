package position

import "time"

type Position struct {
	Id         uint16    `json:"id"`                        // ticket
	BrokerId   uint16    `json:"brokerId" db:"broker_id"`   // FK
	AccountId  uint16    `json:"accountId" db:"account_id"` // FK
	OrderId    uint16    `json:"orderId" db:"order_id"`     // FK order that opened/closed/modify the position
	Pair       string    `json:"pair"`
	PairId     uint16    `json:"pairId" db:"pair_id"`
	Volume     float32   `json:"volume"`
	Timestamp  time.Time `json:"timestamp"`
	Type       uint8     `json:"type"` // 1- market, 2 - pending, 0 - modify // execution type
	Price      float32   `json:"price"`
	SL         float32   `json:"sl"` // Stop Loss
	TS         float32   `json:"ts"` // Trailing stop
	TP         float32   `json:"tp"` // Take profit
	Profit     float32   `json:"profit"`
	Memo       string    `json:"memo"`
	Change     float32   `json:"change"`
	Commission float32   `json:"commission"`
	Margin     float32   `json:"margin"`    // margin used to open positions
	Direction  uint8     `json:"direction"` //1 - buy, 2 - sell
	Swap       float32   `json:"swap"`

	// Joined from AccountSpecs on PairID
	PipSize       float32 `json:"pipSize" db:"pip_size"`
	SpreadPips    float32 `json:"spreadPips" db:"spread_pips"`
	SpreadPerLot  float32 `json:"spreadPerLot" db:"spread_per_lot"`
	SwapShortPips float32 `json:"swapShortPips" db:"swap_short_pips"`
	SwapLongPips  float32 `json:"swapLongPips" db:"swap_long_pips"`
	Base          string  `json:"base"`
	Quote         string  `json:"quote"`
}
