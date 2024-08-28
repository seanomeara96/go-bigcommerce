package bigcommerce

import (
	"fmt"
	"strconv"
)

type ProductCustomField struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Value string `json:"value"`
}

func (client *Client) GetCustomFields(productID int, params ProductCustomFieldsRequestParams) ([]ProductCustomField, error) {
	type ResponseObject struct {
		Data []ProductCustomField `json:"data"`
		Meta MetaData             `json:"meta"`
	}
	var response ResponseObject

	getCustomFieldPath, err := urlWithQueryParams(client.constructURL("catalog", "products", strconv.Itoa(productID), "custom-fields"), params)
	if err != nil {
		return nil, fmt.Errorf("failed to construct URL for GetCustomFields (product ID: %d): %w", productID, err)
	}

	if err := client.Get(getCustomFieldPath, &response); err != nil {
		return nil, fmt.Errorf("failed to get custom fields for product ID %d: %w", productID, err)
	}

	return response.Data, nil
}

func (client *Client) CreateCustomField(productID int, params CreateCustomFieldParams) (ProductCustomField, error) {
	type ResponseObject struct {
		Data ProductCustomField `json:"data"`
		Meta MetaData           `json:"meta"`
	}
	var response ResponseObject

	if params.Name == "" || params.Value == "" {
		return response.Data, fmt.Errorf("invalid params for CreateCustomField: name and value cannot be empty (product ID: %d)", productID)
	}

	createCustomFieldpath := client.constructURL("catalog", "products", strconv.Itoa(productID), "custom-fields")

	err := client.Post(createCustomFieldpath, params, &response)
	if err != nil {
		return response.Data, fmt.Errorf("failed to create custom field for product ID %d: %w", productID, err)
	}

	return response.Data, nil
}

func (client *Client) GetCustomField(productID int, customFieldID int) (ProductCustomField, error) {
	type ResponseObject struct {
		Data ProductCustomField `json:"data"`
		Meta MetaData           `json:"meta"`
	}
	var response ResponseObject

	getCustomFieldPath := client.constructURL("catalog", "products", strconv.Itoa(productID), "custom-fields", strconv.Itoa(customFieldID))

	err := client.Get(getCustomFieldPath, &response)
	if err != nil {
		return response.Data, fmt.Errorf("failed to get custom field ID %d for product ID %d: %w", customFieldID, productID, err)
	}

	return response.Data, nil
}

func (client *Client) UpdateCustomField(productID int, customFieldID int, params UpdateCustomFieldParams) (ProductCustomField, error) {
	type ResponseObject struct {
		Data ProductCustomField `json:"data"`
		Meta MetaData           `json:"meta"`
	}
	var response ResponseObject

	updateCustomFieldPath := client.constructURL("/catalog/products", strconv.Itoa(productID), "custom-fields", strconv.Itoa(customFieldID))

	err := client.Put(updateCustomFieldPath, params, &response)
	if err != nil {
		return response.Data, fmt.Errorf("failed to update custom field ID %d for product ID %d: %w", customFieldID, productID, err)
	}

	return response.Data, nil
}

func (client *Client) DeleteCustomField(productID int, customFieldID int) error {
	deleteCustomFieldPath := client.constructURL("/catalog/products", strconv.Itoa(productID), "custom-fields", strconv.Itoa(customFieldID))
	err := client.Delete(deleteCustomFieldPath, nil)
	if err != nil {
		return fmt.Errorf("failed to delete custom field ID %d for product ID %d: %w", customFieldID, productID, err)
	}

	return nil
}

type ProductCustomFieldsRequestParams struct {
	IncludeFields string `url:"include_fields,omitempty"`
	ExcludeFields string `url:"exclude_fields,omitempty"`
	Page          int    `url:"page,omitempty"`
	Limit         int    `url:"limit,omitempty"`
}

type CreateCustomFieldParams struct {
	Name  string `json:"name" validate:"required,min=1,max=250"`
	Value string `json:"value" validate:"required,min=1,max=250"`
}

type UpdateCustomFieldParams struct {
	ID    int    `json:"id"`
	Name  string `json:"name" validate:"required,min=1,max=250"`
	Value string `json:"value" validate:"required,min=1,max=250"`
}
