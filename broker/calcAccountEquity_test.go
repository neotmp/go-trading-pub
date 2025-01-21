package broker_test

import (
	"fmt"
	"testing"
)

func TestCalcAccountEquity(t *testing.T) {

	var balance float32 = 100
	var expected float32 = 94.89
	var profit float32 = -5.110
	var got float32 = 0

	got = balance + profit

	if fmt.Sprintf("%f", got) != fmt.Sprintf("%f", expected) {
		t.Errorf("wrong equity in account, given %f: expected%f ", got, expected)
	}

}
