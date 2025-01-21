package broker_test

import (
	"fmt"
	"testing"

	"github.com/neotmp/go-trading/broker"
)

var positions = []broker.Position{
	{Margin: 15.43053,
		Volume: 0.03,
		Pair:   "EURGBP",
		Id:     38,
	},
	{Margin: 5.14351,
		Volume: 0.01,
		Pair:   "EURGBP",
		Id:     37,
	},
	{Margin: 5.10595,
		Volume: 0.01,
		Pair:   "EURUSD",
		Id:     36,
	},
}

func TestCalcAccountMargin(t *testing.T) {

	var got float32 = 0.0
	var expected float32 = 25.679989

	for _, v := range positions {

		got += v.Margin

	}

	if fmt.Sprintf("%f", got) != fmt.Sprintf("%f", expected) {
		t.Errorf("wrong margin in account, given %f: expected%f ", got, expected)
	}

}
