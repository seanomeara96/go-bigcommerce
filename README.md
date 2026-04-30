This library provides an unofficial Go client for the bigcommerce REST API

### Installation:
```sh
go get github.com/seanomeara96/go-bigcommerce
```


### bigcommerce example usage:

```go
package main

import (
	"fmt"
	"os"

	"github.com/seanomeara96/go-bigcommerce"
)

func main() {
	storeHash := os.Getenv("STORE_HASH")
	xAuthToken := os.Getenv("XAUTHTOKEN")

	store := bigcommerce.NewClient(storeHash, xAuthToken, nil, nil)

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

For update and query parameters where an explicit `false` value must be sent,
use pointer booleans:

```go
_, err := store.V3.UpdateProduct(productID, bigcommerce.UpdateProductParams{
	IsVisible: bigcommerce.BoolPtr(false),
})
```
