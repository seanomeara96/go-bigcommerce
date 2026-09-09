package bigcommerce

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strconv"
)

// StringOrNumber preserves API values that BigCommerce may encode as either a
// JSON string or number.
type StringOrNumber string

func (value *StringOrNumber) UnmarshalJSON(data []byte) error {
	data = bytes.TrimSpace(data)
	if len(data) == 0 {
		return fmt.Errorf("expected a string or number")
	}

	if data[0] == '"' {
		var stringValue string
		if err := json.Unmarshal(data, &stringValue); err != nil {
			return fmt.Errorf("decode string: %w", err)
		}
		*value = StringOrNumber(stringValue)
		return nil
	}

	var numberValue json.Number
	if err := json.Unmarshal(data, &numberValue); err != nil {
		return fmt.Errorf("expected a string or number: %w", err)
	}
	if numberValue == "" {
		return fmt.Errorf("expected a string or number")
	}

	*value = StringOrNumber(numberValue.String())
	return nil
}

type OrderProductDiscountTarget string

const (
	OrderProductDiscountTargetOrder   OrderProductDiscountTarget = "order"
	OrderProductDiscountTargetProduct OrderProductDiscountTarget = "product"
)

type ProductOptionType string

const (
	ProductOptionTypeCheckbox        ProductOptionType = "Checkbox"
	ProductOptionTypeDateField       ProductOptionType = "Date field"
	ProductOptionTypeFileUpload      ProductOptionType = "File Upload"
	ProductOptionTypeMultiLineText   ProductOptionType = "Multi-line text field"
	ProductOptionTypeMultipleChoice  ProductOptionType = "Multiple choice"
	ProductOptionTypeProductPickList ProductOptionType = "Product Pick List"
	ProductOptionTypeSwatch          ProductOptionType = "Swatch"
	ProductOptionTypeTextField       ProductOptionType = "Text field"
)

type OrderProductType string

const (
	OrderProductTypePhysical        OrderProductType = "physical"
	OrderProductTypeDigital         OrderProductType = "digital"
	OrderProductTypeGiftCertificate OrderProductType = "giftcertificate"
)

// OrderProductAppliedDiscount represents a discount applied to the product
type OrderProductAppliedDiscount struct {
	ID     string                     `json:"id"`
	Amount string                     `json:"amount"`
	Name   string                     `json:"name"`
	Code   *string                    `json:"code"`
	Target OrderProductDiscountTarget `json:"target"`
}

// ProductOption represents an option applied to the product
type ProductOption struct {
	ID                   int               `json:"id"`
	OptionID             int               `json:"option_id"`
	OrderProductID       int               `json:"order_product_id"`
	ProductOptionID      int               `json:"product_option_id"`
	DisplayName          string            `json:"display_name"`
	DisplayNameCustomer  string            `json:"display_name_customer"`
	DisplayNameMerchant  string            `json:"display_name_merchant"`
	DisplayValue         string            `json:"display_value"`
	DisplayValueCustomer string            `json:"display_value_customer"`
	DisplayValueMerchant string            `json:"display_value_merchant"`
	Value                string            `json:"value"`
	Type                 ProductOptionType `json:"type"`
	Name                 string            `json:"name"`
	DisplayStyle         string            `json:"display_style"`
}

// OrderProductConfigurableField is a legacy Catalog V2 configurable field.
type OrderProductConfigurableField struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

