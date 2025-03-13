package position

import (
	"fmt"
	"time"

	"github.com/neotmp/go-trading/utils"
)

func (p *Position) CalculateSwap() (*Position, error) {
	// get swaps
	specs, err := DbGetAccountSpecs(p.Pair)
	if err != nil {
		return nil, err
	}

	//fmt.Println(specs, "specs")

	afc, err := PointValue(p.Pair, specs.AccountCurrency)

	if err != nil {
		return nil, err
	}

	// calculate full days
	//s := time.Date(2025, 03, 1, 0, 0, 0, 0, time.Local)
	e := time.Now()
	num := utils.CalcDays(p.Timestamp, e)

	fmt.Println(num, "swap days")

	if p.Direction == 1 {
		// special case for JPY
		// TO DO Change later to id
		if isJPY(p.Pair) {
			p.Swap = (specs.SwapLongPips * float32(num)) * afc
			//fmt.Println(afc, "afc")
		} else {
			p.Swap = (specs.SwapLongPips * 0.1 * float32(num)) * afc
		}

	} else {
		// special case for JPY
		// TO DO Change later to id
		if isJPY(p.Pair) {
			p.Swap = (specs.SwapShortPips * float32(num)) * afc
		} else {
			p.Swap = (specs.SwapShortPips * 0.1 * float32(num)) * afc
		}

	}

	return p, nil
}
