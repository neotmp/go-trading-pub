package transaction

import "time"

type Transaction struct {
	Id         uint16    `json:"id"`
	Timestamp  time.Time `json:"timestamp"`
	BrokerId   uint16    `json:"brokerId" db:"broker_id"`
	AccountId  uint16    `json:"accountId" db:"account_id"`
	CurrencyId uint16    `json:"currencyId" db:"currency_id"`
	Direction  uint8     `json:"direction"` // 1- deposit, 2 - withdrawal, 0 - transfer
	Amount     float32   `json:"amount"`
	Memo       string    `json:"memo"`
	Status     uint8     `json:"status"`
	Fee        float32   `json:"fee"`
}

type TDeposit struct {
	Id         uint16    `json:"id"`
	Timestamp  time.Time `json:"timestamp"`
	BrokerId   uint16    `json:"brokerId" db:"broker_id"`
	AccountId  uint16    `json:"accountId" db:"account_id"`
	CurrencyId uint16    `json:"currencyId" db:"currency_id"`
	Direction  uint8     `json:"direction"` // 1- deposit, 2 - withdrawal, 0 - transfer
	Amount     float32   `json:"amount"`
	Memo       string    `json:"memo"`
}

type TWithdrawal struct {
	Id         uint16    `json:"id"`
	Timestamp  time.Time `json:"timestamp"`
	BrokerId   uint16    `json:"brokerId" db:"broker_id"`
	AccountId  uint16    `json:"accountId" db:"account_id"`
	CurrencyId uint16    `json:"currencyId" db:"currency_id"`
	Direction  uint8     `json:"direction"` // 1- deposit, 2 - withdrawal, 0 - transfer
	Amount     float32   `json:"amount"`
	Memo       string    `json:"memo"`
}

type TTransfer struct {
	Id         uint16    `json:"id"`
	Timestamp  time.Time `json:"timestamp"`
	BrokerId   uint16    `json:"brokerId" db:"broker_id"`
	DebitId    uint16    `json:"debitId" db:"debit_id"`   // account to debit
	CreditId   uint16    `json:"creditId" db:"credit_id"` // account to credit
	CurrencyId uint16    `json:"currencyId" db:"currency_id"`
	Amount     float32   `json:"amount"`
	Memo       string    `json:"memo"`
}
