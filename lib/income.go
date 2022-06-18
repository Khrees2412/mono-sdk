package mono

import (
	"fmt"

	"github.com/khrees2412/mono-sdk/models"
)

// GetIncome /* This resource will return income information to the account. */
func (c *ConnectService) GetIncome(userID string) (*models.Income, error) {
	u := fmt.Sprintf("/accounts/%s/income", userID)
	resp := &models.Income{}
	err := c.client.Call("GET", u, "", nil, &resp)
	return resp, err
}
