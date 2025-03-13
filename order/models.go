package order

import "time"

type Order struct {
	Id           uint16    `json:"id"`
	BrokerId     uint16    `json:"brokerId" db:"broker_id"`
	AccountId    uint16    `json:"accountId" db:"account_id"`
	Type         uint8     `json:"type"` // 1- market, 2 - pending, 0 - modify // execution type
	Volume       float32   `json:"volume"`
	Pair         string    `json:"pair"`
	Timestamp    time.Time `json:"timestamp"`
	TimestampTmp string    `json:"timestamp_tmp"`
	Price        float32   `json:"price"`
	SL           float32   `json:"sl"` // Stop Loss
	TS           float32   `json:"ts"` // Trailing stop
	TP           float32   `json:"tp"` // Take profit
	Memo         string    `json:"memo"`
	Status       uint8     `json:"status"` // 1 - Executed, 2 - Pending, 3- Rejected, 4 - Canceled
	PairId       uint16    `json:"pairId" db:"pair_id"`
	Margin       float32   `json:"margin"` // margin used to open positions
	Commission   float32   `json:"commission"`
	SpreadPips   float32   `json:"spreadPips" db:"spread_pips"`
	Direction    uint8     `json:"direction"` //1 - buy, 0 - sell
	PositionId   uint16    `json:"positionId" db:"position_id"`
}
