package position

import (
	"database/sql"

	"github.com/neotmp/go-trading/database"
)

type SwapStructureTMP struct {
	AccountId       uint16  `jsob:"account_id"`
	AccountCurrency string  `jsob:"account_currency"`
	Pair            string  `jsob:"pair"`
	ContractId      uint16  `jsob:"contract_id"`
	PipSize         float32 `jsob:"pip_size"`
	SwapLongPips    float32 `jsob:"swap_long_pips"`
	SwapShortPips   float32 `jsob:"swap_short_pips"`
}

func DbGetAccountSpecs(p string) (*SwapStructureTMP, error) {

	a := new(SwapStructureTMP)

	q := `SELECT positions.account_id, positions.pair, accounts.contract_id, 
	accounts.currency AS account_currency, 
	account_specs.pip_size, account_specs.swap_long_pips, account_specs.swap_short_pips
	FROM positions 
	JOIN accounts 
	ON positions.account_id = accounts.id JOIN account_specs 
	ON accounts.contract_id = account_specs.account_contract_id 
	WHERE positions.pair = account_specs.symbol 
	AND positions.pair = $1 LIMIT 1`

	if err := database.DB.QueryRow(q,
		p).Scan(&a.AccountId,
		&a.Pair,
		&a.ContractId,
		&a.AccountCurrency,
		&a.PipSize,
		&a.SwapLongPips,
		&a.SwapShortPips,
	); err != nil {

		if err == sql.ErrNoRows {
			return nil, err
		}
		return nil, err
	}

	return a, nil

}
