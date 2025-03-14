package main

import (
	"fmt"
	"time"

	"github.com/neotmp/go-trading/database"
	"github.com/neotmp/go-trading/position"
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

	// t := transaction.Transaction{
	// 	BrokerId:   10,
	// 	AccountId:  32,
	// 	DebitId:    34,
	// 	CreditId:   33,
	// 	CurrencyId: 1,
	// 	Direction:  2,
	// 	Amount:     10,
	// 	Timestamp:  time.Now(),
	// 	Status:     1,
	// 	//External:   true,
	// 	Fee: 0,
	// }

	// result, err := t.Transfer()
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// if result == nil {
	// 	fmt.Println("NIL")
	// }

	// fmt.Println(result, "R")

	// a := account.Account{
	// 	Id:         34,
	// 	BrokerId:   10,
	// 	Balance:    100,
	// 	ContractId: 1,
	// 	CurrencyId: 1,
	// 	Active: true,
	// }

	// res, err := position.AccountUpdate(&a)
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// fmt.Println(res)

	p := position.Position{
		BrokerId:  10,
		AccountId: 34,
		PairId:    22,
		Volume:    0.01,
		Timestamp: time.Now(),
		Direction: 1,
		Type:      1,
	}

	_, err := p.Create()
	if err != nil {
		fmt.Println(err)
	}

	//fmt.Println(open)

}
