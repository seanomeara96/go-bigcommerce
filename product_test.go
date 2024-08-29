package bigcommerce

import (
	"encoding/json"
	"testing"
)

func TestGetProductById(t *testing.T) {
	fs, _ := getTestClient()

	productId := 193

	product, err := fs.V3.GetProduct(productId, LimitedProductQueryParams{})

	if err != nil {
		t.Error(err)
	}

	if product.ID != productId {
		t.Error("Response-product id does not match request product id")
	}
}

func TestGetProductBySKU(t *testing.T) {
	fs, _ := getTestClient()

	productSKU := "14613"

	product, err := fs.V3.GetProductBySKU(productSKU)

	if err != nil {
		t.Error(err)
	}

	if product.ID != 211 {
		t.Errorf("expected id 211. received %d instead", product.ID)
	}
}

func TestGetAllProducts(t *testing.T) {
	fs, err := getTestClient()

	if err != nil {
		t.Error("error getting client")
	}

	products, err := fs.V3.GetAllProducts(ProductQueryParams{Include: []string{"images"}})
	if err != nil {
		t.Error(err)
		return
	}

	if len(products) < 1 {
		t.Error("no products")
		return
	}

	if len(products[1].Images) < 1 {
		t.Error("Expected images")
	}

}

func TestGetFullProductCatalog(t *testing.T) {
	fs, _ := getTestClient()

	products, err := fs.V3.GetAllProducts(ProductQueryParams{})
	if err != nil {
		t.Error(err)
		return
	}

	if len(products) != 87 {
		t.Error("did not fetch all products")
	}
}

func TestMarshalUpdateProductParams(t *testing.T) {
	paramsStruct := CreateUpdateProductParams{Name: "updated name"}
	paramBytes, err := json.Marshal(paramsStruct)
	if err != nil {
		t.Error(err)
		return
	}
	jsonString := string(paramBytes)
	expectedJsonString := `{"name":"updated name"}`
	if jsonString != expectedJsonString {
		t.Errorf("expected %s but received %s instead", expectedJsonString, jsonString)
		return
	}

	paramsStruct = CreateUpdateProductParams{Description: "updated description"}
	paramBytes, err = json.Marshal(paramsStruct)
	if err != nil {
		t.Error(err)
		return
	}
	jsonString = string(paramBytes)
	expectedJsonString = `{"description":"updated description"}`
	if jsonString != expectedJsonString {
		t.Errorf("expected %s but received %s instead", expectedJsonString, jsonString)
		return
	}

	paramsStruct = CreateUpdateProductParams{CustomURL: &CustomURL{
		URL:          "/new-url",
		IsCustomized: true,
	}}
	paramBytes, err = json.Marshal(paramsStruct)
	if err != nil {
		t.Error(err)
		return
	}
	jsonString = string(paramBytes)
	expectedJsonString = `{"custom_url":{"url":"/new-url","is_customized":true}}`
	if jsonString != expectedJsonString {
		t.Errorf("expected %s but received %s instead", expectedJsonString, jsonString)
		return
	}
}
