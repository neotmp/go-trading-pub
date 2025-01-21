package broker_test

import (
	"fmt"
	"testing"

	"github.com/neotmp/go-trading/broker"
)

func TestCalcAccountEquity(t *testing.T) {

	var pos = []broker.Position{
		{Pair: "EURUSD",
			Profit: -0.29004},

		{Pair: "EURGBP",
			Profit: 0.00000},

		{Pair: "EURGBP",
			Profit: 12.17989},
	}

	var expected float32 = 111.88985
	var got float32 = 100

	for _, v := range pos {
		got += v.Profit
	}

	if fmt.Sprintf("%f", got) != fmt.Sprintf("%f", expected) {
		t.Errorf("wrong equity in account, given %f: expected %f: ", got, expected)
	}

}
