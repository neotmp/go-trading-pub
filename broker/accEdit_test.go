package broker_test

import (
	"testing"

	"github.com/neotmp/go-trading/broker"
)

func editAccount(a *broker.Account) *broker.Account {
	return a

}

func TestAccountEdit(t *testing.T) {

	a := broker.Account{}

	a.Name = "New Name"
	a.Memo = "New Memo"
	a.Active = false

	editedAc := editAccount(&a)

	if a.Name != editedAc.Name {
		t.Errorf("failed to edit account, given name %s:, expected %s:", a.Name, editedAc.Name)
	}

	if a.Memo != editedAc.Memo {
		t.Errorf("failed to edit account, given memo %s:, expected %s:", a.Memo, editedAc.Memo)
	}

	if a.Active != editedAc.Active {
		t.Errorf("failed to edit account, given active %t:, expected %t:", a.Active, editedAc.Active)
	}

}
