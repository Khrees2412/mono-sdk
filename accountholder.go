package mono

import (
	"github.com/khrees2412/mono-sdk/models"
)

type IssuingService service

var subpath = "/issuing/v1"

type AccountHolderReq struct {
}

func (c *IssuingService) CreateAccountHolder(p *AccountHolderReq) (interface{}, interface{}) {
	u := subpath + "/accountholders"
	resp := &models.AccountHolder{}
	err := c.client.Call("POST", u, nil, &resp)
	return resp, err
}

type SubAccountHolderReq struct {
}

func (c *IssuingService) CreateSubAccountHolder(p *SubAccountHolderReq) (interface{}, interface{}) {
	u := subpath + "/accountholders"
	resp := &models.AccountHolder{}
	err := c.client.Call("POST", u, nil, &resp)
	return resp, err
}
