package mono

import (
	"fmt"

	"github.com/khrees2412/mono-sdk/models"
)

type CreateVirtualCardbody struct {
	AccountHolder string
	Currency      string
	Disposable    bool
	Amount        int32
}

type FundCardBody struct {
	Amount      string
	FundSources string
}

func (c *IssuingService) CreateVirtualCard(userID string, b CreateVirtualCardbody) (*models.VirtualCardCreate, error) {
	u := subpath + "/cards/virtual"
	resp := &models.VirtualCardCreate{}
	err := c.client.Call("POST", u, "", b, &resp)
	return resp, err
}

func (c *IssuingService) MockCardTxn(cardID, amount string) (Response, error) {
	u := subpath + fmt.Sprintf("/cards/%s/mocktransaction", cardID)
	resp := Response{}
	err := c.client.Call("POST", u, "", amount, &resp)
	return resp, err
}

func (c *IssuingService) FundVirtualCard(cardID string, b FundCardBody) (Response, error) {
	u := subpath + fmt.Sprintf("/cards/%s/fund", cardID)
	resp := Response{}
	err := c.client.Call("POST", u, "", b, &resp)
	return resp, err
}

func (c *IssuingService) LiquidateCard(cardID string) (Response, error) {
	u := subpath + fmt.Sprintf("/cards/%s/liquidate", cardID)
	resp := Response{}
	err := c.client.Call("PATCH", u, "", nil, &resp)
	return resp, err
}

func (c *IssuingService) DeleteCard(cardID string) (Response, error) {
	u := subpath + fmt.Sprintf("/cards/%s", cardID)
	resp := Response{}
	err := c.client.Call("DELETE", u, "", nil, &resp)
	return resp, err
}
