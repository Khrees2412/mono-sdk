package mono

import (
	"fmt"

	"github.com/khrees2412/mono-sdk/models"
)

func (c *IssuingService) GetAllCards() (*models.Card, error) {
	u := subpath + "/cards"
	resp := &models.Card{}
	err := c.client.Call("GET", u, "", nil, &resp)
	return resp, err
}
func (c *IssuingService) GetCardDetails(cardID string) (*models.CardDetails, error) {
	u := subpath + fmt.Sprintf("/cards/%s", cardID)
	resp := &models.CardDetails{}
	err := c.client.Call("GET", u, "", nil, &resp)
	return resp, err
}
func (c *IssuingService) GetCardTxn(cardID string) (*models.CardTxn, error) {
	u := subpath + fmt.Sprintf("/cards/%s/transactions", cardID)
	resp := &models.CardTxn{}
	err := c.client.Call("GET", u, "", nil, &resp)
	return resp, err
}
func (c *IssuingService) FreezeCard(cardID string) (Response, error) {
	u := subpath + fmt.Sprintf("/cards/%s/freeze", cardID)
	resp := Response{}
	err := c.client.Call("PATCH", u, "", nil, &resp)
	return resp, err
}
func (c *IssuingService) UnFreezeCard(cardID string) (Response, error) {
	u := subpath + fmt.Sprintf("/cards/%s/unfreeze", cardID)
	resp := Response{}
	err := c.client.Call("PATCH", u, "", nil, &resp)
	return resp, err
}

type CardSpendLimitBody struct {
	Amount    string
	Channels  string
	Frequency int32
}

func (c *IssuingService) SetCardSpendLimit(cardID string, b *CardSpendLimitBody) (Response, error) {
	u := subpath + fmt.Sprintf("/cards/%s/limits", cardID)
	resp := Response{}
	err := c.client.Call("PATCH", u, "", b, &resp)
	return resp, err
}

func (c *IssuingService) GetCardSpendLimit(cardID string, b *CardSpendLimitBody) (*models.CardSpendLimit, error) {
	u := subpath + fmt.Sprintf("/cards/%s/limits", cardID)
	resp := &models.CardSpendLimit{}
	err := c.client.Call("GET", u, "", b, &resp)
	return resp, err
}
