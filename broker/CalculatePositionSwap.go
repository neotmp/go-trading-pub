package broker

import (
	"fmt"
	"time"
)

func (b *Broker) CalculatePositionSwap(p *Position) (*Broker, error) {

	as, err := GetAccountSpecs(p.AccountId, p.Pair)
	if err != nil {
		return b, err
	}

	// calculate number of days to apply swap
	// nd = now from

	orderPlaced := p.Timestamp

	//days := time.Since(orderPlaced)
	now := time.Now()

	dif := now.Sub(orderPlaced).Hours() / 24

	fmt.Println(now, orderPlaced, dif, "Days")

	fmt.Println(as)

	return b, nil
}
