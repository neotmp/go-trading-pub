package broker_test

import (
	"testing"
	"time"

	"github.com/neotmp/go-trading/broker"
)

func broOpen(b *broker.Broker) *broker.Broker {

	return b

}

func TestBroOpen(t *testing.T) {

	b := broker.Broker{
		Name:     "New Broker",
		Country:  "Brazil",
		Phone:    "12345",
		Email:    "email",
		Memo:     "Some Cool Stuff here",
		OpenedAt: time.Now(),
		Status:   "Active",
	}

	newBr := broOpen(&b)

	if &b != newBr {
		t.Errorf("wrong broker struct returned, given %s: expected %s: ", b.Name, newBr.Name)
	}

}
