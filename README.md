This library provides an unofficial Go client for the bigcommerce REST API

### Installation:
```
go get github.com/seanomeara96/go-bigcommerce
```


### bigcommerce example usage:

```go
package main

import (
	"fmt"

  	"github.com/joho/godotenv"
  	"github.com/seanomeara96/go-bigcommerce"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	storeHash := os.Getenv("STORE_HASH")
	xAuthToken := os.Getenv("XAUTHTOKEN")

	store := bigcommerce.NewClient(storeHash, xAuthToken)

	products, err := store.V3.GetAllProducts(bigcommerce.ProductQueryParams{})
	if err != nil {
		if bcErr, ok := err.(*bigcommerce.BigCommerceError); ok {
			fmt.Printf("BigCommerce API error: Status %d, Message: %s\n", bcErr.StatusCode, bcErr.Message)
			fmt.Printf("Raw response body: %s\n", string(bcErr.RawBody))
		} else {
			fmt.Printf("Other error: %v\n", err)
		}
	}
	fmt.Println(products)
}

```
