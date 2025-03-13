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

	// t := position.Position{
	// 	Id:           50,
	// 	Type:         1,
	// 	BrokerId:     8,
	// 	Timestamp:    time.Now(),
	// 	AccountId:    22,
	// 	Price:        1.05123,
	// 	CurrentPrice: 1.06123, // price to close position
	// 	Volume:       0.01,
	// 	Pair:         "EURUSD",
	// 	PairId:       22,
	// 	SL:           0,
	// 	TS:           0,
	// 	TP:           0,
	// 	Memo:         "First Order to Fill",
	// 	Commission:   0,
	// 	Direction:    2,
	// 	Margin:       5,
	// 	Swap:         0,
	// }

	// result, err := t.Close()
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// fmt.Println(result, "R")

}
