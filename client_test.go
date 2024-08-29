package bigcommerce

import (
	"os"
	"testing"

	"github.com/joho/godotenv"
)

func getTestClient() (*Client, error) {
	var client *Client
	err := godotenv.Load()
	if err != nil {
		return client, err
	}

	storeHash := os.Getenv("FS_STORE_HASH")
	xAuthToken := os.Getenv("FS_XAUTHTOKEN")

	client = NewClient(storeHash, xAuthToken, nil, nil)

	return client, nil
}

func TestNewClient(t *testing.T) {
	NewClient("adsd", "adssda", nil, nil)
}
