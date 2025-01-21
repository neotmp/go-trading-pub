package broker_test

import (
	"fmt"
	"testing"

	"github.com/neotmp/go-trading/broker"
)

var pos = []broker.Position{
	{Profit: -0.29004,
		Pair: "EURUSD"},

	{Profit: 12.17989,
		Pair: "EURGBP"},
	{Profit: 0.00000,
		Pair: "EURGBP"},
}

func TestCalcAccountProfit(t *testing.T) {

	var got float32 = 0.0
	var expected float32 = 11.88985

	for _, v := range pos {

		got += v.Profit

	}

	if fmt.Sprintf("%f", got) != fmt.Sprintf("%f", expected) {
		t.Errorf("wrong profit in account, given %f: expected%f ", got, expected)
	}

}
