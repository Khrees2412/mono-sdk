package mono

import (
	"fmt"

	"github.com/khrees2412/mono-sdk/models"
)

type RequestSingleCardBody struct {
	AccountId string
	DesignId  string
	Address   struct {
		AddressLine1 string
		Lga          string
		City         string
		State        string
	}
}

type RequestBulkCardBody struct {
	Account []struct {
		AccountId string
		DesignId  string
	}
	Address struct {
		AddressLine1 string
		Lga          string
		City         string
		State        string
	}
}

/* This resource allows you to request a physical card */
func (c *IssuingService) RequestSingleCard(b *RequestSingleCardBody) (*models.RequestCard, error) {
	u := subpath + "/cards/physical"
	resp := &models.RequestCard{}
	err := c.client.Call("POST", u, "", b, &resp)
	return resp, err
}

// Not ready
func (c *IssuingService) RequestBulkCard(b *RequestBulkCardBody) (Response, error) {
	u := subpath + "/cards/physical/bulk"
	resp := Response{}
	err := c.client.Call("POST", u, "", b, &resp)
	return resp, err
}

func (c *IssuingService) GetCardDesign() (*models.CardDesign, error) {
	u := subpath + "/cards/designs"
	resp := &models.CardDesign{}
	err := c.client.Call("GET", u, "", nil, &resp)
	return resp, err
}

type PersonalizationBody struct {
	Name     string // not needed for update
	Quantity int32
	Artwork  string
}

func (c *IssuingService) MakePersonalization(b *PersonalizationBody) (*models.CardPersonalization, error) {
	u := subpath + "/cards/personalizations"
	resp := &models.CardPersonalization{}
	err := c.client.Call("POST", u, "", b, &resp)
	return resp, err
}

func (c *IssuingService) UpdatePersonalization(acctID string, b *PersonalizationBody) (*models.CardPersonalizationUpdate, error) {
	u := subpath + fmt.Sprintf("/cards/personalizations/%s", acctID)
	resp := &models.CardPersonalizationUpdate{}
	err := c.client.Call("PATCH", u, "", b, &resp)
	return resp, err
}

func (c *IssuingService) GetAllPersonalizations() (*models.Personalizations, error) {
	u := subpath + "/personalizations"
	resp := &models.Personalizations{}
	err := c.client.Call("GET", u, "", nil, &resp)
	return resp, err
}

func (c *IssuingService) GetPersonalization(id string) (*models.Personalization, error) {
	u := subpath + fmt.Sprintf("/personalizations/%s", id)
	resp := &models.Personalization{}
	err := c.client.Call("GET", u, "", nil, &resp)
	return resp, err
}
