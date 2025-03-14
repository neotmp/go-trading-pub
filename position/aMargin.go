package position

import (
	"github.com/neotmp/go-trading/account"
)

type aMargin struct {
	Count       int
	LongCount   int
	ShortCount  int
	MaxMargin   float32
	TotalMargin float32
}

func AccountMargin(a *account.Account) (*account.Account, []*Position, error) {

	pos, err := dbListPositionsByAccount(a.Id)
	if err != nil {
		return nil, nil, err
	}

	var m = map[uint16]aMargin{}

	for _, v := range pos {

		i := new(aMargin)
		i.Count = m[v.PairId].Count + 1
		if v.Direction == 1 {
			i.LongCount = m[v.PairId].LongCount + 1
			i.ShortCount = m[v.PairId].ShortCount + 0
		} else {
			i.LongCount = m[v.PairId].LongCount + 0
			i.ShortCount = m[v.PairId].ShortCount + 1

		}

		i.MaxMargin = max(v.Margin, m[v.PairId].MaxMargin) // cool

		i.TotalMargin = m[v.PairId].TotalMargin + v.Margin

		m[v.PairId] = *i

	}

	var margin float32 = 0

	for _, v := range m {
		//fmt.Println(max(v.LongCount, v.ShortCount) * int(v.MaxMargin))

		margin += float32(max(v.LongCount, v.ShortCount)) * v.MaxMargin
	}

	a.Margin = margin

	return a, pos, nil

}
