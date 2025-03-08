package broker

import (
	"time"

	"github.com/neotmp/go-trading/account"
)

type Broker struct {
	Id       uint16 `json:"id"`
	Name     string `json:"name"`
	Country  string `json:"country"`
	Phone    string `json:"phone"`
	Email    string `json:"email"`
	Memo     string `json:"memo"`
	OpenedAt string `json:"openedAt" db:"opened_at"`
	EditedAt string `json:"editedAt" db:"edited_at"`
	Active   bool   `json:"active"`
	// AccountContracts  []*AccountContract `json:"accountContracts" db:"account_contracts"`
	// AccountCurrencies []*AccountCurrency `json:"accountCurrencies" db:"account_currencies"`
	Accounts []*account.Account `json:"accounts"`
	// Orders            []*Order           `json:"orders"`
	// Positions         []*Position        `json:"positions"`
	// Transactions      []*Transaction     `json:"transactions"`
}

type Transaction struct {
	Id         uint16    `json:"id"`
	Timestamp  time.Time `json:"timestamp"`
	BrokerId   uint16    `json:"brokerId" db:"broker_id"`     // fk
	AccountId  uint16    `json:"accountId" db:"account_id"`   // FK
	CurrencyId uint16    `json:"currencyId" db:"currency_id"` // fk
	Direction  uint8     `json:"direction"`                   // 1 deposit, 2 - withdrawal, 0 - transfer will create corresponding transaction
	Amount     float32   `json:"amount"`
	Memo       string    `json:"memo"`
	Status     uint8     `json:"status"` // 1 - executed, 2 - pending, 3 - rejected, 0 - canceled
	Fee        float32   `json:"fee"`
}

// type AccountSpecs struct {
// 	Id                uint16  `json:"id"`
// 	AccountContractId uint16  `json:"accountContractId" db:"account_contract_id"` // FK
// 	PairId            uint16  `json:"pairId" db:"pair_id"`
// 	Symbol            string  `json:"symbol"`
// 	PipSize           float32 `json:"pipSize" db:"pip_size"`
// 	SpreadPips        float32 `json:"spreadPips" db:"spread_pips"`
// 	SpreadPerLot      float32 `json:"spreadPerLot" db:"spread_per_lot"`
// 	SwapShortPips     float32 `json:"swapShortPips" db:"swap_short_pips"`
// 	SwapLongPips      float32 `json:"swapLongPips" db:"swap_long_pips"`
// 	Lot               uint32  `json:"lot"`
// }

// type AccountContract struct {
// 	Id       uint32 `json:"id"`
// 	Name     string `json:"name"`
// 	BrokerId uint32 `json:"broker_id"`
// 	Memo     string `json:"memo"`
// }

// type AccountCurrency struct {
// 	Id         uint32 `json:"id"`
// 	Symbol     string `json:"symbol"`
// 	ContractId uint32 `json:"contractId" db:"contract_id"`
// }

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
