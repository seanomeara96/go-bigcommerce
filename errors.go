package bigcommerce

import (
	"fmt"
	"net/http"
)

type BigCommerceError struct {
	StatusCode int
	Message    string
	RawBody    []byte
}

func (e *BigCommerceError) Error() string {
	return fmt.Sprintf("BigCommerce API error (status %d): %s", e.StatusCode, e.Message)
}

func NewBigCommerceError(resp *http.Response, body []byte) *BigCommerceError {
	return &BigCommerceError{
		StatusCode: resp.StatusCode,
		Message:    resp.Status,
		RawBody:    body,
	}
}
