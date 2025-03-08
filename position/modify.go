package position

// Modify amends position:
// Status if currenctly pending
// SL, TP, TS, if not canceled or rejected
// Memo - any status
func (p *Position) Modify() (*Position, error) {

	p, err := p.dbModify()
	if err != nil {
		return p, err
	}

	return p, nil

}
