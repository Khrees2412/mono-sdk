package mono

import "github.com/khrees2412/mono-sdk/models"

type WalletReq struct {
	Amount   string
	Currency string
}

type WalletQuery struct {
	currency string
}

func (c *IssuingService) FundWallet() (interface{}, interface{}) {
	u := subpath + "/wallets/fund"
	resp := &models.Wallet{}
	err := c.client.Call("POST", u, nil, &resp)
	return resp, err
}

func (c *IssuingService) GetWallet(q *WalletQuery) (interface{}, interface{}) {
	u := subpath + "/wallets"
	resp := &models.WalletRes{}
	err := c.client.Call("GET", u, nil, &resp)
	return resp, err
}
