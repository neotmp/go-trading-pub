package account

import (
	"database/sql"
	"fmt"

	"github.com/neotmp/go-trading/database"
)

// GetAccountSpecs returns specs for specific account based on ContactId
func getAccountSpecs(accId uint16, pair string) (*AccountSpecs, error) {

	q := `SELECT id,
	account_contract_id,
	symbol,
	pip_size,
	spread_pips,
	spread_per_lot,
	swap_short_pips,
	swap_long_pips,
	pair_id, 
	lot
	FROM account_specs WHERE account_contract_id = (SELECT contract_id FROM accounts WHERE id = $1) 
	AND symbol = $2  `

	as := new(AccountSpecs)

	if err := database.DB.QueryRow(q, accId, pair).Scan(&as.Id,
		&as.AccountContractId,
		&as.Symbol,
		&as.PipSize,
		&as.SpreadPips,
		&as.SpreadPerLot,
		&as.SwapShortPips,
		&as.SwapLongPips,
		&as.PairId,
		&as.Lot,
	); err != nil {

		if err == sql.ErrNoRows {

			fmt.Println("No Account specs with this code")

			return as, err

		}

		return as, err

	}

	return as, nil

}
