package lib

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/khrees2412/mono-sdk/models"
	"github.com/khrees2412/mono-sdk/utils"
)

// This resource provides a mini customer identity information
func (c *Client) GetIdentity(userID string) (interface{}, interface{}) {
	i := new(models.Identity)

	r, err := http.NewRequest(http.MethodGet,
		fmt.Sprintf("%s/accounts/%s/identity", baseEndpoint, userID), nil)
	if err != nil {
		return i, err.Error()
	}

	resp, err := c.c.Do(r)
	if err != nil {
		return nil, err.Error()
	}

	defer resp.Body.Close()
	respBody, _ := ioutil.ReadAll(resp.Body)

	identity := utils.PrettyPrint(respBody)
	return identity, nil
}
