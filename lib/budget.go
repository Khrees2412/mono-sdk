package mono

import (
	"fmt"

	"github.com/khrees2412/mono-sdk/models"
)

type BudgetReq struct {
	Amount int
}

func (c *IssuingService) SetBudget(userID string, r *BudgetReq) (*models.Budget, error) {
	u := subpath + fmt.Sprintf("/virtualaccounts/%s/budget/set", userID)
	resp := &models.Budget{}
	err := c.client.Call("PATCH", u, "", r, &resp)
	return resp, err
}

func (c *IssuingService) RemoveBudget(userID string, r *BudgetReq) (*models.Budget, error) {
	u := subpath + fmt.Sprintf("/virtualaccounts/%s/budget/remove", userID)
	resp := &models.Budget{}
	err := c.client.Call("PATCH", u, "", r, &resp)
	return resp, err
}

func (c *IssuingService) GetBudget(userID string) (*models.Budget, error) {
	u := subpath + fmt.Sprintf("/virtualaccounts/%s/budget", userID)
	resp := &models.Budget{}
	err := c.client.Call("PATCH", u, "", nil, &resp)
	return resp, err
}
