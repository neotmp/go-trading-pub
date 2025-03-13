package position

import "github.com/neotmp/go-trading/database"

func (p *Position) dbUpdate() (*Position, error) {

	q := `UPDATE positions SET
	current_price = $1,
	profit = $2,
	change = $3,
	swap = $4
	WHERE id = $5 
	RETURNING id`

	if err := database.DB.QueryRow(q,
		&p.CurrentPrice,
		&p.Profit,
		&p.Change,
		&p.Swap,
		&p.Id).Scan(&p.Id); err != nil {
		return p, err
	}

	return p, nil
}
