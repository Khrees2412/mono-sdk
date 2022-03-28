package lib

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/khrees2412/mono-sdk/models"
	"github.com/khrees2412/mono-sdk/utils"
)

type Query struct {
	Start     string
	End       string
	Narration string
	Type      string
	Limit     string
	Paginate  string
}

func (c *Client) GetTransactions(userID string, query *Query) (interface{}, interface{}) {
	t := new(models.Transaction)

	r, err := http.NewRequest(http.MethodGet,
		fmt.Sprintf("%s/accounts/%s/transactions", baseEndpoint, userID), nil)
	if err != nil {
		return t, err.Error()
	}
	q := r.URL.Query() // Get a copy of the query values.

	if query.Start != "" {
		q.Add("start", query.Start)
	}
	if query.End != "" {
		q.Add("end", query.End)
	}
	if query.Narration != "" {
		q.Add("narration", query.Narration)
	}
	if query.Type != "" {
		q.Add("type", query.Type)
	}
	if query.Paginate != "" {
		q.Add("paginate", query.Paginate)
	}
	if query.Limit != "" {
		q.Add("limit", query.Limit)
	}
	resp, err := c.c.Do(r)
	if err != nil {
		return nil, err.Error()
	}

	defer resp.Body.Close()
	respBody, _ := ioutil.ReadAll(resp.Body)

	transactions := utils.PrettyPrint(respBody)
	return transactions, nil
}
