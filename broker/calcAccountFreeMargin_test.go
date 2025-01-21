package broker_test

import (
	"fmt"
	"testing"

	"github.com/neotmp/go-trading/broker"
)

func TestCalcAccountFreeMargin(t *testing.T) {

	var pos = []broker.Position{
		{Margin: 5.10595,
			Pair: "EURUSD"},

		{Margin: 5.14351,
			Pair: "EURGBP"},

		{Margin: 15.43053,
			Pair: "EURGBP"},
	}

	var got float32 = 112.18
	var expected float32 = 86.500015

	for _, v := range pos {

		got -= v.Margin

	}

	if fmt.Sprintf("%f", got) != fmt.Sprintf("%f", expected) {
		t.Errorf("wrong free margin in account, given %f: expected %f: ", got, expected)
	}

}
