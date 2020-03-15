package requests

import (
	"fmt"
	"log"
	"net/http"

	https "github.com/dougkirkley/requests/https"
)

var client *http.Client

// Certs type sctruct
type Certs struct {
	caCert   string
	keyFile  string
	certFile string
}

// Client type struct
type Client struct {
	client *http.Client
}

// New creates http.Client
func New(Certs *Certs) *Client {
    if Certs.caCert != "" {
        return &Client{
			client: https.LoadCert(),
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
		log.Fatal(err)
	}
	resp, err = c.client.Do(req)
	return resp, err
}

// Post performs post request
func (c *Client) Post(url string) (resp *http.Response, err error) {
	req, err := http.NewRequest(http.MethodPost, url, nil)
	if err != nil {
		log.Fatal(err)
	}
	resp, err = c.client.Do(req)
	return resp, err
}

// Put performs put request
func (c *Client) Put(url string) (resp *http.Response, err error) {
	req, err := http.NewRequest(http.MethodPut, url, nil)
	if err != nil {
		log.Fatal(err)
	}
	resp, err = c.client.Do(req)
	return resp, err
}