// OrderProduct represents the structure of an order product
type OrderProduct struct {
	ID                   int              `json:"id"`
	OrderID              int              `json:"order_id"`
	ProductID            int              `json:"product_id"`
	OrderPickupMethodID  int              `json:"order_pickup_method_id"`
	OrderAddressID       int              `json:"order_address_id"`
	Name                 string           `json:"name"`
	NameCustomer         string           `json:"name_customer"`
	NameMerchant         string           `json:"name_merchant"`
	SKU                  string           `json:"sku"`
	UPC                  string           `json:"upc"`
	Type                 OrderProductType `json:"type"`
	BasePrice            string           `json:"base_price"`
	PriceExTax           string           `json:"price_ex_tax"`
	PriceIncTax          string           `json:"price_inc_tax"`
	PriceTax             string           `json:"price_tax"`
	BaseTotal            string           `json:"base_total"`
	TotalExTax           string           `json:"total_ex_tax"`
	TotalIncTax          string           `json:"total_inc_tax"`
	TotalTax             string           `json:"total_tax"`
	Weight               StringOrNumber   `json:"weight"`
	Width                string           `json:"width"`
	Height               string           `json:"height"`
	Depth                string           `json:"depth"`
	Quantity             int              `json:"quantity"`
	BaseCostPrice        string           `json:"base_cost_price"`
	CostPriceIncTax      string           `json:"cost_price_inc_tax"`
	CostPriceExTax       string           `json:"cost_price_ex_tax"`
	CostPriceTax         string           `json:"cost_price_tax"`
	IsRefunded           bool             `json:"is_refunded"`
	QuantityRefunded     float64          `json:"quantity_refunded"`
	RefundAmount         string           `json:"refund_amount"`
	ReturnID             float64          `json:"return_id"`
	WrappingID           int              `json:"wrapping_id"`
	WrappingName         *string          `json:"wrapping_name"`
	BaseWrappingCost     StringOrNumber   `json:"base_wrapping_cost"`
	WrappingCostExTax    string           `json:"wrapping_cost_ex_tax"`
	WrappingCostIncTax   string           `json:"wrapping_cost_inc_tax"`
	WrappingCostTax      string           `json:"wrapping_cost_tax"`
	WrappingMessage      string           `json:"wrapping_message"`
	QuantityShipped      float64          `json:"quantity_shipped"`
	QuantityBackordered  int              `json:"quantity_backordered"`
	BackorderMessage     *string          `json:"backorder_message"`
	EventName            *string          `json:"event_name"`
	EventDate            *string          `json:"event_date"`
	FixedShippingCost    string           `json:"fixed_shipping_cost"`
	EbayItemID           string           `json:"ebay_item_id"`
	EbayTransactionID    string           `json:"ebay_transaction_id"`
	OptionSetID          *int             `json:"option_set_id"`
	ParentOrderProductID *int             `json:"parent_order_product_id"`
	IsBundledProduct     bool             `json:"is_bundled_product"`
	BinPickingNumber     string           `json:"bin_picking_number"`
	ExternalID           *string          `json:"external_id"`
	// FulfillmentSource is returned by some stores but is not in the public schema.
	FulfillmentSource     string                          `json:"fulfillment_source"`
	Brand                 string                          `json:"brand"`
	DiscountedTotalIncTax string                          `json:"discounted_total_inc_tax"`
	AppliedDiscounts      []OrderProductAppliedDiscount   `json:"applied_discounts"`
	ProductOptions        []ProductOption                 `json:"product_options"`
	ConfigurableFields    []OrderProductConfigurableField `json:"configurable_fields"`
	GiftCertificateID     *int                            `json:"gift_certificate_id"`
	VariantID             *int                            `json:"variant_id"`
}

type OrderProductsQueryParams struct {
	Page  int `url:"page,omitempty"`
	Limit int `url:"limit,omitempty"`
}

func (client *V2Client) GetOrderProducts(orderID int, params OrderProductsQueryParams) ([]OrderProduct, MetaData, error) {
	type ResponseData struct {
		Data []OrderProduct `json:"data"`
		Meta MetaData       `json:"meta"`
	}
	var response ResponseData

	getOrdersURL, err := urlWithQueryParams(client.constructURL("orders", strconv.Itoa(orderID), "products"), params)
	if err != nil {
		return nil, MetaData{}, fmt.Errorf("failed to construct URL for GetOrderProducts (order ID: %d): %w", orderID, err)
	}

	if err := client.Get(getOrdersURL, &response.Data); err != nil {
		return nil, MetaData{}, fmt.Errorf("failed to get products for order %d: %w", orderID, err)
	}

	return response.Data, response.Meta, nil
}
