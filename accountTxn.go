package mono

import (
	"fmt"

	"github.com/khrees2412/mono-sdk/models"
)

type TxnQuery struct {
	account_holder string
	account        string
	card           string
}

func (c *IssuingService) GetAccountTxns(acctID string) (interface{}, interface{}) {
	u := subpath + fmt.Sprintf("accountholders/%s/transactions", acctID)
	resp := &models.AccountTxn{}
	err := c.client.Call("GET", u, "", nil, &resp)
	return resp, err

}

func (c *IssuingService) GetAllTxns(t *TxnQuery) (interface{}, interface{}) {
	u := subpath + "/transactions"
	resp := &models.Txns{}
	err := c.client.Call("GET", u, "", nil, &resp)
	return resp, err
}

func (c *IssuingService) GetTxn(txnID string) (interface{}, interface{}) {
	u := subpath + fmt.Sprintf("/transaction/%s", txnID)
	resp := &models.Txn{}
	err := c.client.Call("GET", u, "", nil, &resp)
	return resp, err

}
