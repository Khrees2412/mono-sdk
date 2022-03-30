package mono

import (
	"fmt"
	"net/http"
	

	"github.com/khrees2412/mono-sdk/models"
)

type Query struct {
	Start     string
	End       string
	Narration string
	Type      string
	Limit     string
	Paginate  string
}

/* This resource represents the known transactions on the account. */
func (c *ConnectService) GetTransactions(userID string, query *Query) (interface{}, interface{}) {
	r, err := http.NewRequest(http.MethodGet,
		fmt.Sprintf("%s/accounts/%s/transactions", baseEndpoint, userID), nil)
	
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
	u := fmt.Sprintf("/accounts/%s/transactions?start=%s&end=%s",userID)
	resp := &models.Transaction{}
	err = c.client.Call("GET", u, nil, &resp)
	return resp, err
}
