package mono

import "testing"

func TestGetAccountId(t *testing.T) {
	id, err := cl.Connect.GetAccountId("")
	if err != nil {
		t.Error(err)
	}
	if id.Id == "" {
		t.Error("No id returned, empty string found")
	}
}

func TestGetAccountDetails(t *testing.T) {
	details, err := cl.Connect.GetAccountDetails("")
	if err != nil {
		t.Error(err)
	}
	if details.Account.AccountNumber == "" {
		t.Error("some account details are missing ")
	}
}

func TestUnlink(t *testing.T) {
	_, err := cl.Connect.Unlink("")
	if err != nil {
		t.Error(err)
	}

}
