package bigcommerce

import (
	"encoding/json"
	"net/http"
	"testing"
)

func TestGetCategory(t *testing.T) {
	client, closeServer := newTestClient(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case "/v3/catalog/categories/11":
			http.Error(w, `{"title":"not found"}`, http.StatusNotFound)
		case "/v3/catalog/categories/27":
			writeJSON(t, w, http.StatusOK, map[string]any{
				"data": Category{ID: 27, Name: "Category"},
			})
		default:
			t.Fatalf("unexpected request path: %s", r.URL.Path)
		}
	}))
	defer closeServer()

	if _, err := client.V3.GetCategory(11); err == nil {
		t.Fatal("expected error for missing category")
	}

	category, err := client.V3.GetCategory(27)
	if err != nil {
		t.Fatal(err)
	}
	if category.ID != 27 {
		t.Fatalf("expected category ID 27, got %d", category.ID)
	}
}

func TestGetCategoriesCanQueryExplicitFalse(t *testing.T) {
	client, closeServer := newTestClient(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/v3/catalog/categories" {
			t.Fatalf("unexpected path: %s", r.URL.Path)
		}
		if r.URL.Query().Get("is_visible") != "false" {
			t.Fatalf("expected is_visible=false, got %s", r.URL.RawQuery)
		}
		writeJSON(t, w, http.StatusOK, map[string]any{
			"data": []Category{{ID: 1, Name: "Hidden"}},
			"meta": MetaData{Pagination: Pagination{
				CurrentPage: 1,
				TotalPages:  1,
			}},
		})
	}))
	defer closeServer()

	categories, _, err := client.V3.GetCategories(CategoryQueryParams{IsVisible: BoolPtr(false)})
	if err != nil {
		t.Fatal(err)
	}
	if len(categories) != 1 {
		t.Fatalf("expected 1 category, got %d", len(categories))
	}
}

func TestGetAllCategories(t *testing.T) {
	client, closeServer := newTestClient(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		writeJSON(t, w, http.StatusOK, map[string]any{
			"data": []Category{{ID: 1}, {ID: 2}},
			"meta": MetaData{Pagination: Pagination{
				CurrentPage: 1,
				TotalPages:  1,
			}},
		})
	}))
	defer closeServer()

	categories, err := client.V3.GetAllCategories(CategoryQueryParams{})
	if err != nil {
		t.Fatal(err)
	}
	if len(categories) != 2 {
		t.Fatalf("expected 2 categories, got %d", len(categories))
	}
}

func TestEmptyCategory(t *testing.T) {
	client, closeServer := newTestClient(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch {
		case r.Method == http.MethodGet && r.URL.Path == "/v3/catalog/products":
			if r.URL.Query().Get("categories:in") != "24" {
				t.Fatalf("expected category filter, got %s", r.URL.RawQuery)
			}
			writeJSON(t, w, http.StatusOK, map[string]any{
				"data": []Product{{ID: 7, Categories: []int{24, 30}}},
			})
		case r.Method == http.MethodPut && r.URL.Path == "/v3/catalog/products/7":
			var params UpdateProductParams
			if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
				t.Fatal(err)
			}
			if len(params.Categories) != 1 || params.Categories[0] != 30 {
				t.Fatalf("expected category 24 to be removed, got %#v", params.Categories)
			}
			writeJSON(t, w, http.StatusOK, map[string]any{
				"data": Product{ID: 7, Categories: params.Categories},
			})
		default:
			t.Fatalf("unexpected request: %s %s", r.Method, r.URL.String())
		}
	}))
	defer closeServer()

	if err := client.V3.EmptyCategory(24); err != nil {
		t.Fatal(err)
	}
}
