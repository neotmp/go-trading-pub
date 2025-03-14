package position

import "github.com/neotmp/go-trading/pairs"

func updateAllPositions(accId uint16) error {

	// get all open positions by account
	p, err := ListPositionsByAccount(accId)
	if err != nil {
		return err
	}

	for _, v := range p {

		//fmt.Println(v, "position")

		// get latest prices
		//fmt.Println(v.Pair, "v.pair")
		pr, err := pairs.GetLatestPrice(v.Pair)
		//check for nil in case of error
		if pr == nil || err != nil {
			return err
		}

		// set latest price
		v.CurrentPrice = pr.Close

		v, err := v.CalculateSwap()
		//fmt.Println(v.Swap, "swap returned")
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
