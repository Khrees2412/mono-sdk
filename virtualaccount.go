package mono

import "github.com/khrees2412/mono-sdk/models"

type VirtualAccountReq struct {
	Account_holder string
	Account_type   string
	Disposable     bool
}

func (c *IssuingService) CreateVirtualAccount(p *VirtualAccountReq) (interface{}, interface{}) {
	u := subpath + "/virtualaccount"
	resp := &models.VirtualAccount{}
	err := c.client.Call("POST", u, nil, &resp)
	return resp, err
}
