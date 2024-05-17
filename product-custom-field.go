package bigcommerce

import (
	"encoding/json"
	"fmt"
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
	queryString, err := paramString(params)
	if err != nil {
		return response.Data, err
	}

	getCustomFieldPath := client.BaseURL().JoinPath("/catalog/products", fmt.Sprint(productID), "/custom-fields").String() + queryString

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

	paramBytes, err := json.Marshal(params)
	if err != nil {
		return response.Data, err
	}

	createCustomFieldpath := client.BaseURL().JoinPath("/catalog/products", fmt.Sprint(productID), "/custom-fields").String()

	resp, err := client.Post(createCustomFieldpath, paramBytes)
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
	getCustomFieldPath := client.BaseURL().JoinPath("/catalog/products", fmt.Sprint(productID), "custom-fields", fmt.Sprint(customFieldID)).String()

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

	paramBytes, err := json.Marshal(params)
	if err != nil {
		return response.Data, err
	}

	updateCustomFieldPath := client.BaseURL().JoinPath("/catalog/products", fmt.Sprint(productID), "custom-fields", fmt.Sprint(customFieldID)).String()

	resp, err := client.Put(updateCustomFieldPath, paramBytes)
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
	deleteCustomFieldPath := client.BaseURL().JoinPath("/catalog/products", fmt.Sprint(productID), "custom-fields", fmt.Sprint(customFieldID)).String()
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
