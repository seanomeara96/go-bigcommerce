package bigcommerce

import (
	"strconv"
)

type Category struct {
	ID                 int       `json:"id"`
	ParentID           int       `json:"parent_id"`
	Name               string    `json:"name"`
	Description        string    `json:"description"`
	Views              int       `json:"views"`
	SortOrder          int       `json:"sort_order"`
	PageTitle          string    `json:"page_title"`
	SearchKeywords     string    `json:"search_keywords"`
	MetaKeywords       []string  `json:"meta_keywords"`
	MetaDescription    string    `json:"meta_description"`
	LayoutFile         string    `json:"layout_file"`
	IsVisible          bool      `json:"is_visible"`
	DefaultProductSort string    `json:"default_product_sort"`
	ImageURL           string    `json:"image_url"`
	CustomURL          CustomURL `json:"custom_url"`
}

type CategoryQueryParams struct {
	ID              int      `url:"id,omitempty"`
	IDIn            []int    `url:"id:in,omitempty"`
	IDNotIn         []int    `url:"id:not_in,omitempty"`
	IDMin           []int    `url:"id:min,omitempty"`
	IDMax           []int    `url:"id:max,omitempty"`
	IDGreater       []int    `url:"id:greater,omitempty"`
	IDLess          []int    `url:"id:less,omitempty"`
	Name            string   `url:"name,omitempty"`
	NameLike        []string `url:"name:like,omitempty"`
	ParentID        int      `url:"parent_id,omitempty"`
	ParentIDIn      []int    `url:"parent_id:in,omitempty"`
	ParentIDMin     []int    `url:"parent_id:min,omitempty"`
	ParentIDMax     []int    `url:"parent_id:max,omitempty"`
	ParentIDGreater []int    `url:"parent_id:greater,omitempty"`
	ParentIDLess    []int    `url:"parent_id:less,omitempty"`
	PageTitle       string   `url:"page_title,omitempty"`
	PageTitleLike   []string `url:"page_title:like,omitempty"`
	Keyword         string   `url:"keyword,omitempty"`
	IsVisible       bool     `url:"is_visible,omitempty"`
	Page            int      `url:"page,omitempty"`
	Limit           int      `url:"limit,omitempty"`
	IncludeFields   string   `url:"include_fields,omitempty"`
	ExcludeFields   string   `url:"exclude_fields,omitempty"`
}

func (client *Client) GetCategory(id int) (Category, error) {
	var response struct {
		Data Category `json:"data"`
	}

	categoryURL := client.constructURL("/catalog/categories", strconv.Itoa(id))
	if err := client.Get(categoryURL, &response); err != nil {
		return Category{}, err
	}

	return response.Data, nil
}

func (client *Client) GetCategories(params CategoryQueryParams) ([]Category, MetaData, error) {
	var response struct {
		Data []Category `json:"data"`
		Meta MetaData   `json:"meta"`
	}

	categoriesURL, err := urlWithQueryParams(client.constructURL("/catalog/categories"), params)
	if err != nil {
		return nil, response.Meta, err
	}

	if err := client.Get(categoriesURL, &response); err != nil {
		return nil, response.Meta, err
	}

	return response.Data, response.Meta, nil
}

func (client *Client) GetAllCategories(params CategoryQueryParams) ([]Category, error) {
	var allCategories []Category

	if params.Page < 1 {
		params.Page = 1
	}
	if params.Limit < 1 {
		params.Limit = 250
	}

	for {
		categories, meta, err := client.GetCategories(params)
		if err != nil {
			return nil, err
		}

		allCategories = append(allCategories, categories...)

		if meta.Pagination.CurrentPage >= meta.Pagination.TotalPages {
			break
		}

		params.Page++
	}

	return allCategories, nil
}

func (client *Client) EmptyCategory(id int) error {
	products, _, err := client.GetProducts(ProductQueryParams{CategoriesIn: []int{id}})
	if err != nil {
		return err
	}

	for _, product := range products {
		categories := removeCategory(product.Categories, id)
		_, err = client.UpdateProduct(product.ID, CreateUpdateProductParams{Categories: categories})
		if err != nil {
			return err
		}
	}

	return nil
}

// Helper function to remove a category from a slice of categories
func removeCategory(categories []int, id int) []int {
	result := make([]int, 0, len(categories))
	for _, categoryID := range categories {
		if categoryID != id {
			result = append(result, categoryID)
		}
	}
	return result
}
