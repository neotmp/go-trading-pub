package broker

func (b *Broker) OrdersFind() ([]*Order, error) {

	accs := b.Accounts
	var orders []*Order

	account, err := b.AccountFind(accs, 1)
	if err != nil {

		return orders, err
	}

	for _, v := range b.Orders {
		var o Order

		if v.AccountId == account.Id && b.Id == v.BrokerId {
			o = *v
			orders = append(orders, &o)
		}
	}

	return orders, nil

}
