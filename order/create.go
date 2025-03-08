package order

// Create creates order in db and calls create position function.
// It ignores margin requirements and conducts no checks on the availability of funds
func (o *Order) Create() (*Order, error) {

	o, err := o.dbCreate()
	if err != nil {
		return o, err
	}

	return o, nil
}
