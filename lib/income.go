package lib

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/khrees2412/mono-sdk/models"
	"github.com/khrees2412/mono-sdk/utils"
)

// This resource will return income information to the account.
func (c *Client) GetIncome(userID string) (interface{}, interface{}) {
	i := new(models.Income)

	r, err := http.NewRequest(http.MethodGet,
		fmt.Sprintf("v1/%s/accounts/%s/income", baseEndpoint, userID), nil)
	if err != nil {
		return i, err.Error()
	}

	resp, err := c.c.Do(r)
	if err != nil {
		return nil, err.Error()
	}

	defer resp.Body.Close()
	respBody, _ := ioutil.ReadAll(resp.Body)

	income := utils.PrettyPrint(respBody)
	return income, nil
}
