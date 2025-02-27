package api

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/neotmp/go-trading/broker"
)

func BrokerList(c *fiber.Ctx) error {

	// Server Responds w/ data
	res := new(SRBrokerList)

	l, err := broker.BrokersList()
	if err != nil {
		res.Error = fmt.Sprint(err)
		res.Message = fmt.Sprint("Could not find list of brokers")
		res.Code = 0
		return c.JSON(res)

	}

	//Data structure in case of success
	res.Message = "Requested data has been successfully retrieved."
	res.Code = 1
	res.Data.Brokers = l

	return c.JSON(res)
}
