package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
	"github.com/neotmp/go-trading/database"
	"github.com/neotmp/go-trading/routes"
)

func main() {

	database.Connect()

	err := godotenv.Load()
	if err != nil {
		log.Fatalf("unable to load .env file: %e", err)
	}

	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowHeaders:     "Origin,Content-Type,Accept,Content-Length,Accept-Language,Accept-Encoding,Connection,Access-Control-Allow-Origin",
		AllowOrigins:     os.Getenv("ALLOW_ORIGINS"),
		AllowCredentials: true,
		AllowMethods:     "GET,POST,HEAD,PUT,DELETE,PATCH",
	}))

	routes.Setup(app)

	if os.Getenv("ENV") == "dev" {
		app.Listen(":3333")
	} else {
		log.Fatal(app.ListenTLS(":4753", "./cert.pem", "./cert.key"))
	}

	// database.Connect()

	// brs, err := broker.BrokersList()
	// if err != nil {
	// 	fmt.Println(err, "Problem with listing brokers")
	// }

	// for _, v := range brs {
	// 	fmt.Println(v.Name)
	// }

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

	// closeAcc, err := brs[0].AccountClose(brs[0].Accounts[3])
	// if err != nil {
	// 	fmt.Println("Error with closing account:", err)
	// }

	// for _, v := range closeAcc.Accounts {

	// 	fmt.Println(v.Name, "Name")
	// 	fmt.Println(v.Leverage, "Lev")
	// 	fmt.Println(v.Active, "Active")

	// }

}
