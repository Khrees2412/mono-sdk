package mono

import "testing"

func TestSetBudget(t *testing.T) {
	b := new(BudgetReq)
	b.Amount = 10000
	_, err := cl.Issuing.SetBudget("", b)

	if err != nil {
		t.Error(err)
	}
}

func TestGetBudget(t *testing.T) {
	_, err := cl.Issuing.GetBudget("")

	if err != nil {
		t.Error(err)
	}
}
