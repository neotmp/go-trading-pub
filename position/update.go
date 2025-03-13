package position

import (
	"fmt"

	"github.com/neotmp/go-trading/pairs"
)

func Update() error {

	// get all open positions by account

	p, err := ListPositionsByAccount(22)
	if err != nil {
		return err
	}

	for _, v := range p {

		// get latest prices
		pr := pairs.GetLatestPrice(v.Pair)
		// check for nil in case of error
		if pr == nil {
			return err
		}

		// set latest price
		v.CurrentPrice = pr.Close

		v, err := v.CalculateSwap()
		fmt.Println(v.Swap, "swap returned")
		if err != nil {
			return err
		}

		// calculate profit
		v, err = Profit(v)
		if err != nil {
			return err
		}

		// calculate change
		v, err = v.CalculateChange()
		if err != nil {
			return err
		}

		// update position
		v.dbUpdate()
	}

	return nil

}
