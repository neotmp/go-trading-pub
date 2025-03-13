package broker_test

import (
	"testing"

	"github.com/neotmp/go-trading/broker"
)

func editBroker(b *broker.Broker) *broker.Broker {
	return b

}

func TestEdit(t *testing.T) {

	b := broker.Broker{
		Name:    "Old Name",
		Country: "Old Country",
		Phone:   "12345",
		Email:   "email",
		Memo:    "Some Cool Stuff here",
		//OpenedAt: time.Now(), TO DO
		Active: true,
	}

	b.Name = "New Name"
	b.Country = "New Country"
	b.Phone = "New Phone"
	b.Email = "New Email"
	b.Memo = "New Memo"
	b.Active = true

	editedBr := editBroker(&b)

	if b.Name != editedBr.Name {
		t.Errorf("failed to edit broker, given name %s:, expected %s:", b.Name, editedBr.Name)
	}

	if b.Country != editedBr.Country {
		t.Errorf("failed to edit broker, given country %s:, expected %s:", b.Country, editedBr.Country)
	}

	if b.Phone != editedBr.Phone {
		t.Errorf("failed to edit broker, given phone %s:, expected %s:", b.Phone, editedBr.Phone)
	}

	if b.Email != editedBr.Email {
		t.Errorf("failed to edit broker, given email %s:, expected %s:", b.Email, editedBr.Email)
	}

	if b.Memo != editedBr.Memo {
		t.Errorf("failed to edit broker, given memo %s:, expected %s:", b.Memo, editedBr.Memo)
	}

	if b.Active != editedBr.Active {
		t.Errorf("failed to edit broker, given status %t:, expected %t:", b.Active, editedBr.Active)
	}

}
