package broker

func (b *Broker) FindOrders() ([]*Order, error) {

	accs := b.Accounts
	var orders []*Order

	account, err := b.FindAccount(accs, 1)
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
