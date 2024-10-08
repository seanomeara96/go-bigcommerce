package bigcommerce

import (
	"fmt"
	"strconv"
)

func (c *V3Client) GetAllVariants(queryParams AllProductVariantsQueryParams) ([]ProductVariant, error) {
	page := 1
	if queryParams.Limit == 0 {
		queryParams.Limit = 250
	}
	all := []ProductVariant{}
	for {
		queryParams.Page = page
		res, _, err := c.GetVariants(queryParams)
		if err != nil {
			return []ProductVariant{}, fmt.Errorf("GetAllVariants: failed to get variants for page %d: %w", page, err)
		}
		if len(res) < 1 {
			return all, nil
		}

		all = append(all, res...)

		page++
	}
}

func (c *V3Client) GetVariants(queryParams AllProductVariantsQueryParams) ([]ProductVariant, MetaData, error) {
	type ResponseObject struct {
		Data []ProductVariant `json:"data"`
		Meta MetaData         `json:"meta"`
	}
	var response ResponseObject

	path, err := urlWithQueryParams(c.constructURL("/catalog/variants"), queryParams)
	if err != nil {
		return response.Data, response.Meta, fmt.Errorf("GetVariants: failed to construct URL with query params: %w", err)
	}

	if err := c.Get(path, &response); err != nil {
		return response.Data, response.Meta, fmt.Errorf("GetVariants: failed to make GET request: %w", err)
	}

	return response.Data, response.Meta, nil
}

func (client *V3Client) GetProductVariants(productID int, params ProductVariantQueryParams) ([]ProductVariant, MetaData, error) {
	type ResponseObject struct {
		Data []ProductVariant `json:"data"`
		Meta MetaData         `json:"meta"`
	}
	var response ResponseObject

	getProductVariantsURL, err := urlWithQueryParams(client.constructURL("/catalog/products", strconv.Itoa(productID), "/variants"), params)
	if err != nil {
		return response.Data, response.Meta, fmt.Errorf("GetProductVariants: failed to construct URL with query params for product ID %d: %w", productID, err)
	}

	if err := client.Get(getProductVariantsURL, &response); err != nil {
		return response.Data, response.Meta, fmt.Errorf("GetProductVariants: failed to make GET request for product ID %d: %w", productID, err)
	}

	return response.Data, response.Meta, nil
}

func (client *V3Client) CreateProductVariant(productID int, params ProductVariantCreateParams) (ProductVariant, error) {
	type ResponseObject struct {
		Data ProductVariant `json:"data"`
		Meta MetaData       `json:"meta"`
	}
	var response ResponseObject

	createProductVariantPath := client.constructURL("/catalog/products", fmt.Sprint(productID), "variants")

	if err := client.Post(createProductVariantPath, params, &response); err != nil {
		return response.Data, fmt.Errorf("CreateProductVariant: failed to create variant for product ID %d: %w", productID, err)
	}

	return response.Data, nil
}

type ProductVariantCreateParams struct {
	CostPrice                 float64          `json:"cost_price"`
	Price                     float64          `json:"price"`
	SalePrice                 float64          `json:"sale_price"`
	RetailPrice               float64          `json:"retail_price"`
	Weight                    float64          `json:"weight"`
	Width                     float64          `json:"width"`
	Height                    float64          `json:"height"`
	Depth                     float64          `json:"depth"`
	IsFreeShipping            bool             `json:"is_free_shipping"`
	FixedCostShippingPrice    float64          `json:"fixed_cost_shipping_price"`
	PurchasingDisabled        bool             `json:"purchasing_disabled"`
	PurchasingDisabledMessage string           `json:"purchasing_disabled_message"`
	UPC                       string           `json:"upc"`
	InventoryLevel            int              `json:"inventory_level"`
	InventoryWarningLevel     int              `json:"inventory_warning_level"`
	BinPickingNumber          string           `json:"bin_picking_number"`
	ImageURL                  string           `json:"image_url"`
	GTIN                      string           `json:"gtin"`
	MPN                       string           `json:"mpn"`
	ProductID                 int              `json:"product_id"`
	SKU                       string           `json:"sku"`
	OptionValues              *[]VariantOption `json:"option_values"`
}
type ProductVariant struct {
	ID                        int             `json:"id"`
	ProductID                 int             `json:"product_id"`
	SKU                       string          `json:"sku"`
	SKUID                     int             `json:"sku_id"`
	Price                     float64         `json:"price"`
	CalculatedPrice           float64         `json:"calculated_price"`
	SalePrice                 float64         `json:"sale_price"`
	RetailPrice               float64         `json:"retail_price"`
	MapPrice                  interface{}     `json:"map_price"`
	Weight                    float64         `json:"weight"`
	CalculatedWeight          float64         `json:"calculated_weight"`
	Width                     float64         `json:"width"`
	Height                    float64         `json:"height"`
	Depth                     float64         `json:"depth"`
	IsFreeShipping            bool            `json:"is_free_shipping"`
	FixedCostShippingPrice    float64         `json:"fixed_cost_shipping_price"`
	PurchasingDisabled        bool            `json:"purchasing_disabled"`
	PurchasingDisabledMessage string          `json:"purchasing_disabled_message"`
	ImageURL                  string          `json:"image_url"`
	CostPrice                 float64         `json:"cost_price"`
	UPC                       string          `json:"upc"`
	MPN                       string          `json:"mpn"`
	GTIN                      string          `json:"gtin"`
	InventoryLevel            int             `json:"inventory_level"`
	InventoryWarningLevel     int             `json:"inventory_warning_level"`
	BinPickingNumber          string          `json:"bin_picking_number"`
	OptionValues              []VariantOption `json:"option_values"`
}

type VariantOption struct {
	ID                int    `json:"id"`
	Label             string `json:"label"`
	OptionID          int    `json:"option_id"`
	OptionDisplayName string `json:"option_display_name"`
}

type ProductVariantQueryParams struct {
	Page          int    `url:"page,omitempty"`
	Limit         int    `url:"limit,omitempty"`
	IncludeFields string `url:"include_fields,omitempty"`
	ExcludeFields string `url:"exclude_fields,omitempty"`
}

type AllProductVariantsQueryParams struct {
	ID            int    `url:"id,omitempty"`
	SKU           string `url:"sku,omitempty"`
	UPC           string `url:"upc,omitempty"`
	Page          int    `url:"page,omitempty"`
	Limit         int    `url:"limit,omitempty"`
	IncludeFields string `url:"include_fields,omitempty"`
	ExcludeFields string `url:"exclude_fields,omitempty"`
	ProductID     string `url:"product_id,omitempty"`
}
