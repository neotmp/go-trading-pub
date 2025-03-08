package api

import (
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/neotmp/go-trading/broker"
)

// TO DO Broker Can only be deleted if:
// 1. Current balance = 0
// 2. No positions
// if brokeer has a history of positions they all should be closed
// and balance should be zero, then the broker cann be archived, i.e. active would be false
func BrokerDelete(c *fiber.Ctx) error {

	i := c.Params("*")

	// Server Responds w/ data
	res := new(SRBrokerEdit)

	// get account Id from payload
	b := new(broker.Broker)

	// check if we can convert string to int
	id, err := strconv.Atoi(i)
	if err != nil {
		res.Error = fmt.Sprint(err)
		res.Message = fmt.Sprintf("Could not convert given id: %d", b.Id)
		res.Code = 0
		return c.JSON(res)
	}

	b.Id = uint16(id)

	// handle delete function
	err = b.Delete()
	if err != nil {
		res.Error = fmt.Sprint(err)
		res.Message = fmt.Sprintf("Could not delete broker with given id: %d", b.Id)
		res.Code = 0
		return c.JSON(res)
	}

	// 0 - error, 1 - success at fetching, 2 - success at modifying/editing,
	// 3 - success at creating, 4 - success at deleting

	//Data structure in case of success
	res.Message = "Broker has been successfully deleted."
	res.Code = 4

	return c.JSON(res)
}
