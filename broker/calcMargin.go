package broker

func (b *Broker) CalculateMargin(o *Order) float32 {

	//specs := getAccountSpecs(b.Accounts[0].ContractId, o.Pair)

	m := ((o.Volume * 100_000) * o.Price) / b.Accounts[0].Leverage

	// correct

	return m

}
