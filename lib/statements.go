package mono

import (
	"fmt"

	"github.com/khrees2412/mono-sdk/models"
)

/* Gets the account statement in JSON */
func (c *ConnectService) GetStatementJSON(userID, period string) (interface{}, interface{}) {

	//Set output to JSON
	u := fmt.Sprintf("/accounts/%s/statement", userID)
	resp := &models.StatementJSON{}
	err := c.client.Call("GET", u, "", nil, &resp)
	return resp, err
}

/* Gets the account statement in PDF */
func (c *ConnectService) GetStatementPDF(userID, period string) (interface{}, interface{}) {
	u := fmt.Sprintf("/accounts/%s/statement", userID)
	resp := &models.StatementPDF{}
	err := c.client.Call("GET", u, "", nil, &resp)
	return resp, err
}

/* Gets the account statement status*/
func (c *ConnectService) GetStatementPollStatus(userID, jobID string) (interface{}, interface{}) {
	u := fmt.Sprintf("/accounts/%s/statement/jobs/%s", userID, jobID)
	resp := &models.StatementPDF{}
	err := c.client.Call("GET", u, "", nil, &resp)
	return resp, err
}
