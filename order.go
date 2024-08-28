package bigcommerce

import (
	"fmt"
	"strconv"
)

// SortField represents a field that can be used to sort orders
type OrderSortField string

// OrderSortDirection represents the direction to sort (ascending or descending)
type OrderSortDirection string

// Constants for OrderSortField
const (
	OrderSortFieldID           OrderSortField = "id"
	OrderSortFieldCustomerID   OrderSortField = "customer_id"
	OrderSortFieldDateCreated  OrderSortField = "date_created"
	OrderSortFieldDateModified OrderSortField = "date_modified"
	OrderSortFieldStatusID     OrderSortField = "status_id"
	OrderSortFieldChannelID    OrderSortField = "channel_id"
	OrderSortFieldExternalID   OrderSortField = "external_id"
)

// Constants for OrderSortDirection
const (
	OrderSortDirectionAsc  OrderSortDirection = "asc"
	OrderSortDirectionDesc OrderSortDirection = "desc"
)

// SortQuery represents a sort query with field and direction
type OrderSortQuery struct {
	Field     OrderSortField
	Direction OrderSortDirection
}

// String returns the string representation of the SortQuery
func (s OrderSortQuery) String() string {
	return fmt.Sprintf("%s:%s", s.Field, s.Direction)
}

type Order struct {
	ID                                      int            `json:"id"`
	CustomerID                              int            `json:"customer_id"`
	DateCreated                             string         `json:"date_created"`
	DateModified                            string         `json:"date_modified"`
	DateShipped                             string         `json:"date_shipped"`
	StatusID                                int            `json:"status_id"`
	Status                                  string         `json:"status"`
	SubtotalExTax                           string         `json:"subtotal_ex_tax"`
	SubtotalIncTax                          string         `json:"subtotal_inc_tax"`
	SubtotalTax                             string         `json:"subtotal_tax"`
	BaseShippingCost                        string         `json:"base_shipping_cost"`
	ShippingCostExTax                       string         `json:"shipping_cost_ex_tax"`
	ShippingCostIncTax                      string         `json:"shipping_cost_inc_tax"`
	ShippingCostTax                         string         `json:"shipping_cost_tax"`
	ShippingCostTaxClassID                  int            `json:"shipping_cost_tax_class_id"`
	BaseHandlingCost                        string         `json:"base_handling_cost"`
	HandlingCostExTax                       string         `json:"handling_cost_ex_tax"`
	HandlingCostIncTax                      string         `json:"handling_cost_inc_tax"`
	HandlingCostTax                         string         `json:"handling_cost_tax"`
	HandlingCostTaxClassID                  int            `json:"handling_cost_tax_class_id"`
	BaseWrappingCost                        string         `json:"base_wrapping_cost"`
	WrappingCostExTax                       string         `json:"wrapping_cost_ex_tax"`
	WrappingCostIncTax                      string         `json:"wrapping_cost_inc_tax"`
	WrappingCostTax                         string         `json:"wrapping_cost_tax"`
	WrappingCostTaxClassID                  int            `json:"wrapping_cost_tax_class_id"`
	TotalExTax                              string         `json:"total_ex_tax"`
	TotalIncTax                             string         `json:"total_inc_tax"`
	TotalTax                                string         `json:"total_tax"`
	ItemsTotal                              int            `json:"items_total"`
	ItemsShipped                            int            `json:"items_shipped"`
	PaymentMethod                           string         `json:"payment_method"`
	PaymentProviderID                       string         `json:"payment_provider_id"`
	PaymentStatus                           string         `json:"payment_status"`
	RefundedAmount                          string         `json:"refunded_amount"`
	OrderIsDigital                          bool           `json:"order_is_digital"`
	StoreCreditAmount                       string         `json:"store_credit_amount"`
	GiftCertificateAmount                   string         `json:"gift_certificate_amount"`
	IPAddress                               string         `json:"ip_address"`
	IPAddressV6                             string         `json:"ip_address_v6"`
	GeoIPCountry                            string         `json:"geoip_country"`
	GeoIPCountryISO2                        string         `json:"geoip_country_iso2"`
	CurrencyID                              int            `json:"currency_id"`
	CurrencyCode                            string         `json:"currency_code"`
	CurrencyExchangeRate                    string         `json:"currency_exchange_rate"`
	DefaultCurrencyID                       int            `json:"default_currency_id"`
	DefaultCurrencyCode                     string         `json:"default_currency_code"`
	StaffNotes                              string         `json:"staff_notes"`
	CustomerMessage                         string         `json:"customer_message"`
	DiscountAmount                          string         `json:"discount_amount"`
	CouponDiscount                          string         `json:"coupon_discount"`
	ShippingAddressCount                    int            `json:"shipping_address_count"`
	IsDeleted                               bool           `json:"is_deleted"`
	EbayOrderID                             string         `json:"ebay_order_id"`
	CartID                                  string         `json:"cart_id"`
	BillingAddress                          BillingAddress `json:"billing_address"`
	IsEmailOptIn                            bool           `json:"is_email_opt_in"`
	CreditCardType                          interface{}    `json:"credit_card_type"`
	OrderSource                             string         `json:"order_source"`
	ChannelID                               int            `json:"channel_id"`
	ExternalSource                          interface{}    `json:"external_source"`
	Products                                URLResource    `json:"products"`
	ShippingAddresses                       URLResource    `json:"shipping_addresses"`
	Coupons                                 URLResource    `json:"coupons"`
	ExternalID                              interface{}    `json:"external_id"`
	ExternalMerchantID                      interface{}    `json:"external_merchant_id"`
	TaxProviderID                           string         `json:"tax_provider_id"`
	StoreDefaultCurrencyCode                string         `json:"store_default_currency_code"`
	StoreDefaultToTransactionalExchangeRate string         `json:"store_default_to_transactional_exchange_rate"`
	CustomStatus                            string         `json:"custom_status"`
	CustomerLocale                          string         `json:"customer_locale"`
	ExternalOrderID                         string         `json:"external_order_id"`
}

