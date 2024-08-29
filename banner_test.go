package bigcommerce

import "testing"

func TestGetAllBanners(t *testing.T) {
	client, _ := getTestClient()

	banners, _, err := client.V2.GetBanners(GetBannersParams{})
	if err != nil {
		t.Error(err)
		return
	}

	if len(banners) < 1 {
		t.Error("Expected banners")
		return
	}
}
