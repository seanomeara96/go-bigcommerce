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

type Logger interface {
	Printf(format string, v ...interface{})
}

type BaseVersionClient struct {
	baseURL         *url.URL
	version         int
	authToken       string
	storeHash       string
	rateLimitConfig *RateLimitConfig
	rateLimitStatus *RateLimitStatus
	mu              sync.Mutex
	logger          Logger
}

func (c *BaseVersionClient) BaseURL() *url.URL {
	return c.baseURL
}

type V2Client struct {
	BaseVersionClient
}

type V3Client struct {
	BaseVersionClient
}

type Client struct {
	V3 *V3Client
	V2 *V2Client
}

func NewClient(storeHash string, authToken string, config *RateLimitConfig, logger Logger) *Client {
	v2URL, err := url.Parse(fmt.Sprintf("https://api.bigcommerce.com/stores/%s/v%d", storeHash, 2))
	if err != nil {
		log.Fatalf("Failed to parse BigCommerce API URL: %v", err)
	}

	v3URL, err := url.Parse(fmt.Sprintf("https://api.bigcommerce.com/stores/%s/v%d", storeHash, 3))
	if err != nil {
		log.Fatalf("Failed to parse BigCommerce API URL: %v", err)
	}

	var client Client

	if config == nil {
		config = &RateLimitConfig{
			MinRequestsRemaining: 2,
			EnableWait:           true,
		}
	}

	client.V2 = &V2Client{
		BaseVersionClient: BaseVersionClient{
			baseURL:         v2URL,
			version:         2,
			authToken:       authToken,
			storeHash:       storeHash,
			rateLimitConfig: config,
			logger:          logger,
		},
	}

	client.V3 = &V3Client{
		BaseVersionClient: BaseVersionClient{
			baseURL:         v3URL,
			version:         3,
			authToken:       authToken,
			storeHash:       storeHash,
			rateLimitConfig: config,
			logger:          logger,
		},
	}

	return &client
}

func configureRequest(authToken, httpMethod, relativeUrl string, payload []byte) (*http.Request, error) {
	req, err := http.NewRequest(httpMethod, relativeUrl, bytes.NewBuffer(payload))
	if err != nil {
		return nil, fmt.Errorf("failed to create new request: %w", err)
	}

	req.Header.Set("x-auth-token", authToken)
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")

	return req, nil
}

func (c *BaseVersionClient) setRateLimitStatus(headers http.Header) {
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

func (c *BaseVersionClient) backoff() error {
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

func (c *BaseVersionClient) request(httpMethod string, relativeUrl string, payload []byte, attempt int) (*http.Response, error) {
	if c.logger != nil {
		c.logger.Printf("Attempting %s request to %s (attempt %d)", httpMethod, relativeUrl, attempt+1)
	}

	if err := c.backoff(); err != nil {
		return nil, fmt.Errorf("backoff failed: %w", err)
	}

	req, err := configureRequest(c.authToken, httpMethod, relativeUrl, payload)
	if err != nil {
		return nil, fmt.Errorf("failed to configure request: %w", err)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		if c.logger != nil {
			c.logger.Printf("Request failed: %v", err)
		}
		if attempt < 3 {
			log.Printf("Request failed. Retrying in 3 seconds. Error: %v", err)
			time.Sleep(3 * time.Second)
			return c.request(httpMethod, relativeUrl, payload, attempt+1)
		}
		return nil, fmt.Errorf("request failed after 3 attempts: %w", err)
	}

	if c.logger != nil {
		c.logger.Printf("Response received: Status %d", resp.StatusCode)
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
			return nil, fmt.Errorf("429 - Rate limit exceeded. Max retries reached for %s request to %s", httpMethod, relativeUrl)
		}
		return nil, fmt.Errorf("server error: %d. Max retries reached for %s request to %s", resp.StatusCode, httpMethod, relativeUrl)
	}

	if resp.StatusCode >= 400 && resp.StatusCode < 500 {
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, fmt.Errorf("failed to read error response body: %v", err)
		}
		return nil, fmt.Errorf("bigCommerce 4xx error response: %s", string(body))
	}

	return resp, nil
}

func (client *BaseVersionClient) requestAndDecode(httpMethod string, relativeUrl string, payload []byte, dest any) error {
	res, err := client.request(httpMethod, relativeUrl, payload, 0)
	if err != nil {
		return fmt.Errorf("request failed: %w", err)
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return fmt.Errorf("failed to read response body: %w", err)
	}

	if client.logger != nil {
		client.logger.Printf("Response body: %s", string(body))
	}

	if dest != nil {
		if err := json.Unmarshal(body, dest); err != nil {
			if client.logger != nil {
				client.logger.Printf("Failed to decode response: %s", string(body))
			}
			return fmt.Errorf("failed to decode response: %w", err)
		}
	}
	return nil
}

func (client *BaseVersionClient) marshalJSONandRequestAndDecode(httpMethod string, relativeUrl string, params any, dest any) error {
	var payload []byte
	if params != nil {
		p, err := json.Marshal(params)
		if err != nil {
			return fmt.Errorf("failed to marshal params: %w", err)
		}
		payload = p
	}
	return client.requestAndDecode(httpMethod, relativeUrl, payload, dest)
}

func (client *BaseVersionClient) Get(url *url.URL, dest any) error {
	return client.marshalJSONandRequestAndDecode("GET", url.String(), nil, dest)
}

func (client *BaseVersionClient) Put(url *url.URL, params any, dest any) error {
	return client.marshalJSONandRequestAndDecode("PUT", url.String(), params, dest)
}

func (client *BaseVersionClient) Post(url *url.URL, params any, dest any) error {
	return client.marshalJSONandRequestAndDecode("POST", url.String(), params, dest)
}

func (client *BaseVersionClient) Delete(url *url.URL, dest any) error {
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
