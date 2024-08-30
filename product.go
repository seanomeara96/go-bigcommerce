package bigcommerce

import (
	"errors"
	"fmt"
	"strconv"
	"time"
)

type ResponseObject struct {
	Data Product  `json:"data"`
	Meta MetaData `json:"meta"`
}

type Product struct {
	ID                          int                      `json:"id"`
	Name                        string                   `json:"name"`
	Type                        string                   `json:"type"`
	SKU                         string                   `json:"sku"`
	Description                 string                   `json:"description"`
	Weight                      float64                  `json:"weight"`
	Width                       float64                  `json:"width"`
	Depth                       float64                  `json:"depth"`
	Height                      float64                  `json:"height"`
	Price                       float64                  `json:"price"`
	CostPrice                   float64                  `json:"cost_price"`
	RetailPrice                 float64                  `json:"retail_price"`
	SalePrice                   float64                  `json:"sale_price"`
	MapPrice                    float64                  `json:"map_price"`
	TaxClassID                  int                      `json:"tax_class_id"`
	ProductTaxCode              string                   `json:"product_tax_code"`
	CalculatedPrice             float64                  `json:"calculated_price"`
	Categories                  []int                    `json:"categories"`
	BrandID                     int                      `json:"brand_id"`
	OptionSetID                 int                      `json:"option_set_id"`
	OptionSetDisplay            string                   `json:"option_set_display"`
	InventoryLevel              int                      `json:"inventory_level"`
	InventoryWarningLevel       int                      `json:"inventory_warning_level"`
	InventoryTracking           string                   `json:"inventory_tracking"`
	ReviewsRatingSum            int                      `json:"reviews_rating_sum"`
	ReviewsCount                int                      `json:"reviews_count"`
	TotalSold                   int                      `json:"total_sold"`
	FixedCostShippingPrice      float64                  `json:"fixed_cost_shipping_price"`
	IsFreeShipping              bool                     `json:"is_free_shipping"`
	IsVisible                   bool                     `json:"is_visible"`
	IsFeatured                  bool                     `json:"is_featured"`
	RelatedProducts             []int                    `json:"related_products"`
	Warranty                    string                   `json:"warranty"`
	BinPickingNumber            string                   `json:"bin_picking_number"`
	LayoutFile                  string                   `json:"layout_file"`
	UPC                         string                   `json:"upc"`
	MPN                         string                   `json:"mpn"`
	GTIN                        string                   `json:"gtin"`
	SearchKeywords              string                   `json:"search_keywords"`
	Availability                string                   `json:"availability"`
	AvailabilityDescription     string                   `json:"availability_description"`
	GiftWrappingOptionsType     string                   `json:"gift_wrapping_options_type"`
	GiftWrappingOptionsList     []int                    `json:"gift_wrapping_options_list"`
	SortOrder                   int                      `json:"sort_order"`
	Condition                   string                   `json:"condition"`
	IsConditionShown            bool                     `json:"is_condition_shown"`
	OrderQuantityMinimum        int                      `json:"order_quantity_minimum"`
	OrderQuantityMaximum        int                      `json:"order_quantity_maximum"`
	PageTitle                   string                   `json:"page_title"`
	MetaKeywords                []string                 `json:"meta_keywords"`
	MetaDescription             string                   `json:"meta_description"`
	DateCreated                 string                   `json:"date_created"`
	DateModified                string                   `json:"date_modified"`
	ViewCount                   int                      `json:"view_count"`
	PreorderReleaseDate         string                   `json:"preorder_release_date"`
	PreorderMessage             string                   `json:"preorder_message"`
	IsPreorderOnly              bool                     `json:"is_preorder_only"`
	IsPriceHidden               bool                     `json:"is_price_hidden"`
	PriceHiddenLabel            string                   `json:"price_hidden_label"`
	CustomURL                   CustomURL                `json:"custom_url"`
	BaseVariantID               int                      `json:"base_variant_id"`
	OpenGraphType               string                   `json:"open_graph_type"`
	OpenGraphTitle              string                   `json:"open_graph_title"`
	OpenGraphDescription        string                   `json:"open_graph_description"`
	OpenGraphUseMetaDescription bool                     `json:"open_graph_use_meta_description"`
	OpenGraphUseProductName     bool                     `json:"open_graph_use_product_name"`
	OpenGraphUseImage           bool                     `json:"open_graph_use_image"`
	Variants                    []ProductVariant         `json:"variants"`
	Images                      []ProductImage           `json:"images"`
	CustomFields                []ProductCustomField     `json:"custom_fields"`
	BulkPricingRules            []ProductBulkPricingRule `json:"bulk_pricing_rules"`
	// TODO need appr. type
	PrimaryImage string `json:"primary_image"`
	// TODO need appr. type
	Modifiers string          `json:"modifiers"`
	Options   []ProductOption `json:"options"`
	Videos    []ProductVideo  `json:"video"`
}

