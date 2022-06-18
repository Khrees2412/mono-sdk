package mono

import (
	"fmt"

	"github.com/khrees2412/mono-sdk/models"
)

type VirtualAccountReq struct {
	AccountHolder string
	AccountType   string
	Disposable    bool
}

func (c *IssuingService) CreateVirtualAccount(p *VirtualAccountReq) (*models.VirtualAccount, error) {
	u := subpath + "/virtualaccounts"
	resp := &models.VirtualAccount{}
	err := c.client.Call("POST", u, "", nil, &resp)
	return resp, err
}

func (c *IssuingService) GetVirtualAccount(acctID string) (*models.VirtualAccountRes, error) {
	u := subpath + fmt.Sprintf("/virtualaccounts/%s", acctID)
	resp := &models.VirtualAccountRes{}
	err := c.client.Call("GET", u, "", nil, &resp)
	return resp, err
}

func (c *IssuingService) GetAllVirtualAccounts() (*models.VirtualAccountsRes, error) {
	u := subpath + "/virtualaccounts"
	resp := &models.VirtualAccountsRes{}
	err := c.client.Call("GET", u, "", nil, &resp)
	return resp, err
}

func (c *IssuingService) GetVirtualAccountTxns(acctID string) (*models.Txns, error) {
	u := subpath + fmt.Sprintf("/virtualaccounts/%s", acctID)
	resp := &models.Txns{}
	err := c.client.Call("GET", u, "", nil, &resp)
	return resp, err
}

func (c *IssuingService) GetVirtualAccountBalance(acctID string) (*models.VirtualAccountBalance, error) {
	u := subpath + fmt.Sprintf("/virtualaccounts/%s/balance", acctID)
	resp := &models.VirtualAccountBalance{}
	err := c.client.Call("GET", u, "", nil, &resp)
	return resp, err
}

func (c *IssuingService) GetVirtualAccountKyc(acctID string) (*models.VirtualAccountKyc, error) {
	u := subpath + fmt.Sprintf("/virtualaccounts/%s/limits", acctID)
	resp := &models.VirtualAccountKyc{}
	err := c.client.Call("GET", u, "", nil, &resp)
	return resp, err
}

func (c *IssuingService) LiquidateAccount(acctID string) (Response, error) {
	u := subpath + fmt.Sprintf("/virtualaccounts/%s/liquidate", acctID)
	resp := Response{}
	err := c.client.Call("PATCH", u, "", nil, &resp)
	return resp, err
}

func (c *IssuingService) CloseVirtualAccount(acctID string) (Response, error) {
	u := subpath + fmt.Sprintf("/virtualaccounts/%s", acctID)
	resp := Response{}
	err := c.client.Call("DELETE", u, "", nil, &resp)
	return resp, err
}

func (c *IssuingService) UpgradeVirtualAccount(acctID string) (*models.Txns, error) {
	u := subpath + fmt.Sprintf("/virtualaccounts/%s/upgrade", acctID)
	resp := &models.Txns{}
	err := c.client.Call("PATCH", u, "", nil, &resp)
	return resp, err
}

type MockTxnBody struct {
	Amount string
	Type   string
}

func (c *IssuingService) MockVirtualAccountTxns(acctID string) (interface{}, interface{}) {
	u := subpath + fmt.Sprintf("/virtualaccounts/%s/mocktransaction", acctID)
	resp := &models.MockTxn{}
	err := c.client.Call("POST", u, "", nil, &resp)
	return resp, err
}
