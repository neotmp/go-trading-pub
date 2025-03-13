package broker

import (
	"github.com/neotmp/go-trading/account"
	"github.com/neotmp/go-trading/order"
	"github.com/neotmp/go-trading/position"
	"github.com/neotmp/go-trading/transaction"
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
	Accounts     []*account.Account         `json:"accounts"`
	Orders       []*order.Order             `json:"orders"`
	Positions    []*position.Position       `json:"positions"`
	Transactions []*transaction.Transaction `json:"transactions"`
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