type LimitedProductQueryParams struct {
	Include       []string `url:"include,omitempty,comma"`
	IncludeFields []string `url:"include_fields,omitempty,comma"`
	ExcludeFields []string `url:"exclude_fields,omitempty,comma"`
}

type ProductQueryParams struct {
	ID                    int      `url:"id,omitempty"`
	IDIn                  []int    `url:"id:in,omitempty,comma"`
	IDNotIn               []int    `url:"id:not_in,omitempty,comma"`
	IDMin                 []int    `url:"id:min,omitempty,comma"`
	IDMax                 []int    `url:"id:max,omitempty,comma"`
	IDGreater             []int    `url:"id:greater,omitempty,comma"`
	IDLess                []int    `url:"id:less,omitempty,comma"`
	Name                  string   `url:"name,omitempty"`
	UPC                   string   `url:"upc,omitempty"`
	Price                 float64  `url:"price,omitempty"`
	Weight                float64  `url:"weight,omitempty"`
	Condition             string   `url:"condition,omitempty"`
	BrandID               int      `url:"brand_id,omitempty"`
	DateModified          string   `url:"date_modified,omitempty"`
	DateModifiedMax       string   `url:"date_modified:max,omitempty"`
	DateModifiedMin       string   `url:"date_modified:min,omitempty"`
	DateLastImported      string   `url:"date_last_imported,omitempty"`
	DateLastImportedMax   string   `url:"date_last_imported:max,omitempty"`
	DateLastImportedMin   string   `url:"date_last_imported:min,omitempty"`
	IsVisible             bool     `url:"is_visible,omitempty"`
	IsFeatured            int      `url:"is_featured,omitempty"`
	IsFreeShipping        int      `url:"is_free_shipping,omitempty"`
	InventoryLevel        int      `url:"inventory_level,omitempty"`
	InventoryLevelIn      []int    `url:"inventory_level:in,omitempty,comma"`
	InventoryLevelNotIn   []int    `url:"inventory_level:not_in,omitempty,comma"`
	InventoryLevelMin     []int    `url:"inventory_level:min,omitempty,comma"`
	InventoryLevelMax     []int    `url:"inventory_level:max,omitempty,comma"`
	InventoryLevelGreater []int    `url:"inventory_level:greater,omitempty,comma"`
	InventoryLevelLess    []int    `url:"inventory_level:less,omitempty,comma"`
	InventoryLow          int      `url:"inventory_low,omitempty"`
	OutOfStock            int      `url:"out_of_stock,omitempty"`
	TotalSold             int      `url:"total_sold,omitempty"`
	Type                  string   `url:"type,omitempty"`
	Categories            int      `url:"categories,omitempty"`
	Keyword               string   `url:"keyword,omitempty"`
	KeywordContext        string   `url:"keyword_context,omitempty"`
	Status                int      `url:"status,omitempty"`
	Include               []string `url:"include,omitempty,comma"`
	IncludeFields         []string `url:"include_fields,omitempty,comma"`
	ExcludeFields         []string `url:"exclude_fields,omitempty,comma"`
	Availability          string   `url:"availability,omitempty"`
	Page                  int      `url:"page,omitempty"`
	Limit                 int      `url:"limit,omitempty"`
	Direction             string   `url:"direction,omitempty"`
	Sort                  string   `url:"sort,omitempty"`
	CategoriesIn          []int    `url:"categories:in,omitempty,comma"`
	SKU                   string   `url:"sku,omitempty"`
	SKUIn                 []string `url:"sku:in,omitempty,comma"`
}

