package lib

import (
	"errors"
	"net/http"
	"strings"
	"time"
)

type Client struct {
	c      *http.Client
	secret string
}
type Option func(c *Client)

func WithHTTPClient(cl *http.Client) Option {
	return func(c *Client) {
		c.c = cl
	}
}

func WithSecretKey(s string) Option {
	return func(c *Client) {
		c.secret = s
	}
}

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
func New(opts ...Option) (*Client, error) {
	c := &Client{}

	for _, opt := range opts {
		opt(c)
	}

	if IsStringEmpty(c.secret) {
		return nil, errors.New("please provide your secret key")
	}

	if c.c == nil {
		c.c = &http.Client{
			Transport: &xAPIKeyAuthTransport{
				originalTransport: http.DefaultTransport,
				secret:            c.secret,
			},
			Timeout: time.Second * 5,
		}
	}

	c.secret = ""

	return c, nil
}
