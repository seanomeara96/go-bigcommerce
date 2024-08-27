package bigcommerce

import (
	"fmt"
	"net/url"

	"github.com/google/go-querystring/query"
)

func urlWithQueryParams(u *url.URL, params interface{}) (*url.URL, error) {
	values, err := query.Values(params)
	if err != nil {
		return nil, fmt.Errorf("failed to convert struct to url.Values: %w", err)
	}
	u.RawQuery = values.Encode()
	return u, nil
}

func (c *Client) constructURL(pathComponents ...string) *url.URL {
	u := c.BaseURL().JoinPath(pathComponents...)
	return u
}
