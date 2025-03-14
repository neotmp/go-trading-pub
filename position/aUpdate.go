package position

import (
	"fmt"

	"github.com/neotmp/go-trading/account"
)

// AccountUpdate updates trade account: equity, margin, margin level, free margin,
// profit, change. Triggered by actions:
// 1. Open Position,
// 2. Modify Position,
// 3. Close Position,
// 4. Update Current prices,
// 5. Internal transfer.
func AccountUpdate(a *account.Account) (*account.Account, error) {

	fmt.Println(a.Balance, "FUNC BALANce")

	// 2. update margin
	a, p, err := AccountMargin(a)
	if err != nil {
		return nil, err
	}

	// 5. update profit
	a, err = AccountProfitAll(a, p)
	if err != nil {
		return nil, err
	}

	// 1. update equity
	a, err = AccountEquity(a)
	if err != nil {
		return nil, err
	}

	fmt.Println(a.Equity, "A Equity")

	// 3. update margin level
	a, err = AccountMarginLevel(a)
	if err != nil {
		return nil, err
	}

	// 4. update free margin
	a, err = AccountFreeMargin(a)
	if err != nil {
		return nil, err
	}

	// update change

	a, err = AccountChange(a)
	if err != nil {
		return nil, err
	}

	// 6. update db

	fmt.Println(a, "FUNC ACC")

	a, err = a.DbUpdate()
	if err != nil {
		return nil, err
	}

	return a, err
}
