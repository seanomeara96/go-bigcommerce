package bigcommerce

import (
	"encoding/json"
	"net/http"
	"testing"
)

const productIdDoesExist = 79

func TestCreateCoupon(t *testing.T) {
	client, closeServer := newTestClient(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost || r.URL.Path != "/v2/coupons" {
			t.Fatalf("unexpected request: %s %s", r.Method, r.URL.String())
		}

		var params CreateCouponParams
		if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
			t.Fatal(err)
		}
		if params.Name != "Test Coupon" || !params.Enabled {
			t.Fatalf("unexpected coupon params: %#v", params)
		}

		writeJSON(t, w, http.StatusOK, Coupon{ID: 100, Name: params.Name, Code: params.Code})
	}))
	defer closeServer()

	coupon, err := client.V2.CreateCoupon(CreateCouponParams{
		Name:    "Test Coupon",
		Type:    "per_item_discount",
		Amount:  "1.00",
		Enabled: true,
		Code:    "TEST",
		AppliesTo: &AppliesTo{
			IDs:    []int{productIdDoesExist},
			Entity: "products",
		},
	})
	if err != nil {
		t.Fatal(err)
	}
	if coupon.ID != 100 {
		t.Fatalf("expected coupon ID 100, got %d", coupon.ID)
	}
}

func TestUpdateCouponCanSendExplicitFalse(t *testing.T) {
	client, closeServer := newTestClient(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPut || r.URL.Path != "/v2/coupons/100" {
			t.Fatalf("unexpected request: %s %s", r.Method, r.URL.String())
		}

		var body map[string]any
		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			t.Fatal(err)
		}
		if enabled, ok := body["enabled"].(bool); !ok || enabled {
			t.Fatalf("expected enabled=false in body, got %#v", body)
		}

		writeJSON(t, w, http.StatusOK, Coupon{ID: 100, Name: "Updated", Enabled: false})
	}))
	defer closeServer()

	coupon, err := client.V2.UpdateCoupon(100, UpdateCouponParams{
		Name:    "Updated",
		Enabled: BoolPtr(false),
	})
	if err != nil {
		t.Fatal(err)
	}
	if coupon.ID != 100 {
		t.Fatalf("expected coupon ID 100, got %d", coupon.ID)
	}
}

func TestGetCoupons(t *testing.T) {
	client, closeServer := newTestClient(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet || r.URL.Path != "/v2/coupons" {
			t.Fatalf("unexpected request: %s %s", r.Method, r.URL.String())
		}
		if r.URL.Query().Get("limit") != "10" {
			t.Fatalf("expected limit=10, got %s", r.URL.RawQuery)
		}
		writeJSON(t, w, http.StatusOK, []Coupon{{ID: 1}, {ID: 2}})
	}))
	defer closeServer()

	coupons, err := client.V2.GetCoupons(CouponQueryParams{Limit: 10})
	if err != nil {
		t.Fatal(err)
	}
	if len(coupons) != 2 {
		t.Fatalf("expected 2 coupons, got %d", len(coupons))
	}
}

func TestGetCoupon(t *testing.T) {
	client, closeServer := newTestClient(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet || r.URL.Path != "/v2/coupons/100" {
			t.Fatalf("unexpected request: %s %s", r.Method, r.URL.String())
		}
		writeJSON(t, w, http.StatusOK, Coupon{ID: 100, Name: "Coupon"})
	}))
	defer closeServer()

	coupon, err := client.V2.GetCoupon(100)
	if err != nil {
		t.Fatal(err)
	}
	if coupon.ID != 100 {
		t.Fatalf("expected coupon ID 100, got %d", coupon.ID)
	}
}

func TestDeleteCoupon(t *testing.T) {
	client, closeServer := newTestClient(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodDelete || r.URL.Path != "/v2/coupons/100" {
			t.Fatalf("unexpected request: %s %s", r.Method, r.URL.String())
		}
		w.WriteHeader(http.StatusNoContent)
	}))
	defer closeServer()

	if err := client.V2.DeleteCoupon(100); err != nil {
		t.Fatal(err)
	}
}
