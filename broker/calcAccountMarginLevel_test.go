package broker_test

import (
	"fmt"
	"testing"
)

func TestCalcAccountMarginLevel(t *testing.T) {

	// Margin Level = (Equity / Margin) * 100

	var equity float32 = 112.18
	var margin float32 = 25.68
	var got float32
	var expected float32 = 436.8380062

	got = (equity / margin) * 100

	if fmt.Sprintf("%f", got) != fmt.Sprintf("%f", expected) {
		t.Errorf("wrong margin level in account, given %f: expected %f:", got, expected)
	}

}
