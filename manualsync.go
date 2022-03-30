package mono

import (
	"fmt"

	"github.com/khrees2412/mono-sdk/models"
)

/* This resource enables you the user data */
func (c *ConnectService) Sync(userID string) (interface{}, interface{}) {
	u := fmt.Sprintf("/accounts/%s/sync", userID)
	resp := &models.Identity{}
	err := c.client.Call("POST", u, nil, &resp)
	return resp, err
}

func (c *ConnectService) Reauthorise(userID string) (interface{}, interface{}) {
	u := fmt.Sprintf("/accounts/%s/reauthorise", userID)
	resp := &models.Identity{}
	err := c.client.Call("POST", u, nil, &resp)
	return resp, err
}
