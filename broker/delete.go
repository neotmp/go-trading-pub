package broker

func (b *Broker) Delete() error {

	// Should not be able to delete if any applies:
	// 1. Open position exists
	// 2. Non-zero balance exist
	// 3. Pending order exist - should just delete it as well

	err := b.dbDelete()
	if err != nil {
		return err
	}

	return nil
}
