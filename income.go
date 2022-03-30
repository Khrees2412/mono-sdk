package mono

import (
	"fmt"

	"github.com/khrees2412/mono-sdk/models"
)

/* This resource will return income information to the account. */
func (c *ConnectService) GetIncome(userID string) (interface{}, interface{}) {
	u := fmt.Sprintf("/accounts/%s/income", userID)
	resp := &models.Income{}
	err := c.client.Call("GET", u, nil, &resp)
	return resp, err
}
