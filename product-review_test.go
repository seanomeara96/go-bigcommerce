package bigcommerce

import (
	"encoding/json"
	"net/http"
	"reflect"
	"testing"
)

func TestProductReviewCRUD(t *testing.T) {
	var requests int
	client, closeServer := newTestClient(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requests++
		if r.URL.Path != "/v3/catalog/products/42/reviews" && r.URL.Path != "/v3/catalog/products/42/reviews/7" {
			t.Fatalf("unexpected path: %s", r.URL.Path)
		}

		switch requests {
		case 1:
			if r.Method != http.MethodGet || r.URL.Query().Get("status") != "0" || r.URL.Query().Get("limit") != "25" {
				t.Fatalf("unexpected list request: %s %s", r.Method, r.URL.RawQuery)
			}
			writeJSON(t, w, http.StatusOK, map[string]any{
				"data": []ProductReview{{ID: 7, ProductID: 42, Title: "Useful", Status: ProductReviewStatusPending}},
				"meta": MetaData{Pagination: Pagination{CurrentPage: 1, TotalPages: 1}},
			})
		case 2:
			if r.Method != http.MethodGet {
				t.Fatalf("expected GET, got %s", r.Method)
			}
			writeJSON(t, w, http.StatusOK, map[string]any{"data": ProductReview{ID: 7, ProductID: 42, Title: "Useful"}})
		case 3:
			if r.Method != http.MethodPost {
				t.Fatalf("expected POST, got %s", r.Method)
			}
			var body CreateProductReviewParams
			if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
				t.Fatal(err)
			}
			if body.Title != "New review" || body.DateReviewed != "2026-08-04T12:00:00Z" || body.Rating != 5 {
				t.Fatalf("unexpected create body: %#v", body)
			}
			writeJSON(t, w, http.StatusOK, map[string]any{"data": ProductReview{ID: 7, ProductID: 42, Title: body.Title}})
		case 4:
			if r.Method != http.MethodPut {
				t.Fatalf("expected PUT, got %s", r.Method)
			}
			var body map[string]any
			if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
				t.Fatal(err)
			}
			if !reflect.DeepEqual(body, map[string]any{"rating": float64(0), "text": ""}) {
				t.Fatalf("update did not preserve zero values: %#v", body)
			}
			writeJSON(t, w, http.StatusOK, map[string]any{"data": ProductReview{ID: 7, ProductID: 42, Rating: 0}})
		case 5:
			if r.Method != http.MethodDelete {
				t.Fatalf("expected DELETE, got %s", r.Method)
			}
			w.WriteHeader(http.StatusNoContent)
		default:
			t.Fatalf("unexpected extra request")
		}
	}))
	defer closeServer()

	pending := 0
	reviews, meta, err := client.V3.GetProductReviews(42, ProductReviewQueryParams{Status: &pending, Limit: 25})
	if err != nil || len(reviews) != 1 || reviews[0].ID != 7 || meta.Pagination.TotalPages != 1 {
		t.Fatalf("unexpected list result: %#v %#v %v", reviews, meta, err)
	}
	if _, err := client.V3.GetProductReview(42, 7); err != nil {
		t.Fatal(err)
	}
	if _, err := client.V3.CreateProductReview(42, CreateProductReviewParams{Title: "New review", DateReviewed: "2026-08-04T12:00:00Z", Rating: 5}); err != nil {
		t.Fatal(err)
	}
	zero, empty := 0, ""
	if _, err := client.V3.UpdateProductReview(42, 7, UpdateProductReviewParams{Rating: &zero, Text: &empty}); err != nil {
		t.Fatal(err)
	}
	if err := client.V3.DeleteProductReview(42, 7); err != nil {
		t.Fatal(err)
	}
}

func TestGetAllProductReviewsPaginates(t *testing.T) {
	client, closeServer := newTestClient(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		page := r.URL.Query().Get("page")
		if r.URL.Query().Get("limit") != "250" {
			t.Fatalf("expected maximum page size, got %q", r.URL.Query().Get("limit"))
		}
		if page == "1" {
			writeJSON(t, w, http.StatusOK, map[string]any{"data": []ProductReview{{ID: 1}}, "meta": MetaData{Pagination: Pagination{CurrentPage: 1, TotalPages: 2}}})
			return
		}
		writeJSON(t, w, http.StatusOK, map[string]any{"data": []ProductReview{{ID: 2}}, "meta": MetaData{Pagination: Pagination{CurrentPage: 2, TotalPages: 2}}})
	}))
	defer closeServer()

	reviews, err := client.V3.GetAllProductReviews(42, ProductReviewQueryParams{})
	if err != nil {
		t.Fatal(err)
	}
	if len(reviews) != 2 || reviews[0].ID != 1 || reviews[1].ID != 2 {
		t.Fatalf("unexpected reviews: %#v", reviews)
	}
}

func TestCreateProductReviewValidatesRequiredFields(t *testing.T) {
	client := NewClient("store", "token", nil, nil)
	if _, err := client.V3.CreateProductReview(42, CreateProductReviewParams{}); err == nil {
		t.Fatal("expected required-field validation error")
	}
}
