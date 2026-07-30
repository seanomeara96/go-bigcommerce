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

type ResponseDecodeError struct {
	Body []byte
	Err  error
}

func (e *ResponseDecodeError) Error() string {
	return fmt.Sprintf("failed to decode response: %v", e.Err)
}

func (e *ResponseDecodeError) Unwrap() error {
	return e.Err
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
