package bigcommerce

import (
	"os"
	"testing"

	"github.com/joho/godotenv"
)

func TestGetProductVriantOptionsById(t *testing.T) {
	var client *Client
	err := godotenv.Load()
	if err != nil {
		t.Error(err)
	}

	storeHash := os.Getenv("BF_STORE_HASH")
	xAuthToken := os.Getenv("BF_XAUTHTOKEN")

	client = NewClient(storeHash, xAuthToken, 3, nil)

	_, err = client.GetProductVariantOptions(6073)
	if err != nil {
		t.Error(err)
	}

}
