package mono

import (
	"fmt"

	"github.com/khrees2412/mono-sdk/models"
)

type VirtualAcctTransferBody struct {
	Amount         string
	Reference      string
	Narration      string
	Account_number string
}

type BankTransferBody struct {
	Amount          string
	Reference       string
	Bank_code       string
	Narration       string
	Virtual_account string
	Meta            struct {
		Cust string
	}
}

type CreditVirtualAcctBody struct {
	Amount    string
	Narration string
}

func (c *IssuingService) MakeTransferToBank(acctID string, b *BankTransferBody) (interface{}, interface{}) {
	u := subpath + fmt.Sprintf("/virtualaccounts/%s/transfer", acctID)
	resp := &models.FundTransfer{}
	err := c.client.Call("POST", u, "", nil, &resp)
	return resp, err
}

func (c *IssuingService) MakeTransferToVirtualAcct(acctID string, b *VirtualAcctTransferBody) (interface{}, interface{}) {
	u := subpath + fmt.Sprintf("/virtualaccounts/%s/transfer", acctID)
	resp := &models.FundTransfer{}
	err := c.client.Call("POST", u, "", nil, &resp)
	return resp, err
}
func (c *IssuingService) CreditVirtualAcct(acctID string, b *CreditVirtualAcctBody) (interface{}, interface{}) {
	u := subpath + fmt.Sprintf("/virtualaccounts/%s/credit", acctID)
	resp := Response{}
	err := c.client.Call("POST", u, "", nil, &resp)
	return resp, err
}
