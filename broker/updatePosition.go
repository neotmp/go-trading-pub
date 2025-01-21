package broker

// UpdatePosition updates values in the fields and returns pointer to Broker.
// Values that may be updated:
// Volume (only in situations where account type is not HEDGE),
// SL &|| TP,
// TS &|| TP,
// TP || SL || TS,
// Profit,
// Memo,
// Change,
// SwapLongPips,
// SwapShortPips.
func (b *Broker) UpdatePosition(p *Position) (*Broker, error) {

	for _, v := range b.Positions {
		// TO DO Add other fields for updating position
		if v.Id == p.Id {
			//v.Volume = p.Volume
			v.SL = p.SL
			v.TS = p.TS
			v.TP = p.TP
			v.Profit = p.Profit
			v.Memo = p.Memo
			v.Change = p.Change
			v.SwapLongPips = p.SwapLongPips
			v.SwapShortPips = p.SwapShortPips

		}
	}

	updated, err := b.dbUpdatePosition(p)
	if err != nil {
		return b, err
	}

	return updated, nil

}
