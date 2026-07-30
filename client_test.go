package bigcommerce

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"sync/atomic"
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

func TestGetRetriesTruncatedJSONResponse(t *testing.T) {
	var requests atomic.Int32
	client, closeServer := newTestClient(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if requests.Add(1) == 1 {
			_, _ = w.Write([]byte(`[{"id": 1}`))
			return
		}
		writeJSON(t, w, http.StatusOK, []map[string]int{{"id": 1}})
	}))
	defer closeServer()

	var response []map[string]int
	if err := client.V2.Get(client.V2.baseURL, &response); err != nil {
		t.Fatalf("expected truncated response to be retried: %v", err)
	}
	if requests.Load() != 2 {
		t.Fatalf("expected 2 requests, got %d", requests.Load())
	}
	if len(response) != 1 || response[0]["id"] != 1 {
		t.Fatalf("unexpected decoded response: %#v", response)
	}
}

func TestGetAcceptsEmptySuccessfulResponse(t *testing.T) {
	var requests atomic.Int32
	client, closeServer := newTestClient(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requests.Add(1)
		w.WriteHeader(http.StatusNoContent)
	}))
	defer closeServer()

	var response []map[string]int
	if err := client.V2.Get(client.V2.baseURL, &response); err != nil {
		t.Fatalf("expected empty successful response to represent an empty result: %v", err)
	}
	if requests.Load() != 1 {
		t.Fatalf("expected empty response not to be retried; got %d requests", requests.Load())
	}
	if len(response) != 0 {
		t.Fatalf("expected empty result, got %#v", response)
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
