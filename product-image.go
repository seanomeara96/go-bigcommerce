package bigcommerce

import (
	"encoding/json"
	"fmt"
)

type ProductImage struct {
	ImageFile    string `json:"image_file"`
	IsThumbnail  bool   `json:"is_thumbnail"`
	SortOrder    int    `json:"sort_order"`
	Description  string `json:"description"`
	ImageURL     string `json:"image_url"`
	ID           int    `json:"id"`
	ProductID    int    `json:"product_id"`
	URLZoom      string `json:"url_zoom"`
	URLStandard  string `json:"url_standard"`
	URLThumbnail string `json:"url_thumbnail"`
	URLTiny      string `json:"url_tiny"`
	DateModified string `json:"date_modified"`
}

func (client *Client) GetAllProductImages(productID int) ([]ProductImage, error) {
	type ResponseObject struct {
		Data []ProductImage `json:"data"`
		Meta MetaData       `json:"meta"`
	}
	var response ResponseObject

	getAllImagesPath := client.BaseURL().JoinPath("/catalog/products", fmt.Sprint(productID), "images").String()

	resp, err := client.Get(getAllImagesPath)
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
func (client *Client) GetProductImage(productID int, imageID int) (ProductImage, error) {
	type ResponseObject struct {
		Data ProductImage `json:"data"`
		Meta MetaData     `json:"meta"`
	}
	var response ResponseObject

	getProductImagePath := client.BaseURL().JoinPath("/catalog/products", fmt.Sprint(productID), "images", fmt.Sprint(imageID)).String()

	resp, err := client.Get(getProductImagePath)
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
func (client *Client) CreateProductImage(productID int, params CreateProductImageParams) (ProductImage, error) {
	type ResponseObject struct {
		Data ProductImage `json:"data"`
		Meta MetaData     `json:"meta"`
	}
	var response ResponseObject
	// POST /catalog/products/{product_id}/images
	createProductImagePath := client.BaseURL().JoinPath("/catalog/products", fmt.Sprint(productID), "images").String()

	paramBytes, err := json.Marshal(params)
	if err != nil {
		return response.Data, err
	}

	resp, err := client.Post(createProductImagePath, paramBytes)
	if err != nil {
		return response.Data, err
	}

	err = expectStatusCode(200, resp)
	if err != nil {
		return response.Data, err
	}

	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return response.Data, err
	}

	return response.Data, nil

}
func (client *Client) UpdateProductImage(productID int, imageID int, params UpdateProductImageParams) (ProductImage, error) {
	type ResponseObject struct {
		Data ProductImage `json:"data"`
		Meta MetaData     `json:"meta"`
	}
	var response ResponseObject
	// PUT /catalog/products/{product_id}/images/{image_id}
	updateProductImagePath := client.BaseURL().JoinPath("/catalog/products", fmt.Sprint(productID), "images", fmt.Sprint(imageID)).String()

	paramBytes, err := json.Marshal(params)
	if err != nil {
		return response.Data, err
	}

	resp, err := client.Put(updateProductImagePath, paramBytes)
	if err != nil {
		return response.Data, err
	}

	err = expectStatusCode(200, resp)
	if err != nil {
		err = expectStatusCode(201, resp)
		if err != nil {
			return response.Data, err
		}
	}

	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return response.Data, err
	}

	return response.Data, nil

}
func (client *Client) DeleteProductImage(productID int, imageID int) (bool, error) {
	deleteProductImagePath := client.BaseURL().JoinPath("/catalog/products", fmt.Sprint(productID), "images", fmt.Sprint(imageID)).String()

	resp, err := client.Delete(deleteProductImagePath)
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()

	if err = expectStatusCode(204, resp); err != nil {
		return false, err
	}

	return true, nil
}

type CreateProductImageParams struct {
	ProductID    int    `json:"product_id"`
	ImageFile    string `json:"image_file,omitempty"`
	URLZoom      string `json:"url_zoom,omitempty"`
	URLStandard  string `json:"url_standard,omitempty"`
	URLThumbnail string `json:"url_thumbnail,omitempty"`
	URLTiny      string `json:"url_tiny,omitempty"`
	DateModified string `json:"date_modified,omitempty"`
	IsThumbnail  bool   `json:"is_thumbnail,omitempty"`
	SortOrder    int    `json:"sort_order,omitempty"`
	Description  string `json:"description,omitempty"`
	ImageURL     string `json:"image_url,omitempty"`
}

type UpdateProductImageParams struct {
	ProductID    int    `json:"product_id,omitempty"`
	URLZoom      string `json:"url_zoom,omitempty"`
	URLStandard  string `json:"url_standard,omitempty"`
	URLThumbnail string `json:"url_thumbnail,omitempty"`
	URLTiny      string `json:"url_tiny,omitempty"`
	ImageFile    string `json:"image_file,omitempty"`
	IsThumbnail  bool   `json:"is_thumbnail,omitempty"`
	SortOrder    int    `json:"sort_order,omitempty"`
	Description  string `json:"description,omitempty"`
	ImageURL     string `json:"image_url,omitempty"`
}
