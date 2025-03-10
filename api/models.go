package api

import (
	"github.com/neotmp/go-trading/broker"
	"github.com/neotmp/go-trading/position"
)

// Data Shape of Broker List
type SRBrokerList struct {
	Code    uint8  `json:"code"`
	Message string `json:"message"`
	Error   string `json:"error"`
	Data    Data   `json:"data"`
}

// Data Shape of Broker Edit
type SRBrokerEdit struct {
	Code    uint8          `json:"code"`
	Message string         `json:"message"`
	Error   string         `json:"error"`
	Data    *broker.Broker `json:"data"`
}

// Data Shape of Broker Create
type SRBrokerCreate struct {
	Code    uint8          `json:"code"`
	Message string         `json:"message"`
	Error   string         `json:"error"`
	Data    *broker.Broker `json:"data"`
}

// Data Shape of Broker Delete
type SRBrokerDelete struct {
	Code    uint8  `json:"code"`
	Message string `json:"message"`
	Error   string `json:"error"`
	//Data    any    `json:"data"`
}

// Data Shape of Account Edit
type SRAccountEdit struct {
	Code    uint8  `json:"code"`
	Message string `json:"message"`
	Error   string `json:"error"`
	//Data    *broker.Account `json:"data"`
}

// Data Shape of Account Create
type SRAccountCreate struct {
	Code    uint8  `json:"code"`
	Message string `json:"message"`
	Error   string `json:"error"`
	//Data    *broker.Account `json:"data"`
}

// Data Shape of Broker Delete
type SRAccountDelete struct {
	Code    uint8  `json:"code"`
	Message string `json:"message"`
	Error   string `json:"error"`
	Data    any    `json:"data"`
}

// Data Shape of Account Edit
type SROrderEdit struct {
	Code    uint8  `json:"code"`
	Message string `json:"message"`
	Error   string `json:"error"`
	//Data    *broker.Order `json:"data"`
}

// Data Shape of Position Create
type SRPositionCreate struct {
	Code    uint8              `json:"code"`
	Message string             `json:"message"`
	Error   string             `json:"error"`
	Data    *position.Position `json:"data"`
}

type SRPositionClose struct {
	Code    uint8  `json:"code"`
	Message string `json:"message"`
	Error   string `json:"error"`
	Data    any    `json:"data"`
}

// Covers all cases where funds are transfered
type SRTransaction struct {
	Code    uint8  `json:"code"`
	Message string `json:"message"`
	Error   string `json:"error"`
	Data    any    `json:"data"`
}

type Data struct {
	Brokers []*broker.Broker `json:"brokers"`
	//Accounts     []*account.Account         `json:"accounts"`
	//Banks        []*bank.Bank               `json:"banks"`
	//Payees       []*payee.Payee             `json:"payees"`
	//Currencies   []*currency.Currency       `json:"currencies"`
	//Categories   []*category.Category       `json:"categories"`
	//Transactions []*transaction.Transaction `json:"transactions"`
	//Scheduled    []*schedule.Schedule       `json:"scheduled"`
	//AccountTypes []*account.AccountType     `json:"accountTypes"`
}
