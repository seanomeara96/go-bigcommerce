package bigcommerce

import (
	"encoding/json"
	"net/http"
	"strings"
	"testing"
)

func TestGetProductById(t *testing.T) {
	client, closeServer := newTestClient(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet || r.URL.Path != "/v3/catalog/products/193" {
			t.Fatalf("unexpected request: %s %s", r.Method, r.URL.String())
		}
		writeJSON(t, w, http.StatusOK, map[string]any{
			"data": Product{ID: 193, Name: "Test product"},
		})
	}))
	defer closeServer()

	product, err := client.V3.GetProduct(193, LimitedProductQueryParams{})
	if err != nil {
		t.Fatal(err)
	}

	if product.ID != 193 {
		t.Fatalf("expected product ID 193, got %d", product.ID)
	}
}

func TestGetProductBySKU(t *testing.T) {
	client, closeServer := newTestClient(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case "/v3/catalog/variants":
			if r.URL.Query().Get("sku") != "14613" {
				t.Fatalf("expected SKU query, got %s", r.URL.RawQuery)
			}
			writeJSON(t, w, http.StatusOK, map[string]any{
				"data": []ProductVariant{{ID: 1, ProductID: 211, SKU: "14613"}},
			})
		case "/v3/catalog/products/211":
			writeJSON(t, w, http.StatusOK, map[string]any{
				"data": Product{ID: 211, Name: "SKU product"},
			})
		default:
			t.Fatalf("unexpected request path: %s", r.URL.Path)
		}
	}))
	defer closeServer()

	product, err := client.V3.GetProductBySKU("14613")
	if err != nil {
		t.Fatal(err)
	}

	if product.ID != 211 {
		t.Fatalf("expected product ID 211, got %d", product.ID)
	}
}

func TestGetAllProducts(t *testing.T) {
	client, closeServer := newTestClient(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet || r.URL.Path != "/v3/catalog/products" {
			t.Fatalf("unexpected request: %s %s", r.Method, r.URL.String())
		}
		if r.URL.Query().Get("include") != "images" {
			t.Fatalf("expected include=images, got %s", r.URL.RawQuery)
		}
		writeJSON(t, w, http.StatusOK, map[string]any{
			"data": []Product{
				{ID: 1, Name: "One"},
				{ID: 2, Name: "Two", Images: []ProductImage{{ID: 20, ProductID: 2}}},
			},
		})
	}))
	defer closeServer()

	products, err := client.V3.GetAllProducts(ProductQueryParams{Include: []string{"images"}})
	if err != nil {
		t.Fatal(err)
	}

	if len(products) != 2 {
		t.Fatalf("expected 2 products, got %d", len(products))
	}
	if len(products[1].Images) != 1 {
		t.Fatal("expected images on second product")
	}
}

func TestMarshalUpdateProductParams(t *testing.T) {
	cases := []struct {
		name     string
		params   UpdateProductParams
		expected string
	}{
		{
			name:     "name",
			params:   UpdateProductParams{Name: "updated name"},
			expected: `{"name":"updated name"}`,
		},
		{
			name:     "description",
			params:   UpdateProductParams{Description: "updated description"},
			expected: `{"description":"updated description"}`,
		},
		{
			name: "custom URL",
			params: UpdateProductParams{CustomURL: &CustomURL{
				URL:          "/new-url",
				IsCustomized: BoolPtr(true),
			}},
			expected: `{"custom_url":{"url":"/new-url","is_customized":true}}`,
		},
		{
			name:     "explicit false bool",
			params:   UpdateProductParams{IsVisible: BoolPtr(false)},
			expected: `{"is_visible":false}`,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			paramBytes, err := json.Marshal(tc.params)
			if err != nil {
				t.Fatal(err)
			}
			if string(paramBytes) != tc.expected {
				t.Fatalf("expected %s but received %s", tc.expected, string(paramBytes))
			}
		})
	}
}

func TestCreateProductValidationAndErrorReturn(t *testing.T) {
	client, closeServer := newTestClient(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost || r.URL.Path != "/v3/catalog/products" {
			t.Fatalf("unexpected request: %s %s", r.Method, r.URL.String())
		}
		http.Error(w, `{"title":"bad request"}`, http.StatusBadRequest)
	}))
	defer closeServer()

	_, err := client.V3.CreateProduct(CreateProductParams{
		Name:   "Valid product",
		Type:   "physical",
		Weight: 0,
		Price:  10,
	})
	if err == nil {
		t.Fatal("expected API error to be returned")
	}
	if !strings.Contains(err.Error(), "status 400") {
		t.Fatalf("expected status 400 error, got %v", err)
	}

	_, err = client.V3.CreateProduct(CreateProductParams{Name: "Bad product", Type: "service", Weight: 1, Price: 10})
	if err == nil {
		t.Fatal("expected invalid type error")
	}
}

func TestForEachProductDoesNotRequireLiveStore(t *testing.T) {
	client, closeServer := newTestClient(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet || r.URL.Path != "/v3/catalog/products" {
			t.Fatalf("unexpected request: %s %s", r.Method, r.URL.String())
		}
		writeJSON(t, w, http.StatusOK, map[string]any{
			"data": []Product{},
		})
	}))
	defer closeServer()

	if err := client.V3.ForEachProduct([]func(p *Product) bool{}); err != nil {
		t.Fatal(err)
	}
}
