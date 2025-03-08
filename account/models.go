package account

import "time"

type Account struct {
	Id          uint16    `json:"id"`
	BrokerId    uint16    `json:"brokerId" db:"broker_id"` //FK
	Broker      string    `json:"broker" db:"broker"`
	Balance     float32   `json:"balance"`
	Contract    string    `json:"contract"`
	Type        uint      `json:"type"`     // 1 - trade, 2 - transit, 0 - cash
	Leverage    float32   `json:"leverage"` // 30, 50, 100, 200, 400, 800, etc
	Lot         uint32    `json:"lot"`      // 100_000
	StopOut     float32   `json:"stopOut" db:"stop_out"`
	Equity      float32   `json:"equity"`
	FreeMargin  float32   `json:"freeMargin" db:"free_margin"`
	MarginLevel float32   `json:"marginLevel" db:"margin_level"` //percent
	Margin      float32   `json:"margin"`                        // margin as sum of all opened positions' margins
	OpenedAt    time.Time `json:"openedAt" db:"opened_at"`
	ClosedAt    string    `json:"closedAt" db:"closed_at"`
	Active      bool      `json:"active"`
	Hedge       bool      `json:"hedge"`
	Memo        string    `json:"memo"`
	ContractId  uint16    `json:"contractId" db:"contract_id"` // FK
	CurrencyId  uint16    `json:"currencyId" db:"currency_id"` // FK
	Currency    string    `json:"currency"`                    // USD JOIN
	Profit      float32   `json:"profit"`
	Name        string    `json:"name"`
	Change      float32   `json:"change"`
}

type AccountSpecs struct {
	Id                uint16  `json:"id"`
	AccountContractId uint16  `json:"accountContractId" db:"account_contract_id"` // FK
	PairId            uint16  `json:"pairId" db:"pair_id"`
	Symbol            string  `json:"symbol"`
	PipSize           float32 `json:"pipSize" db:"pip_size"`
	SpreadPips        float32 `json:"spreadPips" db:"spread_pips"`
	SpreadPerLot      float32 `json:"spreadPerLot" db:"spread_per_lot"`
	SwapShortPips     float32 `json:"swapShortPips" db:"swap_short_pips"`
	SwapLongPips      float32 `json:"swapLongPips" db:"swap_long_pips"`
	Lot               uint32  `json:"lot"`
}

type AccountContract struct {
	Id       uint32 `json:"id"`
	Name     string `json:"name"`
	BrokerId uint32 `json:"broker_id"`
	Memo     string `json:"memo"`
}

type AccountCurrency struct {
	Id         uint32 `json:"id"`
	Symbol     string `json:"symbol"`
	ContractId uint32 `json:"contractId" db:"contract_id"`
}
