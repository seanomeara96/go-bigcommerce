package bigcommerce

import (
	"strconv"
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

	getAllImagesPath := client.constructURL("/catalog/products", strconv.Itoa(productID), "images")

	err := client.Get(getAllImagesPath, &response)
	if err != nil {
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

	getProductImagePath := client.constructURL("/catalog/products", strconv.Itoa(productID), "images", strconv.Itoa(imageID))

	err := client.Get(getProductImagePath, &response)
	if err != nil {
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
	createProductImagePath := client.constructURL("catalog", "products", strconv.Itoa(productID), "images")

	err := client.Post(createProductImagePath, params, &response)
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
	updateProductImagePath := client.constructURL("catalog", "products", strconv.Itoa(productID), "images", strconv.Itoa(imageID))

	err := client.Put(updateProductImagePath, params, &response)
	if err != nil {
		return response.Data, err
	}

	return response.Data, nil

}
func (client *Client) DeleteProductImage(productID int, imageID int) (bool, error) {
	deleteProductImagePath := client.constructURL("catalog", "products", strconv.Itoa(productID), "images", strconv.Itoa(imageID))

	err := client.Delete(deleteProductImagePath, nil)
	if err != nil {
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
