package position

func (p *Position) CalculateChange() (*Position, error) {

	// for buy direction
	if p.Direction == 1 {
		p.Change = ((p.CurrentPrice - p.Price) / p.Price) * 100
	} else {
		p.Change = ((p.Price - p.CurrentPrice) / p.CurrentPrice) * 100
	}

	return p, nil
}
