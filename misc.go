package mono

import (
	"fmt"

	"github.com/khrees2412/mono-sdk/models"
)

type MiscService service

func (c *MiscService) GetStates(state string) (interface{}, interface{}) {
	u := subpath + "/misc/states"
	resp := Response{}
	err := c.client.Call("GET", u, state, nil, &resp)
	return resp, err
}
func (c *MiscService) CacLookup(company string) (interface{}, interface{}) {
	u := "v1/cac/lookup"
	resp := &models.Cac{}
	err := c.client.Call("GET", u, company, nil, &resp)
	return resp, err
}
func (c *MiscService) CompanyShareholders(companyID string) (interface{}, interface{}) {
	u := fmt.Sprintf("v1/cac/company/%s", companyID)
	resp := &models.Shareholders{}
	err := c.client.Call("GET", u, "", nil, &resp)
	return resp, err
}

func (c *MiscService) GetBanks() (interface{}, interface{}) {
	u := subpath + "/misc/banks"
	resp := Response{}
	err := c.client.Call("GET", u, "", nil, &resp)
	return resp, err
}

type VerifyAccountQuery struct {
	Account_number string
	Bank_code      string
}

func (c *MiscService) VerifyAccount(b *VerifyAccountQuery) (interface{}, interface{}) {
	u := "/misc/verify/account"
	resp := &models.VerifyAcct{}
	err := c.client.Call("GET", u, "", b, &resp)
	return resp, err
}
