package account_test

import (
	"testing"

	"github.com/neotmp/go-trading/account"
)

func editAccount(a *account.Account) *account.Account {
	return a

}

func TestAccountEdit(t *testing.T) {

	a := account.Account{}

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
