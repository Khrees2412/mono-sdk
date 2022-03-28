package lib

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/khrees2412/mono-sdk/models"
	"github.com/khrees2412/mono-sdk/utils"
)

var baseEndpoint = "https://api.withmono.com"

func (c *Client) GetAccountId(code string) (interface{}, interface{}) {
	postBody, _ := json.Marshal(map[string]string{
		"code": code,
	})
	body := bytes.NewBuffer([]byte(postBody))
	req, err := http.NewRequest(http.MethodGet,
		fmt.Sprintf("%s/account/auth", baseEndpoint), body)
	if err != nil {
		return nil, err.Error()
	}
	resp, err := c.c.Do(req)
	if err != nil {
		return nil, err.Error()
	}

	defer resp.Body.Close()

	respBody, _ := ioutil.ReadAll(resp.Body)

	mono_id := utils.PrettyPrint(respBody)
	return mono_id, nil
}

// This resource represents the account details with the financial institution.
func (c *Client) GetAccountDetails(userID string) (interface{}, interface{}) {
	d := new(models.Details)

	r, err := http.NewRequest(http.MethodGet,
		fmt.Sprintf("%s/accounts/%s", baseEndpoint, userID), nil)
	if err != nil {
		return d, err.Error()
	}

	resp, err := c.c.Do(r)
	if err != nil {
		return nil, err.Error()
	}

	defer resp.Body.Close()
	respBody, _ := ioutil.ReadAll(resp.Body)

	details := utils.PrettyPrint(respBody)
	return details, nil

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
