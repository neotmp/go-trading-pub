package broker

import (
	"time"
)

type Broker struct {
	Id        uint16      `json:"id"`
	Name      string      `json:"name"`
	Country   string      `json:"country"`
	Phone     string      `json:"phone"`
	Email     string      `json:"email"`
	Memo      string      `json:"memo"`
	OpenedAt  time.Time   `json:"opened_at"`
	EditedAt  string      `json:"edited_at"`
	Status    string      `json:"status"`
	Accounts  []*Account  `json:"accounts"`
	Orders    []*Order    `json:"orders"`
	Positions []*Position `json:"positions"`
}

type Order struct {
	Id         uint16    `json:"id"`
	BrokerId   uint16    `json:"broker_id"`
	AccountId  uint16    `json:"account_id"`
	Type       uint8     `json:"type"` // 1 - buy, 0 - sell
	Volume     float32   `json:"volume"`
	Pair       string    `json:"pair"`
	Timestamp  time.Time `json:"timestamp"`
	Price      float32   `json:"price"`
	SL         float32   `json:"sl"` // Stop Loss
	TS         float32   `json:"ts"` // Trailing stop
	TP         float32   `json:"tp"` // Take profit
	Memo       string    `json:"memo"`
	Status     uint8     `json:"status"` // 1 - Executed, 2 - Pending, 3- Rejected, 4 - Canceled
	PairId     uint16    `json:"pair_id"`
	Margin     float32   `json:"margin"` // margin used to open positions
	Commission float32   `json:"commission"`
	SpreadPips float32   `json:"spread_pips"`
}

type Position struct {
	Id         uint16    `json:"id"`         // ticket
	BrokerId   uint16    `json:"broker_id"`  // FK
	AccountId  uint16    `json:"account_id"` // FK
	OrderId    uint16    `json:"order_id"`   // FK order that opened/closed/modify the position
	Pair       string    `json:"pair"`
	PairId     uint16    `json:"pair_id"`
	Volume     float32   `json:"volume"`
	Timestamp  time.Time `json:"timestamp"`
	Type       uint8     `json:"type"` // 1 - buy, 2 - sell
	Price      float32   `json:"price"`
	SL         float32   `json:"sl"` // Stop Loss
	TS         float32   `json:"ts"` // Trailing stop
	TP         float32   `json:"tp"` // Take profit
	Profit     float32   `json:"profit"`
	Memo       string    `json:"memo"`
	Change     float32   `json:"change"`
	Commission float32   `json:"commission"`
	Margin     float32   `json:"margin"` // margin used to open positions

	// Joined from AccountSpecs on PairID
	PipSize       float32 `json:"pip_size"`
	SpreadPips    float32 `json:"spread_pips"`
	SpreadPerLot  float32 `json:"spread_per_lot"`
	SwapShortPips float32 `json:"swap_short_pips"`
	SwapLongPips  float32 `json:"swap_long_pips"`
	Base          string  `json:"base"`
	Quote         string  `json:"quote"`
}

type Account struct {
	Id          uint16    `json:"id"`
	BrokerId    uint16    `json:"broker_id"` //FK
	Balance     float32   `json:"balance"`
	Contract    string    `json:"contract"`
	Type        uint      `json:"type"`     // 1 - trade, 2 - transit, 0 - cash
	Currency    string    `json:"currency"` // USD
	Leverage    float32   `json:"leverage"` // 30, 50, 100, 200, 400, 800, etc
	Lot         uint32    `json:"lot"`      // 100_000
	StopOut     float32   `json:"stop_out"`
	Equity      float32   `json:"equity"`
	FreeMargin  float32   `json:"free_margin"`
	MarginLevel float32   `json:"margin_level"` //percent
	Margin      float32   `json:"margin"`       // margin as sum of all opened positions' margins
	OpenedAt    time.Time `json:"opened_at"`
	ClosedAt    string    `json:"closed_at"`
	Active      bool      `json:"active"`
	Hedge       bool      `json:"hedge"`
	Memo        string    `json:"memo"`
	ContractId  uint16    `json:"contract_id"` // FK
	CurrencyId  uint16    `json:"currency_id"` // FK
	Profit      float32   `json:"profit"`
	Name        string    `json:"name"`
}

type AccountSpecs struct {
	Id                uint16  `json:"id"`
	AccountContractId uint16  `json:"account_contract_id"` // FK
	PairId            uint16  `json:"pair_id"`
	Symbol            string  `json:"symbol"`
	PipSize           float32 `json:"pip_size"`
	SpreadPips        float32 `json:"spread_pips"`
	SpreadPerLot      float32 `json:"spread_per_lot"`
	SwapShortPips     float32 `json:"swap_short_pips"`
	SwapLongPips      float32 `json:"swap_long_pips"`
	Lot               uint32  `json:"lot"`
}

type AccountContract struct {
	Id       uint32 `json:"id"`
	Name     string `json:"name"`
	BrokerId uint32 `json:"broker_id"`
}

type AccountCurrency struct {
	Id         uint32 `json:"id"`
	Symbol     string `json:"symbol"`
	ContractId uint32 `json:"contract_id"`
}

// methods
// list all brokers
// get broker
// open broker
// close broker
// edit broker
// types of contracts
// specs for contracts
// open account
// close account
// calculate acc margin
// calc acc equity
// calc margin level
// clac free margin
// calc
// transfer funds to and from
// edit account
// list positions
// list orders
// get order
// modify order
// cancel order
//
