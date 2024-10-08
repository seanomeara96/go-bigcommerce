package bigcommerce

import (
	"errors"
	"fmt"
	"strconv"
)

type ProductVariantOption struct {
	ID           int             `json:"id"`
	ProductID    int             `json:"product_id"`
	DisplayName  string          `json:"display_name"`
	Type         string          `json:"type"`
	Config       []*OptionConfig `json:"config"`
	SortOrder    int             `json:"sort_order"`
	OptionValues []*OptionValue  `json:"option_values"`
	Name         string          `json:"name"`
}

type OptionConfig struct {
	DefaultValue                string   `json:"default_value"`
	CheckedByDefault            bool     `json:"checked_by_default"`
	CheckboxLabel               string   `json:"checkbox_label"`
	DateLimited                 bool     `json:"date_limited"`
	DateLimitMode               string   `json:"date_limit_mode"`
	DateEarliestValue           string   `json:"date_earliest_value"`
	DateLatestValue             string   `json:"date_latest_value"`
	FileTypesMode               string   `json:"file_types_mode"`
	FileTypesSupported          []string `json:"file_types_supported"`
	FileTypesOther              []string `json:"file_types_other"`
	FileMaxSize                 int      `json:"file_max_size"`
	TextCharactersLimited       bool     `json:"text_characters_limited"`
	TextMinLength               int      `json:"text_min_length"`
	TextMaxLength               int      `json:"text_max_length"`
	TextLinesLimited            bool     `json:"text_lines_limited"`
	TextMaxLines                int      `json:"text_max_lines"`
	NumberLimited               bool     `json:"number_limited"`
	NumberLimitMode             string   `json:"number_limit_mode"`
	NumberLowestValue           int      `json:"number_lowest_value"`
	NumberHighestValue          int      `json:"number_highest_value"`
	NumberIntegersOnly          bool     `json:"number_integers_only"`
	ProductListAdjustsInventory bool     `json:"product_list_adjusts_inventory"`
	ProductListAdjustsPricing   bool     `json:"product_list_adjusts_pricing"`
	ProductListShippingCalc     string   `json:"product_list_shipping_calc"`
}

type OptionValue struct {
	IsDefault bool        `json:"is_default"`
	Label     string      `json:"label"`
	SortOrder int         `json:"sort_order"`
	ValueData interface{} `json:"value_data"`
	ID        int         `json:"id"`
}

type CreateUpdateProductVariantOptions struct {
	ProductID    int       `json:"product_id"`
	DisplayName  string    `json:"display_name"`
	Type         string    `json:"type"`
	Config       *Config   `json:"config"`
	OptionValues []*Option `json:"option_values"`
}

type Config struct {
	DefaultValue            string   `json:"default_value,omitempty"`
	CheckedByDefault        bool     `json:"checked_by_default,omitempty"`
	CheckboxLabel           string   `json:"checkbox_label,omitempty"`
	DateLimited             bool     `json:"date_limited,omitempty"`
	DateLimitMode           string   `json:"date_limit_mode,omitempty"`
	DateEarliestValue       string   `json:"date_earliest_value,omitempty"`
	DateLatestValue         string   `json:"date_latest_value,omitempty"`
	FileTypesMode           string   `json:"file_types_mode,omitempty"`
	FileTypesSupported      []string `json:"file_types_supported,omitempty"`
	FileTypesOther          []string `json:"file_types_other,omitempty"`
	FileMaxSize             int      `json:"file_max_size,omitempty"`
	TextCharactersLimited   bool     `json:"text_characters_limited,omitempty"`
	TextMinLength           int      `json:"text_min_length,omitempty"`
	TextMaxLength           int      `json:"text_max_length,omitempty"`
	TextLinesLimited        bool     `json:"text_lines_limited,omitempty"`
	TextMaxLines            int      `json:"text_max_lines,omitempty"`
	NumberLimited           bool     `json:"number_limited,omitempty"`
	NumberLimitMode         string   `json:"number_limit_mode,omitempty"`
	NumberLowestValue       int      `json:"number_lowest_value,omitempty"`
	NumberHighestValue      int      `json:"number_highest_value,omitempty"`
	NumberIntegersOnly      bool     `json:"number_integers_only,omitempty"`
	ProductListAdjustsInv   bool     `json:"product_list_adjusts_inventory,omitempty"`
	ProductListAdjustsPrc   bool     `json:"product_list_adjusts_pricing,omitempty"`
	ProductListShippingCalc string   `json:"product_list_shipping_calc,omitempty"`
	SortOrder               int      `json:"sort_order,omitempty"`
}

