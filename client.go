package bigcommerce

import (
	"bytes"
	"log"
	"net/http"
	"net/url"
)

type Client struct {
	BaseURL    *url.URL
	AuthToken  string
	httpClient *http.Client
	Version    string
}

func NewClient(version string, storeHash string, authToken string) Client {
	var client Client
	url, err := url.Parse("https://api.bigcommerce.com/stores/" + storeHash + "/v" + version)
	if err != nil {
		log.Fatal(err)
	}
	client.BaseURL = url
	client.AuthToken = authToken
	client.httpClient = http.DefaultClient
	client.Version = version
	return client
}

func (c *Client) configureRequest(httpMethod string, relativeUrl string, payload *bytes.Buffer) (*http.Request, error) {
	var req *http.Request
	var err error
	// I dont understand why this works...
	// solves a nil pointer dereference issue but not sure why
	if payload != nil {
		req, err = http.NewRequest(httpMethod, relativeUrl, payload)
	} else {
		req, err = http.NewRequest(httpMethod, relativeUrl, nil)
	}
	if err != nil {
		return nil, err
	}
	req.Header.Set("x-auth-token", c.AuthToken)
	req.Header.Set("Accept", "application/json")

	return req, nil
}

func (c *Client) Request(httpMethod string, relativeUrl string, payload *bytes.Buffer) (*http.Response, error) {
	req, err := c.configureRequest(httpMethod, relativeUrl, payload)
	if err != nil {
		return nil, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (client *Client) Get(url string) (*http.Response, error) {
	return client.Request("GET", url, nil)
}

func (client *Client) Put(url string, payload *bytes.Buffer) (*http.Response, error) {
	return client.Request("PUT", url, payload)
}

func (client *Client) Post(url string, payload *bytes.Buffer) (*http.Response, error) {
	return client.Request("POST", url, payload)
}

func (client *Client) Delete(url string) (*http.Response, error) {
	return client.Request("DELETE", url, nil)
}
