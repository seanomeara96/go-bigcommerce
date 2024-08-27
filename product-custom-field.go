package bigcommerce

import (
	"encoding/json"
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
	// /catalog/products/{product_id}/custom-fields
	getCustomFieldPath, err := urlWithQueryParams(client.constructURL("/catalog/products", strconv.Itoa(productID), "/custom-fields"), params)
	if err != nil {
		return response.Data, err
	}
	resp, err := client.Get(getCustomFieldPath)
	if err != nil {
		return response.Data, err
	}
	defer resp.Body.Close()

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
func (client *Client) CreateCustomField(productID int, params CreateCustomFieldParams) (ProductCustomField, error) {
	type ResponseObject struct {
		Data ProductCustomField `json:"data"`
		Meta MetaData           `json:"meta"`
	}
	var response ResponseObject

	if params.Name == "" || params.Value == "" {
		return response.Data, fmt.Errorf("check params, no empty values allowed name: %s, value: %s", params.Name, params.Value)
	}

	createCustomFieldpath := client.constructURL("catalog", "products", strconv.Itoa(productID), "custom-fields")

	resp, err := client.Post(createCustomFieldpath, params)
	if err != nil {
		return response.Data, err
	}
	defer resp.Body.Close()

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
func (client *Client) GetCustomField(productID int, customFieldID int) (ProductCustomField, error) {
	type ResponseObject struct {
		Data ProductCustomField `json:"data"`
		Meta MetaData           `json:"meta"`
	}
	var response ResponseObject
	// /catalog/products/{product_id}/custom-fields/{custom_field_id}
	getCustomFieldPath := client.constructURL("catalog", "products", strconv.Itoa(productID), "custom-fields", strconv.Itoa(customFieldID))

	resp, err := client.Get(getCustomFieldPath)
	if err != nil {
		return response.Data, err
	}
	defer resp.Body.Close()

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
func (client *Client) UpdateCustomField(productID int, customFieldID int, params UpdateCustomFieldParams) (ProductCustomField, error) {
	type ResponseObject struct {
		Data ProductCustomField `json:"data"`
		Meta MetaData           `json:"meta"`
	}
	var response ResponseObject

	updateCustomFieldPath := client.constructURL("/catalog/products", strconv.Itoa(productID), "custom-fields", strconv.Itoa(customFieldID))

	resp, err := client.Put(updateCustomFieldPath, params)
	if err != nil {
		return response.Data, err
	}
	defer resp.Body.Close()

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
func (client *Client) DeleteCustomField(productID int, customFieldID int) error {
	deleteCustomFieldPath := client.constructURL("/catalog/products", strconv.Itoa(productID), "custom-fields", strconv.Itoa(customFieldID))
	resp, err := client.Delete(deleteCustomFieldPath)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	err = expectStatusCode(204, resp)
	if err != nil {
		return err
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
