package api

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/neotmp/go-trading/position"
)

func PositionClose(c *fiber.Ctx) error {

	// Server Responds w/ data
	res := new(SRPositionClose)

	// get order from payload
	p := new(position.Position)

	if err := c.BodyParser(&p); err != nil {
		c.Status(503).Send([]byte(err.Error()))
		return err
	}

	p, err := p.Close()
	// If we return nil, we don't have access to position values!!!!
	if err != nil {
		res.Error = fmt.Sprint(err)
		res.Message = fmt.Sprintf("Could not close position with broker id: %d, account id: %d, currency id: %d", p.BrokerId, p.AccountId, p.PairId)
		res.Code = 0
		return c.JSON(res)

	}

	// 0 - error, 1 - success at fetching, 2 - success at modifying/editing,
	// 3 - success at creating, 4 - success at deleting

	//Data structure in case of success
	res.Message = "Position has been successfully closed."
	res.Code = 3
	res.Data = p

	return c.JSON(res)
}
