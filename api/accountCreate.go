package api

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/neotmp/go-trading/account"
)

func AccountCreate(c *fiber.Ctx) error {

	// Server Responds w/ data
	res := new(SRAccountCreate)

	// get account Id and brokerId from payload
	a := new(account.Account)

	if err := c.BodyParser(&a); err != nil {
		c.Status(503).Send([]byte(err.Error()))
		return err
	}

	a, err := a.Create()
	if err != nil {
		res.Error = fmt.Sprint(err)
		res.Message = fmt.Sprintf("Could not create account with given name: %s", a.Name)
		res.Code = 0
		return c.JSON(res)

	}

	// 0 - error, 1 - success at fetching, 2 - success at modifying/editing,
	// 3 - success at creating, 4 - success at deleting

	//Data structure in case of success
	res.Message = "Account has been successfully created."
	res.Code = 3
	res.Data = a

	return c.JSON(res)
}
