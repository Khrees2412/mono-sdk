package mono

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/khrees2412/mono-sdk/models"
	"github.com/khrees2412/mono-sdk/utils"
	"github.com/mitchellh/mapstructure"
)

var baseEndpoint = "https://api.withmono.com"

type service struct {
	client *Client
}

// Client manages communication with the Mono API
type Client struct {
	common  service
	c       *http.Client
	baseURL string
	secret  string

	Connect   *ConnectService
	Directpay *DirectpayService
	Issuing   *IssuingService
	Misc      *MiscService
}

// func WithHTTPClient(cl *http.Client) Option {
// 	return func(c *Client) {
// 		c.c = cl
// 	}
// }

// func WithSecretKey(s string) Option {
// 	return func(c *Client) {
// 		c.secret = s
// 	}
// }

func IsStringEmpty(s string) bool { return len(strings.TrimSpace(s)) == 0 }

type xAPIKeyAuthTransport struct {
	originalTransport http.RoundTripper
	secret            string
}

func (c *xAPIKeyAuthTransport) RoundTrip(r *http.Request) (*http.Response, error) {
	r.Header.Add("Content-Type", "application/json")
	r.Header.Add("mono-sec-key", c.secret)
	return c.originalTransport.RoundTrip(r)
}
func NewClient(secret string) (*Client, error) {
	c := &Client{}

	if IsStringEmpty(secret) {
		return nil, errors.New("please provide your secret key")
	}
	c.secret = secret
	c.baseURL = baseEndpoint

	c.common.client = c
	c.Connect = (*ConnectService)(&c.common)
	c.Directpay = (*DirectpayService)(&c.common)
	c.Issuing = (*IssuingService)(&c.common)
	c.Misc = (*MiscService)(&c.common)

	if c.c == nil {
		c.c = &http.Client{
			Transport: &xAPIKeyAuthTransport{
				originalTransport: http.DefaultTransport,
				secret:            c.secret,
			},
			Timeout: time.Second * 10,
		}
	}
	return c, nil
}

// Call actually does the HTTP request to Mono API
func (c *Client) Call(method, path string, query string, body, v interface{}) error {
	var buf io.ReadWriter
	var u string
	if body != nil {
		buf = new(bytes.Buffer)
		err := json.NewEncoder(buf).Encode(body)
		if err != nil {
			return err
		}
	}
	if len(query) > 1 {
		u = c.baseURL + path + "?" + query
	} else {
		u = c.baseURL + path

	}

	req, err := http.NewRequest(method, u, buf)

	if err != nil {
		return err
	}

	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	resp, err := c.c.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()
	return c.decodeResponse(resp, v)
}

// u, _ := c.baseURL.Parse(r.path)

func mapstruct(data interface{}, v interface{}) error {
	config := &mapstructure.DecoderConfig{
		Result:           v,
		TagName:          "json",
		WeaklyTypedInput: true,
	}
	decoder, err := mapstructure.NewDecoder(config)
	if err != nil {
		return err
	}
	err = decoder.Decode(data)
	return err
}

// Response represents arbitrary response data
type Response map[string]interface{}

// decodeResponse decodes the JSON response from the Mono API.
// The actual response will be written to the `v` parameter
func (c *Client) decodeResponse(httpResp *http.Response, v interface{}) error {
	var resp Response
	respBody, err := ioutil.ReadAll(httpResp.Body)
	json.Unmarshal(respBody, &resp)
	if err != nil {
		return err
	}
	if status, _ := resp["status"].(bool); !status || httpResp.StatusCode >= 400 {
		return utils.NewAPIError(httpResp)
	}

	if data, ok := resp["data"]; ok {
		switch t := resp["data"].(type) {
		case map[string]interface{}:
			return mapstruct(data, v)
		default:
			_ = t
			return mapstruct(resp, v)
		}
	}
	// if response data does not contain data key, map entire response to v
	return mapstruct(resp, v)
}

type BVN struct {
	Bvn string
}

/* This resource returns all the financial accounts that are linked to the BVN specified within Mono's Ecosystem. */
func (c *Client) View360(bvn *BVN) (interface{}, interface{}) {
	u := "/360view"
	resp := &models.View360{}
	err := c.Call("POST", u, "", bvn, &resp)
	p := &PaymentReq{
		Start: "",
		End:   "",
	}
	c.Directpay.GetAllPayments(p)
	return resp, err
}
