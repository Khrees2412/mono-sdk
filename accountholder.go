package mono

import (
	"fmt"

	"github.com/khrees2412/mono-sdk/models"
)

type IssuingService service

var subpath = "/issuing/v1"

type AccountHolderReq struct {
	Entity     string
	First_name string
	Last_name  string
	Bvn        string
	Phone      string
	Address    struct {
		Address_line1 string
		Lga           string
		City          string
		State         string
	}
	Identity struct {
		Type   string
		Number string
		Url    string
	}
}

func (c *IssuingService) CreateAccountHolder(p *AccountHolderReq) (interface{}, interface{}) {
	u := subpath + "/accountholders"
	resp := &models.AccountHolder{}
	err := c.client.Call("POST", u, nil, &resp)
	return resp, err
}

type SubAccountHolderReq struct {
	Main_account string
	// ID of the main account holder
	First_name string
	Last_name  string
	Phone      string
}

func (c *IssuingService) CreateSubAccountHolder(p *SubAccountHolderReq) (interface{}, interface{}) {
	u := subpath + "/accountholders"
	resp := &models.AccountHolder{}
	err := c.client.Call("POST", u, nil, &resp)
	return resp, err
}

func (c *IssuingService) DeleteAccountHolder(acctID string, p *SubAccountHolderReq) (interface{}, interface{}) {
	u := subpath + fmt.Sprintf("/accountholders/%s", acctID)
	resp := Response{}
	err := c.client.Call("DELETE", u, nil, &resp)
	return resp, err
}

type UpdateAccountHolderReq struct {
	Main_account string
	// ID of the main account holder
	First_name string
	Last_name  string
	Phone      string
}

/*
type UpdateAccountHolderReq struct {
	Main_account string
	// ID of the main account holder
	First_name string
	Last_name  string
	Phone      string
}
*/
func (c *IssuingService) UpdateAccountHolder(acctID string, p *AccountHolderReq) (interface{}, interface{}) {
	u := subpath + fmt.Sprintf("/accountholders/%s", acctID)
	resp := &models.UpdateAccountHolder{}
	err := c.client.Call("PATCH", u, nil, &resp)
	return resp, err
}

func (c *IssuingService) FetchAllAccounts() (interface{}, interface{}) {
	u := subpath + "/accountholders"
	resp := &models.Accounts{}
	err := c.client.Call("GET", u, nil, &resp)
	return resp, err

}

func (c *IssuingService) GetAccount(acctID string) (interface{}, interface{}) {
	u := subpath + fmt.Sprintf("/accountholders/%s", acctID)
	resp := &models.Account{}
	err := c.client.Call("GET", u, nil, &resp)
	return resp, err

}
