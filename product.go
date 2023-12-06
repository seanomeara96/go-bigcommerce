package bigcommerce

import (
	"encoding/json"
	"fmt"
)

//TODO add a GetBrand/GetCategories function that returns brand or slice of categories arg can be pointer and if nil can call the api

func (client *Client) GetProduct(id int) (Product, error) {
	type ResponseObject struct {
		Data Product  `json:"data"`
		Meta MetaData `json:"meta"`
	}
	var response ResponseObject

	getProductUrl := client.BaseURL.JoinPath("/catalog/products/", fmt.Sprint(id)).String()

	// Send the request
	resp, err := client.Get(getProductUrl)
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

//TODO maybe change this to getproduct, getproducts and getAllProducts, and have the ability to pass params to get all products

func (client *Client) GetAllProducts(params ProductQueryParams) ([]Product, MetaData, error) {
	type ResponseObject struct {
		Data []Product `json:"data"`
		Meta MetaData  `json:"meta"`
	}
	var response ResponseObject

	queryParams, err := paramString(params)

	if err != nil {
		return response.Data, response.Meta, err
	}

	getProductsUrl := client.BaseURL.JoinPath("/catalog/products").String() + queryParams

	resp, err := client.Get(getProductsUrl)
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

func (client *Client) GetFullProductCatalog(limit int) ([]Product, error) {
	var products []Product
	page := 1
	end := false

	for !end {
		p, _, err := client.GetAllProducts(ProductQueryParams{Limit: limit, Page: page})
		if err != nil {
			return products, err
		}

		for i := 0; i < len(p); i++ {
			products = append(products, p[i])
		}

		if len(p) < limit {
			end = true
			break
		}

		page++
	}

	return products, nil
}

func (client *Client) UpdateProduct(productId int, params CreateUpdateProductParams) (Product, error) {
	type ResponseObject struct {
		Data Product  `json:"data"`
		Meta MetaData `json:"meta"`
	}
	var response ResponseObject

	updateProductPath := client.BaseURL.JoinPath("/catalog/products", fmt.Sprint(productId)).String()

	payloadBytes, err := json.Marshal(params)
	if err != nil {
		return response.Data, err
	}

	resp, err := client.Put(updateProductPath, payloadBytes)
	if err != nil {
		return response.Data, err
	}
	defer resp.Body.Close()

	return response.Data, nil
}

func (client *Client) CreateProduct(params CreateUpdateProductParams) (Product, error) {
	type ResponseObject struct {
		Data Product  `json:"data"`
		Meta MetaData `json:"meta"`
	}
	var response ResponseObject

	noNameSupplied := params.Name == ""
	invalidType := params.Type == "physical" || params.Type == "digital"
	invalidWeight := params.Weight <= 0

	if noNameSupplied || invalidType || invalidWeight {
		return response.Data, fmt.Errorf("failed check of name, type and weight")
	}

	createProductPath := client.BaseURL.JoinPath("/catalog/products").String()

	payloadBytes, err := json.Marshal(params)
	if err != nil {
		return response.Data, nil
	}

	resp, err := client.Post(createProductPath, payloadBytes)
	if err != nil {
		return response.Data, nil
	}

	err = expectStatusCode(200, resp)
	if err != nil {
		err = expectStatusCode(207, resp)
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

func (client *Client) DeleteProduct(productID int) error {
	path := client.BaseURL.JoinPath("/catalog/products", fmt.Sprint(productID)).String()
	resp, err := client.Delete(path)
	if err != nil {
		return err
	}
	err = expectStatusCode(204, resp)
	if err != nil {
		return err
	}

	return nil
}

func (client *Client) RemoveCategoryFromProduct(productID, categoryToRemoveID int) (Product, error) {
	product, err := client.GetProduct(productID)
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

	return client.UpdateProduct(productID, CreateUpdateProductParams{Categories: categoriesToKeep})
}

func (client *Client) AddCategoryToProduct(productID, categoryToAddID int) (Product, error) {
	product, err := client.GetProduct(productID)
	if err != nil {
		return product, err
	}
	updatedProductCategories := append(product.Categories, categoryToAddID)
	return client.UpdateProduct(productID, CreateUpdateProductParams{Categories: updatedProductCategories})
}

type Product struct {
	ID                      int       `json:"id"`
	Name                    string    `json:"name"`
	Type                    string    `json:"type"`
	SKU                     string    `json:"sku"`
	Description             string    `json:"description"`
	Weight                  float64   `json:"weight"`
	Width                   float64   `json:"width"`
	Depth                   float64   `json:"depth"`
	Height                  float64   `json:"height"`
	Price                   float64   `json:"price"`
	CostPrice               float64   `json:"cost_price"`
	RetailPrice             float64   `json:"retail_price"`
	SalePrice               float64   `json:"sale_price"`
	MapPrice                float64   `json:"map_price"`
	TaxClassID              float64   `json:"tax_class_id"`
	ProductTaxCode          string    `json:"product_tax_code"`
	CalculatedPrice         float64   `json:"calculated_price"`
	Categories              []int     `json:"categories"`
	BrandID                 int       `json:"brand_id"`
	OptionSetID             int       `json:"option_set_id"`
	OptionSetDisplay        string    `json:"option_set_display"`
	InventoryLevel          int       `json:"inventory_level"`
	InventoryWarningLevel   int       `json:"inventory_warning_level"`
	InventoryTracking       string    `json:"inventory_tracking"`
	ReviewsRatingSum        float64   `json:"reviews_rating_sum"`
	ReviewsCount            int       `json:"reviews_count"`
	TotalSold               int       `json:"total_sold"`
	FixedCostShippingPrice  float64   `json:"fixed_cost_shipping_price"`
	IsFreeShipping          bool      `json:"is_free_shipping"`
	IsVisible               bool      `json:"is_visible"`
	IsFeatured              bool      `json:"is_featured"`
	RelatedProducts         []int     `json:"related_products"`
	Warranty                string    `json:"warranty"`
	BinPickingNumber        string    `json:"bin_picking_number"`
	LayoutFile              string    `json:"layout_file"`
	UPC                     string    `json:"upc"`
	MPN                     string    `json:"mpn"`
	GTIN                    string    `json:"gtin"`
	SearchKeywords          string    `json:"search_keywords"`
	Availability            string    `json:"availability"`
	AvailabilityDescription string    `json:"availability_description"`
	GiftWrappingOptionsType string    `json:"gift_wrapping_options_type"`
	GiftWrappingOptionsList []string  `json:"gift_wrapping_options_list"`
	SortOrder               int       `json:"sort_order"`
	Condition               string    `json:"condition"`
	IsConditionShown        bool      `json:"is_condition_shown"`
	OrderQuantityMinimum    int       `json:"order_quantity_minimum"`
	OrderQuantityMaximum    int       `json:"order_quantity_maximum"`
	PageTitle               string    `json:"page_title"`
	MetaKeywords            []string  `json:"meta_keywords"`
	MetaDescription         string    `json:"meta_description"`
	DateCreated             string    `json:"date_created"`
	DateModified            string    `json:"date_modified"`
	ViewCount               int       `json:"view_count"`
	PreorderReleaseDate     string    `json:"preorder_release_date"`
	PreorderMessage         string    `json:"preorder_message"`
	IsPreorderOnly          bool      `json:"is_preorder_only"`
	IsPriceHidden           bool      `json:"is_price_hidden"`
	PriceHiddenLabel        string    `json:"price_hidden_label"`
	CustomURL               CustomURL `json:"custom_url"`
	BaseVariantID           int       `json:"base_variant_id"`
	OpenGraphType           string    `json:"open_graph_type"`
	OpenGraphTitle          string    `json:"open_graph_title"`
	OpenGraphDescription    string    `json:"open_graph_description"`
	OpenGraphUseMetaDesc    bool      `json:"open_graph_use_meta_description"`
	OpenGraphUseProductName bool      `json:"open_graph_use_product_name"`
	OpenGraphUseImage       bool      `json:"open_graph_use_image"`
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
	Include               string   `url:"include,omitempty"`
	IncludeFields         string   `url:"include_fields,omitempty"`
	ExcludeFields         string   `url:"exclude_fields,omitempty"`
	Availability          string   `url:"availability,omitempty"`
	Page                  int      `url:"page,omitempty"`
	Limit                 int      `url:"limit,omitempty"`
	Direction             string   `url:"direction,omitempty"`
	Sort                  string   `url:"sort,omitempty"`
	CategoriesIn          []int    `url:"categories:in,omitempty,comma"`
	SKU                   string   `url:"sku,omitempty"`
	SKUIn                 []string `url:"sku:in,omitempty,comma"`
}

type CreateUpdateProductParams struct {
	Name                    string                   `json:"name,omitempty" validate:"required,min=1,max=250"`
	Type                    string                   `json:"type,omitempty" validate:"required,oneof=physical digital"`
	SKU                     string                   `json:"sku,omitempty" validate:"max=255"`
	Description             string                   `json:"description,omitempty"`
	Weight                  float64                  `json:"weight,omitempty" validate:"required"`
	Width                   float64                  `json:"width,omitempty"`
	Depth                   float64                  `json:"depth,omitempty"`
	Height                  float64                  `json:"height,omitempty"`
	Price                   float64                  `json:"price,omitempty" validate:"required"`
	CostPrice               float64                  `json:"cost_price,omitempty"`
	RetailPrice             float64                  `json:"retail_price,omitempty"`
	SalePrice               float64                  `json:"sale_price,omitempty"`
	MapPrice                float64                  `json:"map_price,omitempty"`
	TaxClassID              int                      `json:"tax_class_id,omitempty"`
	ProductTaxCode          string                   `json:"product_tax_code,omitempty" validate:"max=255"`
	Categories              []int                    `json:"categories,omitempty"`
	BrandID                 int                      `json:"brand_id,omitempty"`
	InventoryLevel          int                      `json:"inventory_level,omitempty"`
	InventoryWarningLevel   int                      `json:"inventory_warning_level,omitempty"`
	InventoryTracking       string                   `json:"inventory_tracking,omitempty"`
	FixedCostShippingPrice  float64                  `json:"fixed_cost_shipping_price,omitempty"`
	IsFreeShipping          bool                     `json:"is_free_shipping,omitempty"`
	IsVisible               bool                     `json:"is_visible,omitempty"`
	IsFeatured              bool                     `json:"is_featured,omitempty"`
	RelatedProducts         []int                    `json:"related_products,omitempty"`
	Warranty                string                   `json:"warranty,omitempty"`
	BinPickingNumber        string                   `json:"bin_picking_number,omitempty"`
	LayoutFile              string                   `json:"layout_file,omitempty"`
	UPC                     string                   `json:"upc,omitempty"`
	SearchKeywords          string                   `json:"search_keywords,omitempty"`
	AvailabilityDescription string                   `json:"availability_description,omitempty"`
	Availability            string                   `json:"availability,omitempty"`
	GiftWrappingOptionsType string                   `json:"gift_wrapping_options_type,omitempty"`
	GiftWrappingOptionsList []int                    `json:"gift_wrapping_options_list,omitempty"`
	SortOrder               int                      `json:"sort_order,omitempty"`
	Condition               string                   `json:"condition,omitempty"`
	IsConditionShown        bool                     `json:"is_condition_shown,omitempty"`
	OrderQuantityMinimum    int                      `json:"order_quantity_minimum,omitempty"`
	OrderQuantityMaximum    int                      `json:"order_quantity_maximum,omitempty"`
	PageTitle               string                   `json:"page_title,omitempty"`
	MetaKeywords            []string                 `json:"meta_keywords,omitempty"`
	MetaDescription         string                   `json:"meta_description,omitempty"`
	ViewCount               int                      `json:"view_count,omitempty"`
	PreorderReleaseDate     string                   `json:"preorder_release_date,omitempty"`
	Message                 string                   `json:"preorder_message,omitempty"`
	IsPreorderOnly          bool                     `json:"is_preorder_only,omitempty"`
	IsPriceHidden           bool                     `json:"is_price_hidden,omitempty"`
	PriceHiddenLabel        string                   `json:"price_hidden_label,omitempty"`
	CustomURL               *CustomURL               `json:"custom_url,omitempty"`
	OpenGraphType           string                   `json:"open_graph_type,omitempty"`
	OpenGraphTitle          string                   `json:"open_graph_title,omitempty"`
	OpenGraphDescription    string                   `json:"open_graph_description,omitempty"`
	OpenGraphUseMetaDesc    bool                     `json:"open_graph_use_meta_description,omitempty"`
	OpenGraphUseProductName bool                     `json:"open_graph_use_product_name,omitempty"`
	OpenGraphUseImage       bool                     `json:"open_graph_use_image,omitempty"`
	BrandName               string                   `json:"brand_name,omitempty"`
	GTIN                    string                   `json:"gtin,omitempty"`
	MPN                     string                   `json:"mpn,omitempty"`
	ReviewsRatingSum        int                      `json:"reviews_rating_sum,omitempty"`
	ReviewsCount            int                      `json:"reviews_count,omitempty"`
	TotalSold               int                      `json:"total_sold,omitempty"`
	CustomFields            []ProductCustomField     `json:"custom_fields,omitempty"`
	BulkPricingRules        []ProductBulkPricingRule `json:"bulk_pricing_rules,omitempty"`
	Images                  []ProductImage           `json:"images,omitempty"`
	Videos                  []ProductVideo           `json:"videos,omitempty"`
	Variants                []ProductVariant         `json:"variants,omitempty"`
}

type ProductBulkPricingRule struct {
	ID          int    `json:"id"`
	QuantityMin int    `json:"quantity_min"`
	QuantityMax int    `json:"quantity_max"`
	Type        string `json:"type"`
	Amount      string `json:"amount"`
}
