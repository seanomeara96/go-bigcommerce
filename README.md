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

	store := bigcommerce.NewClient(storeHash, xAuthToken, 3)

	products, err := store.GetAllProducts(bigcommerce.ProductQueryParams{})
	if err != nil {
		panic(err)
	}
	fmt.Println(products)
}

```