type Option struct {
	IsDefault bool        `json:"is_default,omitempty"`
	Label     string      `json:"label"`
	SortOrder int         `json:"sort_order"`
	ValueData OptionValue `json:"value_data"`
}

func ValidateCreateUpdateProductVariantOptions(options CreateUpdateProductVariantOptions) error {
	if options.ProductID <= 0 {
		return errors.New("product_id must be a positive integer")
	}

	if len(options.DisplayName) < 1 || len(options.DisplayName) > 255 {
		return errors.New("display_name must be between 1 and 255 characters")
	}

	validTypes := [6]string{"radio_buttons", "rectangles", "dropdown", "product_list", "product_list_with_images", "swatch"}
	if !containsValidOptionType(validTypes, options.Type) {
		return errors.New("type is not valid")
	}

	for i := 0; i < len(options.OptionValues); i++ {
		option := options.OptionValues[i]
		if len(option.Label) == 0 {
			return errors.New("label is required for option values")
		}

		if option.SortOrder < 0 {
			return errors.New("sort_order must be a non-negative integer")
		}
	}

	return nil
}

func containsValidOptionType(slice [6]string, value string) bool {
	for _, item := range slice {
		if item == value {
			return true
		}
	}
	return false
}

func (client *V3Client) GetProductVariantOptions(product_id int) ([]ProductVariantOption, error) {
	type ResponseObject struct {
		Data []ProductVariantOption `json:"data"`
		Meta MetaData               `json:"meta"`
	}
	var response ResponseObject
	path := client.constructURL("catalog", "products", strconv.Itoa(product_id), "options")

	err := client.Get(path, &response)
	if err != nil {
		return nil, fmt.Errorf("failed to get product variant options for product ID %d: %w", product_id, err)
	}

	return response.Data, nil
}
func (client *V3Client) CreateProductVariantOption(product_id int, params CreateUpdateProductVariantOptions) (ProductVariantOption, error) {
	type ResponseObject struct {
		Data ProductVariantOption `json:"data"`
		Meta MetaData             `json:"meta"`
	}
	var response ResponseObject

	err := ValidateCreateUpdateProductVariantOptions(params)
	if err != nil {
		return ProductVariantOption{}, fmt.Errorf("invalid params for CreateProductVariantOption (product ID: %d): %w", product_id, err)
	}

	path := client.constructURL("catalog", "products", strconv.Itoa(product_id), "options")

	if err := client.Post(path, params, &response); err != nil {
		return ProductVariantOption{}, fmt.Errorf("failed to create product variant option for product ID %d: %w", product_id, err)
	}

	return response.Data, nil
}
func (client *V3Client) GetProductVariantOption(product_id, option_id int) (ProductVariantOption, error) {
	type ResponseObject struct {
		Data ProductVariantOption `json:"data"`
		Meta MetaData             `json:"meta"`
	}
	var response ResponseObject
	path := client.constructURL("catalog", "products", strconv.Itoa(product_id), "options", strconv.Itoa(option_id))

	err := client.Get(path, &response)
	if err != nil {
		return ProductVariantOption{}, fmt.Errorf("failed to get product variant option ID %d for product ID %d: %w", option_id, product_id, err)
	}

	return response.Data, nil
}
func (client *V3Client) UpdateProductVariantOption(product_id, option_id int, params CreateUpdateProductVariantOptions) (ProductVariantOption, error) {
	type ResponseObject struct {
		Data ProductVariantOption `json:"data"`
		Meta MetaData             `json:"meta"`
	}
	var response ResponseObject

	err := ValidateCreateUpdateProductVariantOptions(params)
	if err != nil {
		return ProductVariantOption{}, fmt.Errorf("invalid params for UpdateProductVariantOption (product ID: %d, option ID: %d): %w", product_id, option_id, err)
	}

	path := client.constructURL("/catalog/products/", strconv.Itoa(product_id), "/options", strconv.Itoa(option_id))

	if err := client.Put(path, params, &response); err != nil {
		return ProductVariantOption{}, fmt.Errorf("failed to update product variant option ID %d for product ID %d: %w", option_id, product_id, err)
	}

	return response.Data, nil
}
func (client *V3Client) DeleteProductVariantOption(product_id, option_id int) error {
	path := client.constructURL("catalog", "products", strconv.Itoa(product_id), "options", strconv.Itoa(option_id))
	err := client.Delete(path, nil)
	if err != nil {
		return fmt.Errorf("failed to delete product variant option ID %d for product ID %d: %w", option_id, product_id, err)
	}

	return nil
}
