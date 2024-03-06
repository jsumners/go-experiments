package client

import (
	"errors"
	"net/http"
)

type Client struct {
	http   *http.Client
	apiKey string
}

func New(apiKey string) (*Client, error) {
	return &Client{
		http:   &http.Client{Transport: http.DefaultTransport},
		apiKey: apiKey,
	}, nil
}

func (c *Client) Authenticate() error {
	req, err := http.NewRequest(http.MethodGet, "https://httpbin.org/bearer", nil)
	if err != nil {
		return err
	}

	req.Header.Set("Authorization", "Bearer "+c.apiKey)
	res, err := c.http.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode > 399 {
		return errors.New("authentication failed")
	}
	
	return nil
}
