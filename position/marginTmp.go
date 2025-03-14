package position

import "fmt"

type TmpStruct struct {
	PairId    int
	Direction int // 1 or 2
	Margin    float32
}

type countTmp struct {
	Count       int
	LongCount   int
	ShortCount  int
	MaxMargin   float32
	TotalMargin float32
}

func MarginTmp() {

	posTmp := []TmpStruct{
		{PairId: 22, Direction: 2, Margin: 5.16},
		{PairId: 22, Direction: 1, Margin: 5.15},
		{PairId: 21, Direction: 1, Margin: 4.5},
		{PairId: 21, Direction: 1, Margin: 4},
		{PairId: 22, Direction: 2, Margin: 5},
		{PairId: 22, Direction: 2, Margin: 5},
		{PairId: 22, Direction: 2, Margin: 5},
		{PairId: 22, Direction: 1, Margin: 5},
		{PairId: 22, Direction: 2, Margin: 5},
	}

	//

	var m = map[int]countTmp{}

	for _, v := range posTmp {

		i := new(countTmp)
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

	for _, v := range m {
		fmt.Println(float32(max(v.LongCount, v.ShortCount)) * v.MaxMargin)
	}

}
