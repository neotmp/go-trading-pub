package api

import "github.com/neotmp/go-trading/broker"

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