type CreateProductParams struct {
	// Required Fields
	Name   string  `json:"name"`   // >= 1 character, <= 250 characters
	Type   string  `json:"type"`   // Allowed: physical, digital
	Weight float64 `json:"weight"` // Min: 0, Max: 9999999999
	Price  float64 `json:"price"`  // Min: 0

	// Optional Fields
	SKU                     string                   `json:"sku,omitempty"`                       // <= 255 characters
	Description             string                   `json:"description,omitempty"`               // HTML formatting allowed
	Width                   float64                  `json:"width,omitempty"`                     // Min: 0, Max: 9999999999
	Depth                   float64                  `json:"depth,omitempty"`                     // Min: 0, Max: 9999999999
	Height                  float64                  `json:"height,omitempty"`                    // Min: 0, Max: 9999999999
	CostPrice               float64                  `json:"cost_price,omitempty"`                // Min: 0
	RetailPrice             float64                  `json:"retail_price,omitempty"`              // Min: 0
	SalePrice               float64                  `json:"sale_price,omitempty"`                // Min: 0
	MAPPrice                float64                  `json:"map_price,omitempty"`                 // Min: 0
	TaxClassID              int                      `json:"tax_class_id,omitempty"`              // Min: 0, Max: 255
	ProductTaxCode          string                   `json:"product_tax_code,omitempty"`          // <= 255 characters
	Categories              []int                    `json:"categories,omitempty"`                // Max: 1000
	BrandID                 int                      `json:"brand_id,omitempty"`                  // Min: 0, Max: 1000000000
	BrandName               string                   `json:"brand_name,omitempty"`                // Non-case sensitive
	InventoryLevel          int                      `json:"inventory_level,omitempty"`           // Min: 0, Max: 2147483647
	InventoryWarningLevel   int                      `json:"inventory_warning_level,omitempty"`   // Min: 0, Max: 2147483647
	InventoryTracking       string                   `json:"inventory_tracking,omitempty"`        // Allowed: none, product, variant
	FixedCostShippingPrice  float64                  `json:"fixed_cost_shipping_price,omitempty"` // Min: 0
	IsFreeShipping          bool                     `json:"is_free_shipping,omitempty"`
	IsVisible               bool                     `json:"is_visible,omitempty"`
	IsFeatured              bool                     `json:"is_featured,omitempty"`
	RelatedProducts         []int                    `json:"related_products,omitempty"`
	Warranty                string                   `json:"warranty,omitempty"`                   // <= 65535 characters
	BinPickingNumber        string                   `json:"bin_picking_number,omitempty"`         // <= 255 characters
	LayoutFile              string                   `json:"layout_file,omitempty"`                // <= 500 characters
	UPC                     string                   `json:"upc,omitempty"`                        // <= 14 characters
	SearchKeywords          string                   `json:"search_keywords,omitempty"`            // <= 65535 characters
	AvailabilityDescription string                   `json:"availability_description,omitempty"`   // <= 255 characters
	Availability            string                   `json:"availability,omitempty"`               // Allowed: available, disabled, preorder
	GiftWrappingOptionsType string                   `json:"gift_wrapping_options_type,omitempty"` // Allowed: any, none, list
	GiftWrappingOptionsList []int                    `json:"gift_wrapping_options_list,omitempty"`
	SortOrder               int                      `json:"sort_order,omitempty"` // Min: -2147483648, Max: 2147483647
	Condition               string                   `json:"condition,omitempty"`  // Allowed: New, Used, Refurbished
	IsConditionShown        bool                     `json:"is_condition_shown,omitempty"`
	OrderQuantityMinimum    int                      `json:"order_quantity_minimum,omitempty"` // Min: 0, Max: 1000000000
	OrderQuantityMaximum    int                      `json:"order_quantity_maximum,omitempty"` // Min: 0, Max: 1000000000
	PageTitle               string                   `json:"page_title,omitempty"`             // <= 255 characters
	MetaKeywords            []string                 `json:"meta_keywords,omitempty"`          // <= 65535 characters
	MetaDescription         string                   `json:"meta_description,omitempty"`       // <= 65535 characters
	PreorderReleaseDate     *time.Time               `json:"preorder_release_date,omitempty"`
	PreorderMessage         string                   `json:"preorder_message,omitempty"` // <= 255 characters
	IsPreorderOnly          bool                     `json:"is_preorder_only,omitempty"`
	IsPriceHidden           bool                     `json:"is_price_hidden,omitempty"`
	PriceHiddenLabel        string                   `json:"price_hidden_label,omitempty"` // <= 200 characters
	CustomURL               *CustomURL               `json:"custom_url,omitempty"`
	OpenGraphType           string                   `json:"open_graph_type,omitempty"` // Allowed: product, album, book, etc.
	OpenGraphTitle          string                   `json:"open_graph_title,omitempty"`
	OpenGraphDescription    string                   `json:"open_graph_description,omitempty"`
	OpenGraphUseMetaDesc    bool                     `json:"open_graph_use_meta_description,omitempty"`
	OpenGraphUseProductName bool                     `json:"open_graph_use_product_name,omitempty"`
	OpenGraphUseImage       bool                     `json:"open_graph_use_image,omitempty"`
	GTIN                    string                   `json:"gtin,omitempty"` // <= 14 characters
	MPN                     string                   `json:"mpn,omitempty"`
	DateLastImported        string                   `json:"date_last_imported,omitempty"`
	ReviewsRatingSum        int                      `json:"reviews_rating_sum,omitempty"`
	ReviewsCount            int                      `json:"reviews_count,omitempty"`
	TotalSold               int                      `json:"total_sold,omitempty"`
	CustomFields            []ProductCustomField     `json:"custom_fields,omitempty"` // 200 maximum custom fields per product
	BulkPricingRules        []ProductBulkPricingRule `json:"bulk_pricing_rules,omitempty"`
	Images                  []ProductImage           `json:"images,omitempty"` // A product can have up to 1000 images
	Videos                  []ProductVideo           `json:"videos,omitempty"`
	Variants                []ProductVariant         `json:"variants,omitempty"`
}

