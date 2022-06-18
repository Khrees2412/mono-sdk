package mono

import (
	"fmt"

	"github.com/khrees2412/mono-sdk/models"
)

type DirectpayService service

type DirectPayBody struct {
	Amount      string
	Description string
	Type        string
	Reference   string
	StartDate   string
	EndDate     string
	Duration    int32
	Interval    string
	Account     string
	RedirectUrl string
	Meta        struct {
		Reference string
	}
}

func (c *DirectpayService) MakeDirectPay(d *DirectPayBody) (*models.MakePayment, error) {
	u := "/v1/payments/initiate"
	resp := &models.MakePayment{}
	err := c.client.Call("POST", u, "", d, &resp)
	return resp, err

}

/* Using the reference passed when initiating payment */
type Ref struct {
	Reference string `json:"reference"`
}

func (c *DirectpayService) VerifyPayment(reference *Ref) (*models.VerifyPayment, error) {
	u := "/v1/payments/verify"
	resp := &models.VerifyPayment{}
	err := c.client.Call("POST", u, "", reference, &resp)
	return resp, err
}

// https://docs.mono.co/reference/get-payments
type PaymentReq struct {
	Start      string
	End        string
	CustomerID string
	AccountID  string
	Status     string
	Page       int
}

func (c *DirectpayService) GetAllPayments(p *PaymentReq) (*models.GetPayment, error) {
	u := "/v1/payments/transactions"
	var query = "?"
	if len(p.Start) > 1 {
		if len(p.End) > 1 {
			query += fmt.Sprintf("start=%s&end=%s", p.Start, p.End)
		}
	}
	resp := &models.GetPayment{}
	err := c.client.Call("GET", u, query, nil, &resp)
	return resp, err
}
