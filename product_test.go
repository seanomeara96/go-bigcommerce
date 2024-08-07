package bigcommerce

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"

	"github.com/joho/godotenv"
)

func getClient() (Client, error) {
	var client Client
	err := godotenv.Load()
	if err != nil {
		return client, err
	}

	storeHash := os.Getenv("FS_STORE_HASH")
	xAuthToken := os.Getenv("FS_XAUTHTOKEN")

	client = NewClient(storeHash, xAuthToken, 3)

	return client, nil
}

func TestGetProductById(t *testing.T) {
	fs, _ := getClient()

	productId := 193

	product, err := fs.GetProduct(productId)

	if err != nil {
		t.Error(err)
	}

	if product.ID != productId {
		t.Error("Response-product id does not match repquest product id")
	}
}

func TestGetProductBySKU(t *testing.T) {
	fs, _ := getClient()

	productSKU := "7600"

	product, err := fs.GetProductBySKU(productSKU)

	if err != nil {
		t.Error(err)
	}
	fmt.Println(product.Name)
	t.Error()
}

func TestGetAllProducts(t *testing.T) {
	fs, err := getClient()

	if err != nil {
		t.Error("error getting client")
	}

	products, err := fs.GetAllProducts(ProductQueryParams{})
	if err != nil {
		t.Error(err)
		return
	}

	if len(products) < 1 {
		t.Error("no products")
	}

}

func TestGetFullProductCatalog(t *testing.T) {
	fs, _ := getClient()

	products, err := fs.GetAllProducts(ProductQueryParams{})
	if err != nil {
		t.Error(err)
		return
	}

	if len(products) != 69 {
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
