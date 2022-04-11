package mono

import (
	"fmt"

	"github.com/khrees2412/mono-sdk/models"
)

type RequestSingleCardBody struct {
	Account_id string
	Design_id  string
	Address    struct {
		Address_line1 string
		Lga           string
		City          string
		State         string
	}
}

type RequestBulkCardBody struct {
	Account []struct {
		Account_id string
		Design_id  string
	}
	Address struct {
		Address_line1 string
		Lga           string
		City          string
		State         string
	}
}

/* This resource allows you to request a physical card */
func (c *IssuingService) RequestSingleCard(b *RequestSingleCardBody) (interface{}, interface{}) {
	u := subpath + "/cards/physical"
	resp := &models.RequestCard{}
	err := c.client.Call("POST", u, b, &resp)
	return resp, err
}

// Not ready
func (c *IssuingService) RequestBulkCard(b *RequestBulkCardBody) (interface{}, interface{}) {
	u := subpath + "/cards/physical/bulk"
	resp := Response{}
	err := c.client.Call("POST", u, b, &resp)
	return resp, err
}

func (c *IssuingService) GetCardDesign() (interface{}, interface{}) {
	u := subpath + "/cards/designs"
	resp := &models.CardDesign{}
	err := c.client.Call("GET", u, nil, &resp)
	return resp, err
}

type PersonalizationBody struct {
	Name     string // not needed for update
	Quantity int32
	Artwork  string
}

func (c *IssuingService) MakePersonalization(b *PersonalizationBody) (interface{}, interface{}) {
	u := subpath + "/cards/personalizations"
	resp := &models.CardPersonalization{}
	err := c.client.Call("POST", u, b, &resp)
	return resp, err
}

func (c *IssuingService) UpdatePersonalization(acctID string, b *PersonalizationBody) (interface{}, interface{}) {
	u := subpath + fmt.Sprintf("/cards/personalizations/%s", acctID)
	resp := &models.CardPersonalizationUpdate{}
	err := c.client.Call("PATCH", u, b, &resp)
	return resp, err
}

func (c *IssuingService) GetAllPersonalizations() (interface{}, interface{}) {
	u := subpath + "/personalizations"
	resp := &models.Personalizations{}
	err := c.client.Call("GET", u, nil, &resp)
	return resp, err
}

func (c *IssuingService) GetPersonalization(id string) (interface{}, interface{}) {
	u := subpath + fmt.Sprintf("/personalizations/%s", id)
	resp := &models.Personalization{}
	err := c.client.Call("GET", u, nil, &resp)
	return resp, err
}
