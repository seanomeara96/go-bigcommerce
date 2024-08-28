package bigcommerce

import (
	"testing"
)

func TestGetBrand(t *testing.T) {
	fs, _ := getClient()

	brandId := 49

	brand, err := fs.V3.GetBrand(brandId)

	if err != nil {
		t.Error(err)
	}

	if brand.ID != brandId {
		t.Errorf("response brand-id does not match request brand id. Expected %d got %d", brandId, brand.ID)
	}
}

func TestGetBrands(t *testing.T) {
	fs, _ := getClient()

	brands, _, err := fs.V3.GetBrands(BrandQueryParams{})

	if err != nil {
		t.Error(err)
	}

	if len(brands) < 1 {
		t.Error("no brands")
	}

}