type BillingAddress struct {
	FirstName   string       `json:"first_name"`
	LastName    string       `json:"last_name"`
	Company     string       `json:"company"`
	Street1     string       `json:"street_1"`
	Street2     string       `json:"street_2"`
	City        string       `json:"city"`
	State       string       `json:"state"`
	Zip         string       `json:"zip"`
	Country     string       `json:"country"`
	CountryISO2 string       `json:"country_iso2"`
	Phone       string       `json:"phone"`
	Email       string       `json:"email"`
	FormFields  []FormFields `json:"form_fields"`
}

type FormFields struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type URLResource struct {
	URL      string `json:"url"`
	Resource string `json:"resource"`
}

type OrderQueryParams struct {
	ID              int      `url:"id,omitempty"`
	IDIn            []int    `url:"id:in,omitempty"`
	IDNotIn         []int    `url:"id:not_in,omitempty"`
	IDMin           []int    `url:"id:min,omitempty"`
	IDMax           []int    `url:"id:max,omitempty"`
	IDGreater       []int    `url:"id:greater,omitempty"`
	IDLess          []int    `url:"id:less,omitempty"`
	Name            string   `url:"name,omitempty"`
	NameLike        []string `url:"name:like,omitempty"`
	ParentID        int      `url:"parent_id,omitempty"`
	ParentIDIn      []int    `url:"parent_id:in,omitempty"`
	ParentIDMin     []int    `url:"parent_id:min,omitempty"`
	ParentIDMax     []int    `url:"parent_id:max,omitempty"`
	ParentIDGreater []int    `url:"parent_id:greater,omitempty"`
	ParentIDLess    []int    `url:"parent_id:less,omitempty"`
	PageTitle       string   `url:"page_title,omitempty"`
	PageTitleLike   []string `url:"page_title:like,omitempty"`
	Keyword         string   `url:"keyword,omitempty"`
	IsVisible       bool     `url:"is_visible,omitempty"`
	Page            int      `url:"page,omitempty"`
	Limit           int      `url:"limit,omitempty"`
	IncludeFields   string   `url:"include_fields,omitempty"`
	ExcludeFields   string   `url:"exclude_fields,omitempty"`
	MinID           int      `url:"min_id,omitempty"`
	MaxID           int      `url:"max_id,omitempty"`
	MinTotal        float64  `url:"min_total,omitempty"`
	MaxTotal        float64  `url:"max_total,omitempty"`
	CustomerID      int      `url:"customer_id,omitempty"`
	Email           string   `url:"email,omitempty"`
	StatusID        int      `url:"status_id,omitempty"`
	CartID          string   `url:"cart_id,omitempty"`
	PaymentMethod   string   `url:"payment_method,omitempty"`
	MinDateCreated  string   `url:"min_date_created,omitempty"`
	MaxDateCreated  string   `url:"max_date_created,omitempty"`
	MinDateModified string   `url:"min_date_modified,omitempty"`
	MaxDateModified string   `url:"max_date_modified,omitempty"`
	Sort            string   `url:"sort,omitempty"`
	IsDeleted       bool     `url:"is_deleted,omitempty"`
	ChannelID       int      `url:"channel_id,omitempty"`
}

func (client *Client) GetOrder(orderID int) (Order, error) {
	type ResponseObject struct {
		Data Order    `json:"data"`
		Meta MetaData `json:"meta"`
	}
	var response ResponseObject

	err := client.Version2Required()
	if err != nil {
		return Order{}, fmt.Errorf("API version 2 is required: %w", err)
	}

	getOrderURL := client.constructURL("storefront", "orders", strconv.Itoa(orderID))

	if err := client.Get(getOrderURL, &response); err != nil {
		return Order{}, fmt.Errorf("failed to get order with ID %d: %w", orderID, err)
	}

	return response.Data, nil
}

func (client *Client) GetOrders(params OrderQueryParams) ([]Order, MetaData, error) {
	type ResponseData struct {
		Data []Order  `json:"data"`
		Meta MetaData `json:"meta"`
	}
	var response ResponseData

	err := client.Version2Required()
	if err != nil {
		return nil, MetaData{}, fmt.Errorf("API version 2 is required: %w", err)
	}

	getOrdersURL, err := urlWithQueryParams(client.constructURL("orders"), params)
	if err != nil {
		return nil, MetaData{}, fmt.Errorf("failed to construct URL with query params: %w", err)
	}

	if err := client.Get(getOrdersURL, &response); err != nil {
		return nil, MetaData{}, fmt.Errorf("failed to get orders: %w", err)
	}

	return response.Data, response.Meta, nil
}
