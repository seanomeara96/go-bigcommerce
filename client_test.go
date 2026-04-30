package bigcommerce

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func newTestClient(t *testing.T, handler http.Handler) (*Client, func()) {
	t.Helper()

	server := httptest.NewServer(handler)
	client := NewClient("test-store", "test-token", &RateLimitConfig{EnableWait: false}, nil)

	v2URL, err := url.Parse(server.URL + "/v2")
	if err != nil {
		server.Close()
		t.Fatalf("failed to parse v2 test URL: %v", err)
	}

	v3URL, err := url.Parse(server.URL + "/v3")
	if err != nil {
		server.Close()
		t.Fatalf("failed to parse v3 test URL: %v", err)
	}

	client.V2.baseURL = v2URL
	client.V3.baseURL = v3URL

	return client, server.Close
}

func TestNewClient(t *testing.T) {
	client := NewClient("store-hash", "auth-token", nil, nil)
	if client.V2 == nil || client.V3 == nil {
		t.Fatal("expected both V2 and V3 clients")
	}
}

func writeJSON(t *testing.T, w http.ResponseWriter, status int, v any) {
	t.Helper()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(v); err != nil {
		t.Fatalf("failed to write JSON response: %v", err)
	}
}
