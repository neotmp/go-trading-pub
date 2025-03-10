package api

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/neotmp/go-trading/transaction"
)

func Transfer(c *fiber.Ctx) error {

	// Server Responds w/ data
	res := new(SRTransaction)

	// get order from payload
	t := new(transaction.TTransfer)

	if err := c.BodyParser(&t); err != nil {
		c.Status(503).Send([]byte(err.Error()))
		return err
	}

	t, _, err := t.Transfer()
	// If we return nil, we don't have access to position values!!!!
	if err != nil {
		res.Error = fmt.Sprint(err)
		res.Message = fmt.Sprintf("Could not transfer funds with broker id: %d, debit account id: %d, currency id: %d", t.BrokerId, t.DebitId, t.CurrencyId)
		res.Code = 0
		return c.JSON(res)

	}

	// 0 - error, 1 - success at fetching, 2 - success at modifying/editing,
	// 3 - success at creating, 4 - success at deleting

	//Data structure in case of success
	res.Message = "Transfer has been successfully processed."
	res.Code = 3
	res.Data = t

	return c.JSON(res)
}
