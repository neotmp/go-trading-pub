package api

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/neotmp/go-trading/broker"
)

func BrokerEdit(c *fiber.Ctx) error {
	// Server Responds w/ data
	res := new(SRBrokerEdit)

	// get account Id from payload
	b := new(broker.Broker)

	if err := c.BodyParser(&b); err != nil {
		c.Status(503).Send([]byte(err.Error()))
		return err
	}

	b, err := b.Edit()
	if err != nil {
		res.Error = fmt.Sprint(err)
		res.Message = fmt.Sprintf("Could not edit broker with given id: %d", b.Id)
		res.Code = 0
		return c.JSON(res)

	}

	// 0 - error, 1 - success at fetching, 2 - success at modifying/editing,
	// 3 - success at creating, 4 - success at deleting

	//Data structure in case of success
	res.Message = "Broker has been successfully updated."
	res.Code = 2
	res.Data = b

	return c.JSON(res)
}
