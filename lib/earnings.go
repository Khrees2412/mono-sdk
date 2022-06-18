package mono

import (
	"fmt"

	"github.com/khrees2412/mono-sdk/models"
)

/* This resource represents the earnings of your connected customers in their investment accounts. */
func (c *ConnectService) GetEarnings(userID string) (*models.Earnings, error) {
	u := fmt.Sprintf("/accounts/%s/earnings", userID)
	resp := &models.Earnings{}
	err := c.client.Call("GET", u, "", nil, &resp)
	return resp, err

}
