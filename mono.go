package mono

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"

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
	baseURL *url.URL
	secret  string

	Connect   *ConnectService
	Directpay *DirectpayService
	Issuing   *IssuingService
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
	c.baseURL, _ = url.Parse(baseEndpoint)

	c.common.client = c
	c.Connect = (*ConnectService)(&c.common)
	c.Directpay = (*DirectpayService)(&c.common)
	c.Issuing = (*IssuingService)(&c.common)

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
func (c *Client) Call(method, path string, body, v interface{}) error {
	var buf io.ReadWriter
	if body != nil {
		buf = new(bytes.Buffer)
		err := json.NewEncoder(buf).Encode(body)
		if err != nil {
			return err
		}
	}
	u, _ := c.baseURL.Parse(path)
	req, err := http.NewRequest(method, u.String(), buf)

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
		return newAPIError(httpResp)
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
