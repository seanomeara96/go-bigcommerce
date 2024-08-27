package bigcommerce

import (
	"strconv"
)

type Brand struct {
	ID              int       `json:"id"`
	Name            string    `json:"name"`
	MetaKeywords    []string  `json:"meta_keywords"`
	MetaDescription string    `json:"meta_description"`
	ImageURL        string    `json:"image_url"`
	SearchKeywords  string    `json:"search_keywords"`
	CustomURL       CustomURL `json:"custom_url"`
}

type BrandQueryParams struct {
	ID            int    `url:"id,omitempty"`
	IDIn          []int  `url:"id:in,omitempty"`
	IDNotIn       []int  `url:"id:not_in,omitempty"`
	IDMin         []int  `url:"id:min,omitempty"`
	IDMax         []int  `url:"id:max,omitempty"`
	IDGreater     []int  `url:"id:greater,omitempty"`
	IDLess        []int  `url:"id:less,omitempty"`
	Name          string `url:"name,omitempty"`
	PageTitle     string `url:"page_title,omitempty"`
	Page          int    `url:"page,omitempty"`
	Limit         int    `url:"limit,omitempty"`
	IncludeFields string `url:"include_fields,omitempty"`
	ExcludeFields string `url:"exclude_fields,omitempty"`
}

func (client *Client) GetBrand(id int) (Brand, error) {
	type ResponseObject struct {
		Data Brand    `json:"data"`
		Meta MetaData `json:"meta"`
	}

	var response ResponseObject

	brandURL := client.constructURL("/catalog/brands", strconv.Itoa(id))

	if err := client.Get(brandURL, &response); err != nil {
		return response.Data, nil
	}

	return response.Data, nil
}

func (client *Client) GetBrands(params BrandQueryParams) ([]Brand, MetaData, error) {
	type ResponseObject struct {
		Data []Brand  `json:"data"`
		Meta MetaData `json:"meta"`
	}
	var response ResponseObject

	brandsURL, err := urlWithQueryParams(client.constructURL("/catalog/brands"), params)
	if err != nil {
		return response.Data, response.Meta, err
	}

	if err := client.Get(brandsURL, &response); err != nil {
		return response.Data, response.Meta, err
	}

	return response.Data, response.Meta, nil
}

func (client *Client) GetAllBrands(params BrandQueryParams) ([]Brand, error) {
	var brands []Brand
	params.Page = 1
	params.Limit = 250

	for {
		b, _, err := client.GetBrands(params)
		if err != nil {
			return brands, err
		}
		for i := 0; i < len(b); i++ {
			brands = append(brands, b[i])
		}

		if len(b) < params.Limit {
			break
		}

		params.Page++
	}
	return brands, nil
}
