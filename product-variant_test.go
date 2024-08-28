package bigcommerce

import (
	"testing"
)

func TestGetProductVariants(t *testing.T) {
	fs, _ := getClient()

	_, _, err := fs.V3.GetProductVariants(193, ProductVariantQueryParams{})
	if err != nil {
		t.Error(err)
		return
	}
}
