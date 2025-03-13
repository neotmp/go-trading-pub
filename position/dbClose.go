package position

import (
	"github.com/neotmp/go-trading/database"
)

func (p *Position) dbClose() (*Position, error) {

	q := `UPDATE positions SET
	open = false
	WHERE id = $1
	RETURNING id`

	if err := database.DB.QueryRow(q, &p.Id).Scan(&p.Id); err != nil {
		return nil, err
	}

	return p, nil

}
