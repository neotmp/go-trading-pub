package api

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/neotmp/go-trading/transaction"
)

func Deposit(c *fiber.Ctx) error {

	// Server Responds w/ data
	res := new(SRTransaction)

	// get order from payload
	d := new(transaction.TDeposit)

	if err := c.BodyParser(&d); err != nil {
		c.Status(503).Send([]byte(err.Error()))
		return err
	}

	// account here
	d, _, err := d.Deposit()
	// If we return nil, we don't have access to position values!!!!
	if err != nil {
		res.Error = fmt.Sprint(err)
		res.Message = fmt.Sprintf("Could not transfer funds with broker id: %d, debit account id: %d, currency id: %d", d.BrokerId, d.AccountId, d.CurrencyId)
		res.Code = 0
		return c.JSON(res)

	}

	// 0 - error, 1 - success at fetching, 2 - success at modifying/editing,
	// 3 - success at creating, 4 - success at deleting

	//Data structure in case of success
	res.Message = "Deposit has been successfully processed."
	res.Code = 3
	res.Data = d

	return c.JSON(res)
}
