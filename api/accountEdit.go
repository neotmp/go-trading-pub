package api

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/neotmp/go-trading/account"
)

func AccountEdit(c *fiber.Ctx) error {

	// Server Responds w/ data
	res := new(SRAccountEdit)

	// get account Id and brokerId from payload
	a := new(account.Account)

	if err := c.BodyParser(&a); err != nil {
		c.Status(503).Send([]byte(err.Error()))
		return err
	}

	a, err := a.Edit()
	if err != nil {
		res.Error = fmt.Sprint(err)
		res.Message = fmt.Sprintf("Could not edit account with given id: %d", a.Id)
		res.Code = 0
		return c.JSON(res)

	}

	// 0 - error, 1 - success at fetching, 2 - success at modifying/editing,
	// 3 - success at creating, 4 - success at deleting

	//Data structure in case of success
	res.Message = "Account has been successfully updated."
	res.Code = 2
	res.Data = a

	return c.JSON(res)

}