type UpdateProductParams struct {
	Name                        string                   `json:"name,omitempty" validate:"required,min=1,max=250"`
	Type                        string                   `json:"type,omitempty" validate:"required,oneof=physical digital"`
	SKU                         string                   `json:"sku,omitempty" validate:"omitempty,min=0,max=255"`
	Description                 string                   `json:"description,omitempty"`
	Weight                      float64                  `json:"weight,omitempty" validate:"required,min=0,max=9999999999"`
	Width                       float64                  `json:"width,omitempty" validate:"omitempty,min=0,max=9999999999"`
	Depth                       float64                  `json:"depth,omitempty" validate:"omitempty,min=0,max=9999999999"`
	Height                      float64                  `json:"height,omitempty" validate:"omitempty,min=0,max=9999999999"`
	Price                       float64                  `json:"price,omitempty" validate:"required,min=0"`
	CostPrice                   float64                  `json:"cost_price,omitempty" validate:"omitempty,min=0"`
	RetailPrice                 float64                  `json:"retail_price,omitempty" validate:"omitempty,min=0"`
	SalePrice                   float64                  `json:"sale_price,omitempty" validate:"omitempty,min=0"`
	MapPrice                    float64                  `json:"map_price,omitempty" validate:"omitempty,min=0"`
	TaxClassID                  int                      `json:"tax_class_id,omitempty" validate:"omitempty,min=0,max=255"`
	ProductTaxCode              string                   `json:"product_tax_code,omitempty" validate:"omitempty,min=0,max=255"`
	Categories                  []int                    `json:"categories,omitempty" validate:"omitempty,min=0,max=1000,dive,min=0"`
	BrandID                     int                      `json:"brand_id,omitempty" validate:"omitempty,min=0,max=1000000000"`
	BrandName                   string                   `json:"brand_name,omitempty"`
	InventoryLevel              int                      `json:"inventory_level,omitempty" validate:"omitempty,min=0,max=2147483647"`
	InventoryWarningLevel       int                      `json:"inventory_warning_level,omitempty" validate:"omitempty,min=0,max=2147483647"`
	InventoryTracking           string                   `json:"inventory_tracking,omitempty" validate:"omitempty,oneof=none product variant"`
	FixedCostShippingPrice      float64                  `json:"fixed_cost_shipping_price,omitempty" validate:"omitempty,min=0"`
	IsFreeShipping              bool                     `json:"is_free_shipping,omitempty"`
	IsVisible                   bool                     `json:"is_visible,omitempty"`
	IsFeatured                  bool                     `json:"is_featured,omitempty"`
	RelatedProducts             []int                    `json:"related_products,omitempty"`
	Warranty                    string                   `json:"warranty,omitempty" validate:"omitempty,max=65535"`
	BinPickingNumber            string                   `json:"bin_picking_number,omitempty" validate:"omitempty,min=0,max=255"`
	LayoutFile                  string                   `json:"layout_file,omitempty" validate:"omitempty,min=0,max=500"`
	UPC                         string                   `json:"upc,omitempty" validate:"omitempty,min=0,max=14"`
	SearchKeywords              string                   `json:"search_keywords,omitempty" validate:"omitempty,min=0,max=65535"`
	AvailabilityDescription     string                   `json:"availability_description,omitempty" validate:"omitempty,min=0,max=255"`
	Availability                string                   `json:"availability,omitempty" validate:"omitempty,oneof=available disabled preorder"`
	GiftWrappingOptionsType     string                   `json:"gift_wrapping_options_type,omitempty" validate:"omitempty,oneof=any none list"`
	GiftWrappingOptionsList     []int                    `json:"gift_wrapping_options_list,omitempty"`
	SortOrder                   int                      `json:"sort_order,omitempty" validate:"omitempty,min=-2147483648,max=2147483647"`
	Condition                   string                   `json:"condition,omitempty" validate:"omitempty,oneof=New Used Refurbished"`
	IsConditionShown            bool                     `json:"is_condition_shown,omitempty"`
	OrderQuantityMinimum        int                      `json:"order_quantity_minimum,omitempty" validate:"omitempty,min=0,max=1000000000"`
	OrderQuantityMaximum        int                      `json:"order_quantity_maximum,omitempty" validate:"omitempty,min=0,max=1000000000"`
	PageTitle                   string                   `json:"page_title,omitempty" validate:"omitempty,min=0,max=255"`
	MetaKeywords                []string                 `json:"meta_keywords,omitempty" validate:"omitempty,dive,min=0,max=65535"`
	MetaDescription             string                   `json:"meta_description,omitempty" validate:"omitempty,min=0,max=65535"`
	PreorderReleaseDate         string                   `json:"preorder_release_date,omitempty"`
	PreorderMessage             string                   `json:"preorder_message,omitempty" validate:"omitempty,min=0,max=255"`
	IsPreorderOnly              bool                     `json:"is_preorder_only,omitempty"`
	IsPriceHidden               bool                     `json:"is_price_hidden,omitempty"`
	PriceHiddenLabel            string                   `json:"price_hidden_label,omitempty" validate:"omitempty,min=0,max=200"`
	CustomURL                   *CustomURL               `json:"custom_url,omitempty"`
	OpenGraphType               string                   `json:"open_graph_type,omitempty" validate:"omitempty,oneof=product album book drink food game movie song tv_show"`
	OpenGraphTitle              string                   `json:"open_graph_title,omitempty"`
	OpenGraphDescription        string                   `json:"open_graph_description,omitempty"`
	OpenGraphUseMetaDescription bool                     `json:"open_graph_use_meta_description,omitempty"`
	OpenGraphUseProductName     bool                     `json:"open_graph_use_product_name,omitempty"`
	OpenGraphUseImage           bool                     `json:"open_graph_use_image,omitempty"`
	GTIN                        string                   `json:"gtin,omitempty" validate:"omitempty,min=0,max=14"`
	MPN                         string                   `json:"mpn,omitempty"`
	DateLastImported            string                   `json:"date_last_imported,omitempty"`
	ReviewsRatingSum            int                      `json:"reviews_rating_sum,omitempty"`
	ReviewsCount                int                      `json:"reviews_count,omitempty"`
	TotalSold                   int                      `json:"total_sold,omitempty"`
	CustomFields                []ProductCustomField     `json:"custom_fields,omitempty"`
	BulkPricingRules            []ProductBulkPricingRule `json:"bulk_pricing_rules,omitempty"`
	Images                      []ProductImage           `json:"images,omitempty"`
	Videos                      []ProductVideo           `json:"videos,omitempty"`
	Variants                    []ProductVariant         `json:"variants,omitempty"`
}

