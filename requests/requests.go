package requests

import (
	"fmt"
	"log"
	"bytes"
	"net/http"

	https "github.com/dougkirkley/requests/https"
)


// Client type struct
type Client struct {
	client *http.Client
}

// NewClient creates http.Client
func NewClient(Certs *https.CertsInput) *Client {
    if Certs.caCert != "" {
        return &Client{
			client: https.LoadCerts(Certs),
		}
	}
	return &Client{
		client: &http.Client{},
	}
}

// Get performs get request
func (c *Client) Get(url string) (resp *http.Response, err error) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return req, err
	}
	resp, err = c.client.Do(req)
	return resp, err
}

// Post performs post request
func (c *Client) Post(url string, body *bytes.Buffer) (resp *http.Response, err error) {
	req, err := http.NewRequest(http.MethodPost, url, body)
	if err != nil {
		return req, err
	}
	resp, err = c.client.Do(req)
	return resp, err
}

// Put performs put request
func (c *Client) Put(url string) (resp *http.Response, err error) {
	req, err := http.NewRequest(http.MethodPut, url, nil)
	if err != nil {
		return req, err
	}
	resp, err = c.client.Do(req)
	return resp, err
}