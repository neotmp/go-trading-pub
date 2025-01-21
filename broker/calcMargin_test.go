package broker_test

import (
	"fmt"
	"testing"

	"github.com/neotmp/go-trading/broker"
)

var leverage = float32(200)
var orders = []broker.Order{

	{Pair: "AUDCHF",

		Price:  0.56581,
		Volume: 0.01,
		Margin: 2.82905},

	{Pair: "USDCHF",
		Price:  0.91451,
		Volume: 0.01,
		Margin: 4.57255},

	{Pair: "USDJPY",
		Price:  156.259,
		Volume: 0.01,
		Margin: 781.295},
}

func TestCalcMargin(t *testing.T) {

	for _, v := range orders {

		m := (v.Price * (v.Volume * 100_000)) / leverage

		if fmt.Sprintf("%f", v.Margin) != fmt.Sprintf("%f", m) {
			t.Errorf("wrong in pair %s, given %f: %f * %f not equal %f", v.Pair, m, v.Price, v.Volume, v.Margin)
		}
	}

}