type ProductBulkPricingRule struct {
	ID          int    `json:"id"`
	QuantityMin int    `json:"quantity_min"`
	QuantityMax int    `json:"quantity_max"`
	Type        string `json:"type"`
	Amount      string `json:"amount"`
}

// GetProduct retrieves a single product by its ID from the BigCommerce API.
//
// Parameters:
//   - id: The unique identifier of the product to retrieve.
//   - params: LimitedProductQueryParams to specify additional query parameters for the request.
//
// Returns:
//   - Product: The retrieved product information.
//   - error: An error if the request fails or if there's an issue processing the response.
func (client *V3Client) GetProduct(id int, params LimitedProductQueryParams) (Product, error) {
	var response ResponseObject

	// Add query parameters
	url, err := urlWithQueryParams(client.constructURL("/catalog/products/", strconv.Itoa(id)), params)
	if err != nil {
		return response.Data, err
	}

	// Send the request
	if err = client.Get(url, &response); err != nil {
		return response.Data, err
	}

	return response.Data, nil
}

// GetProductBySKU retrieves a product by its SKU (Stock Keeping Unit) from the BigCommerce API.
//
// This function first fetches all variants matching the given SKU, then retrieves the associated product.
// It ensures that exactly one variant is returned for the given SKU to avoid ambiguity.
//
// Parameters:
//   - sku: The Stock Keeping Unit (SKU) of the product to retrieve.
//
// Returns:
//   - Product: The retrieved product information.
//   - error: An error if the request fails, if no variants are found, if multiple variants are found,
//     or if there's an issue processing the response.
func (client *V3Client) GetProductBySKU(sku string) (Product, error) {
	// Fetch variants matching the SKU
	variants, err := client.GetAllVariants(AllProductVariantsQueryParams{SKU: sku})
	if err != nil {
		return Product{}, err
	}
	if len(variants) < 1 {
		return Product{}, errors.New("this sku returned no results")
	}
	if len(variants) > 1 {
		return Product{}, errors.New("this sku returned too many results")
	}
	// Retrieve the product associated with the variant
	product, err := client.GetProduct(variants[0].ProductID, LimitedProductQueryParams{})
	if err != nil {
		return Product{}, err
	}
	return product, nil
}

