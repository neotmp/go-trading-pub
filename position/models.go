package position

import "time"

type Position struct {
	Id           uint16    `json:"id"`                        // ticket
	BrokerId     uint16    `json:"brokerId" db:"broker_id"`   // FK
	AccountId    uint16    `json:"accountId" db:"account_id"` // FK
	Pair         string    `json:"pair"`
	PairId       uint16    `json:"pairId" db:"pair_id"`
	Volume       float32   `json:"volume"`
	Timestamp    time.Time `json:"timestamp"`
	Type         uint8     `json:"type"` // 1- market, 2 - pending, 0 - modify // execution type
	Price        float32   `json:"price"`
	SL           float32   `json:"sl"` // Stop Loss
	TS           float32   `json:"ts"` // Trailing stop
	TP           float32   `json:"tp"` // Take profit
	Profit       float32   `json:"profit"`
	Memo         string    `json:"memo"`
	Change       float32   `json:"change"`
	Commission   float32   `json:"commission"`
	Margin       float32   `json:"margin"`    // margin used to open positions
	Direction    uint8     `json:"direction"` //1 - buy, 2 - sell
	Swap         float32   `json:"swap"`
	CurrentPrice float32   `json:"currentPrice" db:"current_price"`
}
