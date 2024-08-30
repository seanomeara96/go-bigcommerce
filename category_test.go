package bigcommerce

import (
	"testing"
)

func TestGetCategory(t *testing.T) {
	fs, _ := getTestClient()

	categoryIdDoesNotExist := 11

	_, err := fs.V3.GetCategory(categoryIdDoesNotExist)

	if err == nil {
		t.Error("Expected Error")
	}

	categoryIdDoesExist := 27

	category, err := fs.V3.GetCategory(categoryIdDoesExist)

	if err != nil {
		t.Error(err)
		return
	}

	if category.ID != categoryIdDoesExist {
		t.Error("reponse-category id soes not match request category id")
	}

}

func TestGetCategories(t *testing.T) {

	fs, _ := getTestClient()

	categories, _, err := fs.V3.GetCategories(CategoryQueryParams{})

	if err != nil {
		t.Error(err)
	}

	if len(categories) < 1 {
		t.Error("no catgories")
	}
}

func TestGetAllCategories(t *testing.T) {

	fs, _ := getTestClient()
	categories, err := fs.V3.GetAllCategories(CategoryQueryParams{})
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

	fs, _ := getTestClient()

	products, err := fs.V3.GetAllProducts(ProductQueryParams{CategoriesIn: []int{24}})
	if err != nil {
		return
	}

	toRestore := products

	if len(products) != 2 {
		t.Error("Expected 2 products")
		return
	}

	err = fs.V3.EmptyCategory(24)
	if err != nil {
		t.Error(err)
		return
	}

	products, err = fs.V3.GetAllProducts(ProductQueryParams{CategoriesIn: []int{24}})
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

		_, err = fs.V3.UpdateProduct(product.ID, UpdateProductParams{
			Categories: newCategories,
		})
		if err != nil {
			t.Error("Could not update products")
			return
		}
	}
}
