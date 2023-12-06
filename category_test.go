package bigcommerce

import (
	"os"
	"testing"

	"github.com/joho/godotenv"
)

func TestGetCategory(t *testing.T) {

	err := godotenv.Load()
	if err != nil {
		t.Error(err)
		return
	}

	storeHash := os.Getenv("FS_STORE_HASH")
	xAuthToken := os.Getenv("FS_XAUTHTOKEN")

	fs := NewClient(storeHash, xAuthToken, 3)

	categoryIdDoesNotExist := 11

	_, err = fs.GetCategory(categoryIdDoesNotExist)

	if err == nil {
		t.Error("Expected Error")
	}

	categoryIdDoesExist := 27

	category, err := fs.GetCategory(categoryIdDoesExist)

	if err != nil {
		t.Error(err)
		return
	}

	if category.ID != categoryIdDoesExist {
		t.Error("reponse-category id soes not match request category id")
	}

}

func TestGetCategories(t *testing.T) {
	err := godotenv.Load()
	if err != nil {
		t.Error(err)
		return
	}

	storeHash := os.Getenv("FS_STORE_HASH")
	xAuthToken := os.Getenv("FS_XAUTHTOKEN")

	fs := NewClient(storeHash, xAuthToken, 3)

	categories, _, err := fs.GetCategories(CategoryQueryParams{})

	if err != nil {
		t.Error(err)
	}

	if len(categories) < 1 {
		t.Error("no catgories")
	}
}

func TestGetAllCategories(t *testing.T) {

	fs, _ := getClient()
	categories, err := fs.GetAllCategories(CategoryQueryParams{})
	if err != nil {
		t.Error(err)
		return
	}

	if len(categories) != 10 {
		t.Error("Not enough catgeories")
		return
	}
}

func TestEmptyCategory(t *testing.T) {

	fs, _ := getClient()

	products, _, err := fs.GetAllProducts(ProductQueryParams{CategoriesIn: []int{24}})
	if err != nil {
		return
	}

	toRestore := products

	if len(products) != 2 {
		t.Error("Expected 2 products")
		return
	}

	err = fs.EmptyCategory(24)
	if err != nil {
		t.Error(err)
		return
	}

	products, _, err = fs.GetAllProducts(ProductQueryParams{CategoriesIn: []int{24}})
	if err != nil {
		return
	}

	if len(products) != 0 {
		t.Error("Expected 0 products")
		return
	}

	for _, product := range toRestore {
		newCategories := product.AddCategory(24)
		b := false
		for _, id := range newCategories {
			if id == 24 {
				b = true
			}
		}
		if !b {
			t.Error("add category did not work")
			return
		}

		_, err = fs.UpdateProduct(product.ID, CreateUpdateProductParams{
			Categories: newCategories,
		})
		if err != nil {
			t.Error("Could not update products")
			return
		}
	}
}
