package mono

import (
	"fmt"

	"github.com/khrees2412/mono-sdk/models"
)

type DirectpayService service

func (c *DirectpayService) MakeDirectPay(userID string) (interface{}, interface{}) {

	u := fmt.Sprintf("/accounts/%s/", userID)
	resp := &models.Income{}
	err := c.client.Call("GET", u, nil, &resp)

	return resp, err

}