// TODO maybe change this to getproduct, getproducts and getAllProducts, and have the ability to pass params to get all products
// GetProductsByIDs retrieves multiple products by their IDs from the BigCommerce API.
//
// This function takes a slice of product IDs and fetches the corresponding products.
// It uses the ProductQueryParams to construct the API request, specifically utilizing
// the IDIn parameter to filter products by their IDs.
//
// Parameters:
//   - ids: A slice of integers representing the product IDs to retrieve.
//
// Returns:
//   - []Product: A slice of Product structs containing the retrieved product information.
//   - error: An error if the request fails or if there's an issue processing the response.
func (client *V3Client) GetProductsByIDs(ids []int) ([]Product, error) {
	params := ProductQueryParams{
		IDIn: ids,
	}

	products, _, err := client.GetProducts(params)
	if err != nil {
		return nil, err
	}

	return products, nil
}

// GetProducts retrieves a list of products from the BigCommerce API based on the provided query parameters.
//
// This function sends a GET request to the BigCommerce API's catalog/products endpoint with the specified
// query parameters. It returns a slice of Product structs, metadata about the request, and any error encountered.
//
// Parameters:
//   - params: ProductQueryParams struct containing various query parameters to filter and paginate the products.
//
// Returns:
//   - []Product: A slice of Product structs containing the retrieved product information.
//   - MetaData: Metadata about the API response, including pagination information.
//   - error: An error if the request fails, if there's an issue constructing the URL, or if there's a problem
//     processing the response.
func (client *V3Client) GetProducts(params ProductQueryParams) ([]Product, MetaData, error) {
	type ResponseObject struct {
		Data []Product `json:"data"`
		Meta MetaData  `json:"meta"`
	}
	var response ResponseObject

	getProductsUrl, err := urlWithQueryParams(client.constructURL("/catalog/products"), params)
	if err != nil {
		return response.Data, response.Meta, err
	}
	if err := client.Get(getProductsUrl, &response); err != nil {
		return response.Data, response.Meta, err
	}

	return response.Data, response.Meta, nil
}

