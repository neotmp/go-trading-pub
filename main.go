package main

import (
	"fmt"

	"github.com/neotmp/go-trading/broker"
	"github.com/neotmp/go-trading/database"
)

func main() {

	database.Connect()

	brs, err := broker.BrokersList()
	if err != nil {
		fmt.Println(err, "Problem with listing brokers")
	}

	for _, v := range brs {
		fmt.Println(v.Name)
	}

	// acc := broker.Account{
	// 	Contract:    "Pro",
	// 	Currency:    "USD",
	// 	Leverage:    800,
	// 	Lot:         100_000,
	// 	StopOut:     40,
	// 	OpenedAt:    time.Now(),
	// 	Active:      true,
	// 	Hedge:       true,
	// 	MarginLevel: 0,
	// 	BrokerId:    4,
	// 	Memo:        "New New Account",
	// 	ContractId:  1,
	// 	Type:        1,
	// 	CurrencyId:  2,
	// 	Name:        "EDITED account name",
	// }

	closeAcc, err := brs[0].AccountClose(brs[0].Accounts[3])
	if err != nil {
		fmt.Println("Error with closing account:", err)
	}

	for _, v := range closeAcc.Accounts {

		fmt.Println(v.Name, "Name")
		fmt.Println(v.Leverage, "Lev")
		fmt.Println(v.Active, "Active")

	}

}
