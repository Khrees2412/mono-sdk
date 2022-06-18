package mono

import (
	"fmt"

	"github.com/khrees2412/mono-sdk/models"
)

// GetIdentity /* This resource provides a mini customer identity information */
func (c *ConnectService) GetIdentity(userID string) (*models.Identity, error) {
	u := fmt.Sprintf("/accounts/%s/identity", userID)
	resp := &models.Identity{}
	err := c.client.Call("GET", u, "", nil, &resp)
	return resp, err
}
