package bigcommerce

import (
	"net/http"
	"testing"
)

func TestGetProductVariants(t *testing.T) {
	client, closeServer := newTestClient(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet || r.URL.Path != "/v3/catalog/products/193/variants" {
			t.Fatalf("unexpected request: %s %s", r.Method, r.URL.String())
		}
		writeJSON(t, w, http.StatusOK, map[string]any{
			"data": []ProductVariant{{ID: 1, ProductID: 193, SKU: "SKU-1"}},
		})
	}))
	defer closeServer()

	variants, _, err := client.V3.GetProductVariants(193, ProductVariantQueryParams{})
	if err != nil {
		t.Fatal(err)
	}
	if len(variants) != 1 {
		t.Fatalf("expected 1 variant, got %d", len(variants))
	}
}
