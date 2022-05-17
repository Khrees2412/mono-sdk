package mono

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/khrees2412/mono-sdk/models"
)

type ConnectService service

type AcctId struct {
	Id string
}

func (c *ConnectService) GetAccountId(code string) (*AcctId, error) {

	postBody, _ := json.Marshal(map[string]string{
		"code": code,
	})
	body := bytes.NewBuffer([]byte(postBody))
	u := "/account/auth"
	resp := &AcctId{}
	err := c.client.Call("POST", u, "", body, &resp)
	return resp, err
}

/* This resource represents the account details with the financial institution. */
func (c *ConnectService) GetAccountDetails(userID string) (*models.Details, error) {
	u := fmt.Sprintf("/accounts/%s", userID)
	resp := &models.Details{}
	err := c.client.Call("GET", u, "", nil, &resp)
	return resp, err

}

// if resp.StatusCode > http.StatusCreated {
// 		var s struct {
// 			Message string `json:"message"`
// 		}
// 		if err := json.NewDecoder(resp.Body).Decode(&s); err != nil {
// 			return nil, err
// 		}
// 		return details, errors.New(s.Message)
// 	}
// 	return details, json.NewDecoder(resp.Body).Decode(details)

/* This enables you to provide your customers with the option to unlink their financial account(s) */
func (c *ConnectService) Unlink(userID string) (Response, error) {
	u := fmt.Sprintf("/accounts/%s/unlink", userID)
	resp := Response{}
	err := c.client.Call("POST", u, "", nil, &resp)
	return resp, err
}
