package lib

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/khrees2412/mono-sdk/models"
	"github.com/khrees2412/mono-sdk/utils"
)

func (c *Client) GetAssets(userID string) (interface{}, interface{}) {
	a := new([]models.Asset)

	r, err := http.NewRequest(http.MethodGet,
		fmt.Sprintf("%s/accounts/%s/assets", baseEndpoint, userID), nil)
	if err != nil {
		return a, err.Error()
	}

	resp, err := c.c.Do(r)
	if err != nil {
		return nil, err.Error()
	}

	defer resp.Body.Close()
	respBody, _ := ioutil.ReadAll(resp.Body)

	assets := utils.PrettyPrint(respBody)
	return assets, nil
}
