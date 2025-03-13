package transaction

import (
	"time"

	"github.com/neotmp/go-trading/account"
)

type Transaction struct {
	Id         uint16    `json:"id"`
	Timestamp  time.Time `json:"timestamp"`
	BrokerId   uint16    `json:"brokerId" db:"broker_id"`
	AccountId  uint16    `json:"accountId" db:"account_id"`
	DebitId    uint16    `json:"debitId,omitempty"`  //for transfer only
	CreditId   uint16    `json:"creditId,omitempty"` // for transfer only
	CurrencyId uint16    `json:"currencyId" db:"currency_id"`
	Direction  uint8     `json:"direction,omitempty"` // 1- deposit, 2 - withdrawal, 0 - transfer, not used for transfer
	Amount     float32   `json:"amount"`
	Memo       string    `json:"memo"`
	Status     uint8     `json:"status"`
	External   bool      `json:"external"`
	Fee        float32   `json:"fee"`
}

type Data struct {
	Transactions []*Transaction     `json:"transactions,omitempty"`
	Accounts     []*account.Account `json:"accounts,omitempty"`
}
