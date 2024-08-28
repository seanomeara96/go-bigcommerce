package bigcommerce

import (
	"fmt"
	"strconv"
)

// OrderProductAppliedDiscount represents a discount applied to the product
type OrderProductAppliedDiscount struct {
	ID     string  `json:"id"`
	Amount string  `json:"amount"`
	Name   string  `json:"name"`
	Code   *string `json:"code"`
	Target string  `json:"target"`
}

// ProductOption represents an option applied to the product
type ProductOption struct {
	ID                   int    `json:"id"`
	OptionID             int    `json:"option_id"`
	OrderProductID       int    `json:"order_product_id"`
	ProductOptionID      int    `json:"product_option_id"`
	DisplayName          string `json:"display_name"`
	DisplayNameCustomer  string `json:"display_name_customer"`
	DisplayNameMerchant  string `json:"display_name_merchant"`
	DisplayValue         string `json:"display_value"`
	DisplayValueCustomer string `json:"display_value_customer"`
	DisplayValueMerchant string `json:"display_value_merchant"`
	Value                string `json:"value"`
	Type                 string `json:"type"`
	Name                 string `json:"name"`
	DisplayStyle         string `json:"display_style"`
}

// OrderProduct represents the structure of an order product
type OrderProduct struct {
	ID                    int                           `json:"id"`
	OrderID               int                           `json:"order_id"`
	ProductID             int                           `json:"product_id"`
	OrderAddressID        int                           `json:"order_address_id"`
	Name                  string                        `json:"name"`
	NameCustomer          string                        `json:"name_customer"`
	NameMerchant          string                        `json:"name_merchant"`
	SKU                   string                        `json:"sku"`
	UPC                   string                        `json:"upc"`
	Type                  string                        `json:"type"`
	BasePrice             string                        `json:"base_price"`
	PriceExTax            string                        `json:"price_ex_tax"`
	PriceIncTax           string                        `json:"price_inc_tax"`
	PriceTax              string                        `json:"price_tax"`
	BaseTotal             string                        `json:"base_total"`
	TotalExTax            string                        `json:"total_ex_tax"`
	TotalIncTax           string                        `json:"total_inc_tax"`
	TotalTax              string                        `json:"total_tax"`
	Weight                string                        `json:"weight"`
	Quantity              int                           `json:"quantity"`
	BaseCostPrice         string                        `json:"base_cost_price"`
	CostPriceIncTax       string                        `json:"cost_price_inc_tax"`
	CostPriceExTax        string                        `json:"cost_price_ex_tax"`
	CostPriceTax          string                        `json:"cost_price_tax"`
	IsRefunded            bool                          `json:"is_refunded"`
	QuantityRefunded      int                           `json:"quantity_refunded"`
	RefundAmount          string                        `json:"refund_amount"`
	ReturnID              int                           `json:"return_id"`
	WrappingName          string                        `json:"wrapping_name"`
	BaseWrappingCost      string                        `json:"base_wrapping_cost"`
	WrappingCostExTax     string                        `json:"wrapping_cost_ex_tax"`
	WrappingCostIncTax    string                        `json:"wrapping_cost_inc_tax"`
	WrappingCostTax       string                        `json:"wrapping_cost_tax"`
	WrappingMessage       string                        `json:"wrapping_message"`
	QuantityShipped       int                           `json:"quantity_shipped"`
	EventName             *string                       `json:"event_name"`
	EventDate             *string                       `json:"event_date"`
	FixedShippingCost     string                        `json:"fixed_shipping_cost"`
	EbayItemID            string                        `json:"ebay_item_id"`
	EbayTransactionID     string                        `json:"ebay_transaction_id"`
	OptionSetID           *int                          `json:"option_set_id"`
	ParentOrderProductID  *int                          `json:"parent_order_product_id"`
	IsBundledProduct      bool                          `json:"is_bundled_product"`
	BinPickingNumber      string                        `json:"bin_picking_number"`
	ExternalID            *string                       `json:"external_id"`
	FulfillmentSource     string                        `json:"fulfillment_source"`
	Brand                 string                        `json:"brand"`
	DiscountedTotalIncTax string                        `json:"discounted_total_inc_tax"`
	AppliedDiscounts      []OrderProductAppliedDiscount `json:"applied_discounts"`
	ProductOptions        []ProductOption               `json:"product_options"`
	ConfigurableFields    []interface{}                 `json:"configurable_fields"`
	GiftCertificateID     *int                          `json:"gift_certificate_id"`
}

type OrderProductsQueryParams struct {
	Page  int `url:"page,omitempty"`
	Limit int `url:"limit,omitempty"`
}

func (client *Client) GetOrderProducts(orderID int, params OrderProductsQueryParams) ([]OrderProduct, MetaData, error) {
	type ResponseData struct {
		Data []OrderProduct `json:"data"`
		Meta MetaData       `json:"meta"`
	}
	var response ResponseData

	err := client.Version2Required()
	if err != nil {
		return nil, MetaData{}, fmt.Errorf("version 2 required for GetOrderProducts: %w", err)
	}

	getOrdersURL, err := urlWithQueryParams(client.constructURL("orders", strconv.Itoa(orderID), "products"), params)
	if err != nil {
		return nil, MetaData{}, fmt.Errorf("failed to construct URL for GetOrderProducts (order ID: %d): %w", orderID, err)
	}

	if err := client.Get(getOrdersURL, &response); err != nil {
		return nil, MetaData{}, fmt.Errorf("failed to get products for order %d: %w", orderID, err)
	}

	return response.Data, response.Meta, nil
}
