package mono

import (
	"fmt"

	"github.com/khrees2412/mono-sdk/models"
)

/* This resource enables you to sync the user data and obtain the latest */
func (c *ConnectService) Sync(userID string) (interface{}, interface{}) {
	u := fmt.Sprintf("/accounts/%s/sync", userID)
	resp := &models.Sync{}
	err := c.client.Call("POST", u, nil, &resp)
	return resp, err
}

/* This resource enables you request for reauthorisation to the user data */
func (c *ConnectService) Reauthorise(userID string) (interface{}, interface{}) {
	u := fmt.Sprintf("/accounts/%s/reauthorise", userID)
	resp := &models.Reauth{}
	err := c.client.Call("POST", u, nil, &resp)
	return resp, err
}
