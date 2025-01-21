package broker

import "fmt"

func (b *Broker) CalculatePositionSwap(p *Position) (*Broker, error) {

	as, err := GetAccountSpecs(p.AccountId, p.Pair)
	if err != nil {
		return b, err
	}

	fmt.Println(as)

	return b, nil
}
