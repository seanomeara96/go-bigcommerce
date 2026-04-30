package bigcommerce

import (
	"net/http"
	"testing"
)

func TestGetAllBanners(t *testing.T) {
	client, closeServer := newTestClient(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet || r.URL.Path != "/v2/banners" {
			t.Fatalf("unexpected request: %s %s", r.Method, r.URL.String())
		}
		writeJSON(t, w, http.StatusOK, []Banner{{ID: 1, Name: "Banner"}})
	}))
	defer closeServer()

	banners, _, err := client.V2.GetBanners(GetBannersParams{})
	if err != nil {
		t.Fatal(err)
	}
	if len(banners) != 1 {
		t.Fatalf("expected 1 banner, got %d", len(banners))
	}
}
