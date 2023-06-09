package bigcommerce

import (
	"encoding/json"
	"fmt"
	"log"
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
	type ResponseObject struct {
		Data Category `json:"data"`
		Meta MetaData `json:"meta"`
	}
	var response ResponseObject

	var categoryURL string = client.BaseURL.JoinPath("/catalog/categories", fmt.Sprint(id)).String()

	resp, err := client.Get(categoryURL)
	if err != nil {
		return response.Data, err
	}
	defer resp.Body.Close()

	if err = expectStatusCode(200, resp); err != nil {
		return response.Data, err
	}

	if err = json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return response.Data, err
	}

	return response.Data, nil
}

func (client *Client) GetCategories(params CategoryQueryParams) ([]Category, MetaData, error) {
	type ResponseObject struct {
		Data []Category `json:"data"`
		Meta MetaData   `json:"meta"`
	}
	var response ResponseObject

	queryParams, err := paramString(params)

	if err != nil {
		return response.Data, response.Meta, err
	}

	var categoriesURL string = client.BaseURL.JoinPath("/catalog/categories").String() + queryParams

	resp, err := client.Get(categoriesURL)
	if err != nil {
		return response.Data, response.Meta, err
	}
	defer resp.Body.Close()

	if err = expectStatusCode(200, resp); err != nil {
		return response.Data, response.Meta, err
	}

	if err = json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return response.Data, response.Meta, err
	}

	return response.Data, response.Meta, nil
}

func (client *Client) GetAllCategories(limit int) ([]Category, error) {
	var categories []Category
	page := 1
	end := false

	for !end {
		c, m, err := client.GetCategories(CategoryQueryParams{Limit: limit, Page: page})
		if err != nil {
			return categories, err
		}

		details := m.Pagination
		log.Printf("Fetched %d categories. Current page: %d", details.Count, details.CurrentPage)

		for i := 0; i < len(c); i++ {
			categories = append(categories, c[i])
		}

		if len(c) < limit {
			end = true
			break
		}

		page++
	}

	return categories, nil
}
