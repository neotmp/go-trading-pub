package pairs

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/neotmp/go-trading/database"
)

func GetLatestPrice(pair string) *CurrencyQuote {

	pair = strings.ToLower(pair)
	bc := pair[:3]
	qc := pair[3:]
	str := bc + "_" + qc

	fmt.Println(pair, "PAIR")

	qr := fmt.Sprintf(`SELECT id,
		date,
		open,
		high,
		low,
		close
		FROM %s
		ORDER BY id DESC LIMIT 1`, str)

	cq := new(CurrencyQuote)

	if err := database.DB.QueryRow(qr).Scan(&cq.Id,
		&cq.Date,
		&cq.Open,
		&cq.High,
		&cq.Low,
		&cq.Close,
	); err != nil {

		if err == sql.ErrNoRows {

			fmt.Println("No Currency with this code")

		}

		fmt.Println(err)

	}

	return cq

}
