package broker

import (
	"github.com/neotmp/go-trading/database"
)

// dbUpdatePosition updates table in DB and returns pointer to Broker & error
func (b *Broker) dbPositionUpdate(p *Position) (*Broker, error) {

	q := `UPDATE positions SET
	volume = $1,
	sl = $2,
	ts = $3,
	tp = $4,
	profit = $5,
	memo = $6,
	change = $7,
	swap_long = $8,
	swap_short = $9
	WHERE id = $10 
	RETURNING id`

	if err := database.DB.QueryRow(q,
		&p.Volume,
		&p.SL,
		&p.TS,
		&p.TP,
		&p.Profit,
		&p.Memo,
		&p.Change,
		&p.SwapLongPips,
		&p.SwapShortPips,
		&p.Id).Scan(&p.Id); err != nil {
		return b, err
	}

	return b, nil
}
