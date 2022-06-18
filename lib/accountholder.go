package mono

import (
	"fmt"
	"mime/multipart"

	"github.com/khrees2412/mono-sdk/models"
)

type IssuingService service

var subpath = "/issuing/v1"

type AccountHolderReq struct {
	Entity    string
	FirstName string
	LastName  string
	Bvn       string
	Phone     string
	Address   struct {
		AddressLine1 string
		Lga          string
		City         string
		State        string
	}
	Identity struct {
		Type   string
		Number string
		Url    string
	}
}

func (c *IssuingService) CreateAccountHolder(p *AccountHolderReq) (*models.AccountHolder, error) {
	u := subpath + "/accountholders"
	resp := &models.AccountHolder{}
	err := c.client.Call("POST", u, "", nil, &resp)
	return resp, err
}

type SubAccountHolderReq struct {
	Main_account string
	// ID of the main account holder
	First_name string
	Last_name  string
	Phone      string
}

func (c *IssuingService) CreateSubAccountHolder(p *SubAccountHolderReq) (*models.AccountHolder, error) {
	u := subpath + "/accountholders"
	resp := &models.AccountHolder{}
	err := c.client.Call("POST", u, "", nil, &resp)
	return resp, err
}

func (c *IssuingService) DeleteAccountHolder(acctID string, p *SubAccountHolderReq) (Response, error) {
	u := subpath + fmt.Sprintf("/accountholders/%s", acctID)
	resp := Response{}
	err := c.client.Call("DELETE", u, "", nil, &resp)
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
func (c *IssuingService) UpdateAccountHolder(acctID string, p *AccountHolderReq) (*models.UpdateAccountHolder, error) {
	u := subpath + fmt.Sprintf("/accountholders/%s", acctID)
	resp := &models.UpdateAccountHolder{}
	err := c.client.Call("PATCH", u, "", nil, &resp)
	return resp, err
}

func (c *IssuingService) FetchAllAccounts() (*models.Accounts, error) {
	u := subpath + "/accountholders"
	resp := &models.Accounts{}
	err := c.client.Call("GET", u, "", nil, &resp)
	return resp, err

}

func (c *IssuingService) GetAccount(acctID string) (*models.Account, error) {
	u := subpath + fmt.Sprintf("/accountholders/%s", acctID)
	resp := &models.Account{}
	err := c.client.Call("GET", u, "", nil, &resp)
	return resp, err

}

func (c *IssuingService) UploadFile(file *multipart.FileHeader) (Response, error) {
	u := subpath + "/accountholders/upload"
	resp := Response{}
	err := c.client.Call("POST", u, "", file, &resp)
	return resp, err
}
