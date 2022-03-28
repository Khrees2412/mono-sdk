package lib

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/khrees2412/mono-sdk/models"
	"github.com/khrees2412/mono-sdk/utils"
)

func (c *Client) GetStatementJSON(userID, period string) (interface{}, interface{}) {
	s := new(models.StatementJSON)

	r, err := http.NewRequest(http.MethodGet,
		fmt.Sprintf("%s/accounts/%s/statement", baseEndpoint, userID), nil)
	if err != nil {
		return s, err.Error()
	}

	resp, err := c.c.Do(r)
	if err != nil {
		return nil, err.Error()
	}

	defer resp.Body.Close()
	respBody, _ := ioutil.ReadAll(resp.Body)

	statement := utils.PrettyPrint(respBody)
	return statement, nil
}

func (c *Client) GetStatementPDF(userID, period string) (interface{}, interface{}) {
	s := new(models.StatementPDF)

	r, err := http.NewRequest(http.MethodGet,
		fmt.Sprintf("%s/accounts/%s/statement", baseEndpoint, userID), nil)
	if err != nil {
		return s, err.Error()
	}

	resp, err := c.c.Do(r)
	if err != nil {
		return nil, err.Error()
	}

	defer resp.Body.Close()
	respBody, _ := ioutil.ReadAll(resp.Body)

	statement := utils.PrettyPrint(respBody)
	return statement, nil
}

func (c *Client) GetStatementPollStatus(userID, jobID string) (interface{}, interface{}) {
	s := new(models.StatementPDF)

	r, err := http.NewRequest(http.MethodGet,
		fmt.Sprintf("%s/accounts/%s/statement/jobs/%s", baseEndpoint, userID, jobID), nil)
	if err != nil {
		return s, err.Error()
	}

	resp, err := c.c.Do(r)
	if err != nil {
		return nil, err.Error()
	}

	defer resp.Body.Close()
	respBody, _ := ioutil.ReadAll(resp.Body)

	status := utils.PrettyPrint(respBody)
	return status, nil
}
