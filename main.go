package main

import (
	"fmt"
	"time"

	"github.com/neotmp/go-trading/broker"
	"github.com/neotmp/go-trading/database"
	"github.com/neotmp/go-trading/transaction"
)

func main() {

	database.Connect()

	// err := godotenv.Load()
	// if err != nil {
	// 	log.Fatalf("unable to load .env file: %e", err)
	// }

	// app := fiber.New()
	// app.Use(cors.New(cors.Config{
	// 	AllowHeaders:     "Origin,Content-Type,Accept,Content-Length,Accept-Language,Accept-Encoding,Connection,Access-Control-Allow-Origin",
	// 	AllowOrigins:     os.Getenv("ALLOW_ORIGINS"),
	// 	AllowCredentials: true,
	// 	AllowMethods:     "GET,POST,HEAD,PUT,DELETE,PATCH",
	// }))

	// routes.Setup(app)

	// if os.Getenv("ENV") == "dev" {
	// 	app.Listen(":3333")
	// } else {
	// 	log.Fatal(app.ListenTLS(":4753", "./cert.pem", "./cert.key"))
	// }

	br := broker.Broker{
		Id: 8,
		//Name:     "My Refactored Edited Primary Broker",
		//Country:  "Refactored Edited Zanzibar",
		//Phone:    "+983-234-234-2342",
		//Email:    "zanzibar@broker.br",
		//Memo:     "My cool refactored memo here.",
		//OpenedAt: time.Now().Format("2006-01-02"),
		//Active:   true,
	}

	// list, err := account.List(br.Id)
	// if err != nil {
	// 	fmt.Println(err)

	// }

	// fmt.Println(list, "Accounts list")

	// for _, v := range list {
	// 	fmt.Println(v.OpenedAt, "Accounts")

	// }

	//a := account.Account{
	//	Id: 11,
	//Name:       "Edited Refactored Trade account",
	//BrokerId:   br.Id,
	//Broker:     "Edited Refactored Primary Broker",
	//CurrencyId: 1,
	//Currency:   "USD",
	//Contract:   "Pro",
	//ContractId: 1,
	//Memo:       "Refactored memo edited",
	//Hedge:      false,
	//Type:       1,
	//Active:     true,
	//OpenedAt:   time.Now(),

	//}

	// tr := broker.Transaction{
	// 	Timestamp:  time.Now(),
	// 	BrokerId:   br.Id,
	// 	AccountId:  a.Id,
	// 	CurrencyId: a.CurrencyId,
	// 	Direction:  1,
	// 	Amount:     100,
	// 	Memo:       "My first deposit",
	// 	Fee:        0,
	// }

	// p := position.Position{
	// 	Id:         57,
	// 	Direction:  2,
	// 	Volume:     0.01,
	// 	Pair:       "EURUSD",
	// 	Timestamp:  time.Now(),
	// 	Price:      1.05123,
	// 	Memo:       "First order to fill",
	// 	PairId:     22,
	// 	BrokerId:   br.Id,
	// 	AccountId:  a.Id,
	// 	Type:       1,
	// 	Commission: -1.0,
	// 	Swap:       2.0,
	// 	Margin:     5,
	// }

	// cp, err := p.Close()
	// if err != nil {
	// 	fmt.Println(err, "Could not create account.")
	// }

	tran := transaction.TWithdrawal{
		Timestamp: time.Now(),
		BrokerId:  br.Id,
		AccountId: 11,

		CurrencyId: 1, // usd
		Amount:     500,
		Memo:       "Deposit funds",
	}

	tr, acc, err := tran.Withdraw()
	if err != nil {
		fmt.Println(err, "can't deposit funds to/from this account")
	}

	// if error returns nil as oyur object
	// if d == nil {
	// 	fmt.Println("Could not get account")
	// } else {
	// 	fmt.Println("Cash account:", d.Type == 0, d.Balance, "My Deposit")
	// }

	fmt.Println(tr, acc, "TR, ACC")

}
