package broker

// Margin Level = (Equity / Margin) * 100
// CalculateAccountMarginLevel returns dynamic value of margin level based on current Equity
func (b *Broker) CalculateAccountMarginLevel(a *Account) (*Broker, error) {
	a.MarginLevel = (a.Equity / a.Margin) * 100

	updated, err := b.dbAccountUpdate(a)
	if err != nil {
		return b, err
	}

	return updated, nil
}
