package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/neotmp/go-trading/api"
)

func Setup(app *fiber.App) {

	// broker
	app.Get("/api/broker/list", api.BrokerList)
	app.Post("/api/broker/edit", api.BrokerEdit)
	app.Post("/api/broker/create", api.BrokerCreate)
	app.Get("/api/broker/delete/*", api.BrokerDelete)
}
