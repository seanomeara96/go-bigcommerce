package bigcommerce

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"sync"
	"time"
)

type RateLimitStatus struct {
	MsToReset         int64
	NextWindowTime    time.Time
	WindowSize        int64
	RequestsRemaining int
	RequestsQuota     int
}

type RateLimitConfig struct {
	MinRequestsRemaining int
	EnableWait           bool
}

type Client struct {
	baseURL         *url.URL
	authToken       string
	storeHash       string
	httpClient      *http.Client
	version         int
	rateLimitStatus *RateLimitStatus
	rateLimitConfig RateLimitConfig
	mu              sync.Mutex
}

func (c *Client) StoreHash() string {
	return c.storeHash
}

func (c *Client) BaseURL() *url.URL {
	return c.baseURL
}

func (c *Client) Version() int {
	return c.version
}

func NewClient(storeHash string, authToken string, version int, config *RateLimitConfig) *Client {
	url, err := url.Parse(fmt.Sprintf("https://api.bigcommerce.com/stores/%s/v%d", storeHash, version))
	var client Client
	if err != nil {
		log.Fatal(err)
	}

	client.baseURL = url
	client.authToken = authToken
	client.storeHash = storeHash
	client.httpClient = http.DefaultClient
	client.version = version

	if config == nil {
		client.rateLimitConfig = RateLimitConfig{
			MinRequestsRemaining: 2,
			EnableWait:           true,
		}
	} else {
		client.rateLimitConfig = *config
	}

	return &client
}

func (c *Client) configureRequest(httpMethod string, relativeUrl string, payload []byte) (*http.Request, error) {
	req, err := http.NewRequest(httpMethod, relativeUrl, bytes.NewBuffer(payload))
	if err != nil {
		return nil, err
	}

	req.Header.Set("x-auth-token", c.authToken)
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")

	return req, nil
}

func (c *Client) setRateLimitStatus(headers http.Header) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if msToReset, err := strconv.ParseInt(headers.Get("X-Rate-Limit-Time-Reset-Ms"), 10, 64); err == nil {
		now := time.Now()
		c.rateLimitStatus = &RateLimitStatus{
			MsToReset:         msToReset,
			NextWindowTime:    now.Add(time.Duration(msToReset) * time.Millisecond),
			WindowSize:        parseInt64(headers.Get("X-Rate-Limit-Time-Window-Ms")),
			RequestsRemaining: parseInt(headers.Get("X-Rate-Limit-Requests-Left")),
			RequestsQuota:     parseInt(headers.Get("X-Rate-Limit-Requests-Quota")),
		}
	}
}

func (c *Client) backoff() error {
	c.mu.Lock()
	defer c.mu.Unlock()

	if c.rateLimitStatus != nil {
		isAtRequestThreshold := c.rateLimitStatus.RequestsRemaining <= c.rateLimitConfig.MinRequestsRemaining
		if c.rateLimitConfig.EnableWait && isAtRequestThreshold {
			sleepDuration := time.Until(c.rateLimitStatus.NextWindowTime)
			if sleepDuration > 0 {
				time.Sleep(sleepDuration)
			}
		}
	}
	return nil
}

func (c *Client) request(httpMethod string, relativeUrl string, payload []byte, attempt int) (*http.Response, error) {
	if err := c.backoff(); err != nil {
		return nil, err
	}

	req, err := c.configureRequest(httpMethod, relativeUrl, payload)
	if err != nil {
		return nil, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		if attempt < 3 {
			log.Printf("Request failed. Retrying in 3 seconds. Error: %v", err)
			time.Sleep(3 * time.Second)
			return c.request(httpMethod, relativeUrl, payload, attempt+1)
		}
		return nil, err
	}

	c.setRateLimitStatus(resp.Header)

	if resp.StatusCode == 429 || (resp.StatusCode/100 == 5) {
		if attempt < 3 {
			waitTime := 3 * time.Second
			if resp.StatusCode == 429 {
				waitTime = time.Duration(c.rateLimitStatus.MsToReset) * time.Millisecond
			}
			log.Printf("%s request to %s failed with status code %d. Retrying in %v.", httpMethod, relativeUrl, resp.StatusCode, waitTime)
			time.Sleep(waitTime)
			return c.request(httpMethod, relativeUrl, payload, attempt+1)
		}
		if resp.StatusCode == 429 {
			return nil, fmt.Errorf("429 - Rate limit exceeded. Max retries reached")
		}
		return nil, fmt.Errorf("server error: %d. Max retries reached", resp.StatusCode)
	}

	if resp.StatusCode >= 400 && resp.StatusCode < 500 {
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Printf("Failed to read error response body: %v", err)
		} else {
			log.Printf("BigCommerce 4xx error response: %s", string(body))
		}
		// This line is recreating the response body after it has been read.
		resp.Body = io.NopCloser(bytes.NewBuffer(body))
	}

	return resp, nil
}

func (client *Client) requestAndDecode(httpMethod string, relativeUrl string, payload []byte, dest any) error {
	res, err := client.request(httpMethod, relativeUrl, payload, 0)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if dest != nil {
		if err := json.NewDecoder(res.Body).Decode(dest); err != nil {
			return err
		}
	}
	return nil
}

func (client *Client) marshalJSONandRequestAndDecode(httpMethod string, relativeUrl string, params any, dest any) error {
	var payload []byte
	if params != nil {
		p, err := json.Marshal(params)
		if err != nil {
			return err
		}
		payload = p
	}
	return client.requestAndDecode(httpMethod, relativeUrl, payload, dest)
}

func (client *Client) Get(url *url.URL, dest any) error {
	return client.marshalJSONandRequestAndDecode("GET", url.String(), nil, dest)

}

func (client *Client) Put(url *url.URL, params any, dest any) error {
	return client.marshalJSONandRequestAndDecode("PUT", url.String(), params, dest)
}

func (client *Client) Post(url *url.URL, params any, dest any) error {
	return client.marshalJSONandRequestAndDecode("POST", url.String(), params, dest)
}

func (client *Client) Delete(url *url.URL, dest any) error {
	return client.marshalJSONandRequestAndDecode("DELETE", url.String(), nil, dest)
}

// Helper functions
func parseInt(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}

func parseInt64(s string) int64 {
	i, _ := strconv.ParseInt(s, 10, 64)
	return i
}
