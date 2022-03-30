package mono

import (
	"fmt"

	"github.com/khrees2412/mono-sdk/models"
)

func (c *ConnectService) GetAssets(userID string) (interface{}, interface{}) {
	u := fmt.Sprintf("/accounts/%s/assets",userID)
	resp := &models.Asset{}
	err := c.client.Call("GET", u, nil, &resp)
	return resp, err
}
