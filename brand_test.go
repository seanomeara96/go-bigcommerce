package bigcommerce

import (
	"net/http"
	"testing"
)

func TestGetBrand(t *testing.T) {
	client, closeServer := newTestClient(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet || r.URL.Path != "/v3/catalog/brands/49" {
			t.Fatalf("unexpected request: %s %s", r.Method, r.URL.String())
		}
		writeJSON(t, w, http.StatusOK, map[string]any{
			"data": Brand{ID: 49, Name: "Brand"},
		})
	}))
	defer closeServer()

	brand, err := client.V3.GetBrand(49)
	if err != nil {
		t.Fatal(err)
	}
	if brand.ID != 49 {
		t.Fatalf("expected brand ID 49, got %d", brand.ID)
	}
}

func TestGetBrands(t *testing.T) {
	client, closeServer := newTestClient(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet || r.URL.Path != "/v3/catalog/brands" {
			t.Fatalf("unexpected request: %s %s", r.Method, r.URL.String())
		}
		writeJSON(t, w, http.StatusOK, map[string]any{
			"data": []Brand{{ID: 1, Name: "One"}},
		})
	}))
	defer closeServer()

	brands, _, err := client.V3.GetBrands(BrandQueryParams{})
	if err != nil {
		t.Fatal(err)
	}
	if len(brands) != 1 {
		t.Fatalf("expected 1 brand, got %d", len(brands))
	}
}
