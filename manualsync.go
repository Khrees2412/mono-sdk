package mono

import (
	"fmt"
)

/* This resource enables you the user data */
func (c *ConnectService) Sync(userID string) (interface{}, interface{}) {
	u := fmt.Sprintf("/accounts/%s/sync", userID)
	resp := Response{}
	err := c.client.Call("POST", u, nil, &resp)
	return resp, err
}

func (c *ConnectService) Reauthorise(userID string) (interface{}, interface{}) {
	u := fmt.Sprintf("/accounts/%s/reauthorise", userID)
	resp := Response{}
	err := c.client.Call("POST", u, nil, &resp)
	return resp, err
}
