package order

// List returns a slice of pointers to all orders of broker
// Params id uint16 - broker id
func List(id uint16) ([]*Order, error) {

	orders, err := dbList(id)
	if err != nil {
		return nil, err
	}

	//b.Orders = append(b.Orders, orders...)

	return orders, nil

}
