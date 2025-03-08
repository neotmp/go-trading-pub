package broker_test

import (
	"testing"

	"github.com/neotmp/go-trading/broker"
)

func brokerOpen(b *broker.Broker) *broker.Broker {

	return b

}

func TestBroOpen(t *testing.T) {

	b := broker.Broker{
		Name:    "New Broker",
		Country: "Brazil",
		Phone:   "12345",
		Email:   "email",
		Memo:    "Some Cool Stuff here",
		//OpenedAt: time.Now(), TO DO
		Active: true,
	}

	newBr := brokerOpen(&b)

	if &b != newBr {
		t.Errorf("wrong broker struct returned, given %s: expected %s: ", b.Name, newBr.Name)
	}

}
