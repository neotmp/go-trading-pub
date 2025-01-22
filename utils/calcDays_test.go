package utils_test

import (
	"testing"
	"time"

	"github.com/neotmp/go-trading/utils"
)

func TestCalcDays(t *testing.T) {

	type Date struct {
		Start time.Time
		End   time.Time
		Days  int
	}

	var dates = []Date{
		{Start: time.Date(2024, 10, 15, 0, 0, 0, 0, time.UTC),
			End:  time.Date(2025, 10, 15, 0, 0, 0, 0, time.UTC),
			Days: 365,
		},

		{Start: time.Date(2022, 10, 17, 0, 0, 0, 0, time.UTC),
			End:  time.Date(2023, 1, 17, 0, 0, 0, 0, time.UTC),
			Days: 92,
		},

		{Start: time.Date(2008, 2, 23, 0, 0, 0, 0, time.UTC),
			End:  time.Date(2021, 9, 17, 0, 0, 0, 0, time.UTC),
			Days: 4955,
		},

		{Start: time.Date(2024, 12, 22, 0, 0, 0, 0, time.UTC),
			End:  time.Date(2025, 1, 22, 0, 0, 0, 0, time.UTC),
			Days: 31,
		},
	}

	for _, v := range dates {

		if uint16(v.Days) != utils.CalcDays(v.Start, v.End) {
			t.Errorf("wrong number of days; given %d: expected %d:", utils.CalcDays(v.Start, v.End), v.Days)
		}

	}

}
