package bigcommerce

import (
	"bytes"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

type Client struct {
	BaseURL    *url.URL
	AuthToken  string
	StoreHash  string
	httpClient *http.Client
	Version    int
}

func NewClient(storeHash string, authToken string, version int) Client {
	_version := strconv.Itoa(version)
	var client Client
	url, err := url.Parse("https://api.bigcommerce.com/stores/" + storeHash + "/v" + _version)
	if err != nil {
		log.Fatal(err)
	}
	client.BaseURL = url
	client.AuthToken = authToken
	client.StoreHash = storeHash
	client.httpClient = http.DefaultClient
	client.Version = version
	return client
}

func (c *Client) configureRequest(httpMethod string, relativeUrl string, payload []byte) (*http.Request, error) {
	req, err := http.NewRequest(httpMethod, relativeUrl, bytes.NewBuffer(payload))
	if err != nil {
		return nil, err
	}

	req.Header.Set("x-auth-token", c.AuthToken)
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")

	return req, nil
}

func (c *Client) Request(httpMethod string, relativeUrl string, payload []byte, attempt int) (*http.Response, error) {
	req, err := c.configureRequest(httpMethod, relativeUrl, payload)
	if err != nil {
		return nil, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	if (resp.StatusCode == 429 || resp.StatusCode/100 == 5) && attempt < 3 {
		log.Printf("%s request to %s failed with status code %d. Retrying in 3 seconds.", httpMethod, relativeUrl, resp.StatusCode)
		time.Sleep(3 * time.Second)
		return c.Request(httpMethod, relativeUrl, payload, attempt+1)
	}

	return resp, nil
}

func (client *Client) Get(url string) (*http.Response, error) {
	return client.Request("GET", url, []byte(""), 0)
}

func (client *Client) Put(url string, payload []byte) (*http.Response, error) {
	return client.Request("PUT", url, payload, 0)
}

func (client *Client) Post(url string, payload []byte) (*http.Response, error) {
	return client.Request("POST", url, payload, 0)
}

func (client *Client) Delete(url string) (*http.Response, error) {
	return client.Request("DELETE", url, []byte(""), 0)
}
