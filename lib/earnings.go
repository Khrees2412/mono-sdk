package lib

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/khrees2412/mono-sdk/models"
	"github.com/khrees2412/mono-sdk/utils"
)

// This resource represents the earnings of your connected customers in their investment accounts.
func (c *Client) GetEarnings(userID string) (interface{}, interface{}) {
	e := new([]models.Earnings)

	r, err := http.NewRequest(http.MethodGet,
		fmt.Sprintf("%s/accounts/%s/earnings", baseEndpoint, userID), nil)
	if err != nil {
		return e, err.Error()
	}

	resp, err := c.c.Do(r)
	if err != nil {
		return nil, err.Error()
	}

	defer resp.Body.Close()
	respBody, _ := ioutil.ReadAll(resp.Body)

	earnings := utils.PrettyPrint(respBody)
	return earnings, nil
}