// GetAllProducts retrieves all products from the BigCommerce API based on the provided query parameters.
//
// This function uses pagination to fetch all products that match the given criteria. It automatically
// handles multiple API requests to retrieve all pages of results.
//
// Parameters:
//   - params: ProductQueryParams struct containing various query parameters to filter the products.
//     The Page and Limit fields of this struct will be overwritten by this function.
//
// Returns:
//   - []Product: A slice of Product structs containing all retrieved product information.
//   - error: An error if any request fails or if there's an issue processing the responses.
func (client *V3Client) GetAllProducts(params ProductQueryParams) ([]Product, error) {
	var products []Product
	params.Page = 1
	params.Limit = 250
	for {
		p, _, err := client.GetProducts(params)
		if err != nil {
			return products, err
		}
		for i := 0; i < len(p); i++ {
			products = append(products, p[i])
		}

		if len(p) < params.Limit {
			return products, nil
		}

		params.Page++
	}
}

// ForEach applies a series of functions to each product in the BigCommerce catalog.
//
// This function iterates through all products in the catalog, applying each function
// in the provided slice to every product. If any function returns true, indicating
// that the product was modified, the product is updated via the API.
//
// Parameters:
//   - funcs: A slice of functions, each taking a pointer to a Product and returning a boolean.
//     The boolean indicates whether the product was modified (true) or not (false).
//
// Returns:
//   - error: An error if any API request fails or if there's an issue updating a product.
//
// Note:
//
//	This function uses pagination to process all products in batches of 250.
//	It will continue making API requests until all products have been processed.
func (client *V3Client) ForEachProduct(funcs []func(p *Product) bool) error {
	page := 1
	limit := 250
	for {

		if client.logger != nil {
			client.logger.Printf("Fetching page %d", page)
		}

		params := ProductQueryParams{Page: page, Limit: limit}
		batch, _, err := client.GetProducts(params)
		if err != nil {
			return err
		}

		for i := range batch {
			product := &batch[i]

			updated := false
			for _, fn := range funcs {
				if fn(product) {
					if client.logger != nil {
						client.logger.Printf("Product %d modified", product.ID)
					}
					updated = true
				}
			}

			// Compare modified product to original and update if changed
			if updated {
				updateParams := UpdateProductParams{
					Name:                        product.Name,
					Type:                        product.Type,
					SKU:                         product.SKU,
					Description:                 product.Description,
					Weight:                      product.Weight,
					Width:                       product.Width,
					Depth:                       product.Depth,
					Height:                      product.Height,
					Price:                       product.Price,
					CostPrice:                   product.CostPrice,
					RetailPrice:                 product.RetailPrice,
					SalePrice:                   product.SalePrice,
					MapPrice:                    product.MapPrice,
					TaxClassID:                  product.TaxClassID,
					ProductTaxCode:              product.ProductTaxCode,
					Categories:                  product.Categories,
					BrandID:                     product.BrandID,
					InventoryLevel:              product.InventoryLevel,
					InventoryWarningLevel:       product.InventoryWarningLevel,
					InventoryTracking:           product.InventoryTracking,
					FixedCostShippingPrice:      product.FixedCostShippingPrice,
					IsFreeShipping:              product.IsFreeShipping,
					IsVisible:                   product.IsVisible,
					IsFeatured:                  product.IsFeatured,
					RelatedProducts:             product.RelatedProducts,
					Warranty:                    product.Warranty,
					BinPickingNumber:            product.BinPickingNumber,
					LayoutFile:                  product.LayoutFile,
					UPC:                         product.UPC,
					SearchKeywords:              product.SearchKeywords,
					Availability:                product.Availability,
					AvailabilityDescription:     product.AvailabilityDescription,
					GiftWrappingOptionsType:     product.GiftWrappingOptionsType,
					GiftWrappingOptionsList:     product.GiftWrappingOptionsList,
					SortOrder:                   product.SortOrder,
					Condition:                   product.Condition,
					IsConditionShown:            product.IsConditionShown,
					OrderQuantityMinimum:        product.OrderQuantityMinimum,
					OrderQuantityMaximum:        product.OrderQuantityMaximum,
					PageTitle:                   product.PageTitle,
					MetaKeywords:                product.MetaKeywords,
					MetaDescription:             product.MetaDescription,
					PreorderReleaseDate:         product.PreorderReleaseDate,
					PreorderMessage:             product.PreorderMessage,
					IsPreorderOnly:              product.IsPreorderOnly,
					IsPriceHidden:               product.IsPriceHidden,
					PriceHiddenLabel:            product.PriceHiddenLabel,
					CustomURL:                   &product.CustomURL,
					OpenGraphType:               product.OpenGraphType,
					OpenGraphTitle:              product.OpenGraphTitle,
					OpenGraphDescription:        product.OpenGraphDescription,
					OpenGraphUseMetaDescription: product.OpenGraphUseMetaDescription,
					OpenGraphUseProductName:     product.OpenGraphUseProductName,
					OpenGraphUseImage:           product.OpenGraphUseImage,
					GTIN:                        product.GTIN,
					MPN:                         product.MPN,
					ReviewsRatingSum:            product.ReviewsRatingSum,
					ReviewsCount:                product.ReviewsCount,
					TotalSold:                   product.TotalSold,
					CustomFields:                product.CustomFields,
					BulkPricingRules:            product.BulkPricingRules,
					Images:                      product.Images,
					Videos:                      product.Videos,
					Variants:                    product.Variants,
				}

				_, err := client.UpdateProduct(product.ID, updateParams)
				if err != nil {
					return fmt.Errorf("failed to update product %d: %w", product.ID, err)
				}
				if client.logger != nil {
					client.logger.Printf("Product %d updated", product.ID)
				}
			}
		}
		if len(batch) < limit {
			break
		}
		page++
	}
	return nil
}

func (client *V3Client) UpdateProduct(productId int, params UpdateProductParams) (Product, error) {
	var response ResponseObject

	err := client.Put(client.constructURL("/catalog/products", strconv.Itoa(productId)), params, &response)
	if err != nil {
		return response.Data, err
	}

	return response.Data, nil
}

func (client *V3Client) CreateProduct(params CreateProductParams) (Product, error) {
	var response ResponseObject

	noNameSupplied := params.Name == ""
	invalidType := params.Type == "physical" || params.Type == "digital"
	invalidWeight := params.Weight <= 0

	if noNameSupplied || invalidType || invalidWeight {
		return response.Data, fmt.Errorf("failed check of name, type and weight")
	}

	if err := client.Post(client.constructURL("/catalog/products"), params, &response); err != nil {
		return response.Data, nil
	}

	return response.Data, nil
}

func (client *V3Client) DeleteProduct(productID int) error {
	err := client.Delete(client.constructURL("/catalog/products", strconv.Itoa(productID)), nil)
	if err != nil {
		return err
	}
	return nil
}

func (client *V3Client) RemoveCategoryFromProduct(productID, categoryToRemoveID int) (Product, error) {
	product, err := client.GetProduct(productID, LimitedProductQueryParams{})
	if err != nil {
		return product, err
	}

	categoriesToKeep := []int{}
	for i := 0; i < len(product.Categories); i++ {
		categoryID := product.Categories[i]
		if categoryID != categoryToRemoveID {
			categoriesToKeep = append(categoriesToKeep, categoryID)
		}
	}

	return client.UpdateProduct(productID, UpdateProductParams{Categories: categoriesToKeep})
}

func (client *V3Client) AddCategoryToProduct(productID, categoryToAddID int) (Product, error) {
	product, err := client.GetProduct(productID, LimitedProductQueryParams{})
	if err != nil {
		return product, err
	}
	updatedProductCategories := append(product.Categories, categoryToAddID)
	return client.UpdateProduct(productID, UpdateProductParams{Categories: updatedProductCategories})
}

func (p *Product) AddCategory(c int) []int {
	p.Categories = append(p.Categories, c)
	return p.Categories
}

func (p *Product) ContainsCategory(c int) bool {
	for _, id := range p.Categories {
		if id == c {
			return true
		}
	}
	return false
}

func (p *Product) RemoveCategory(c int) []int {
	idsToKeep := []int{}
	for _, id := range p.Categories {
		if id != c {
			idsToKeep = append(idsToKeep, id)
		}
	}
	p.Categories = idsToKeep
	return p.Categories
}
