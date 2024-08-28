package bigcommerce

import (
	"fmt"
	"strconv"
)

func (client *V2Client) GetOrderShipments(orderID int, params OrderShipmentQueryParams) ([]OrderShipment, MetaData, error) {
	type ResponseData struct {
		Data []OrderShipment `json:"data"`
		Meta MetaData        `json:"meta"`
	}
	var response ResponseData

	getOrdersURL, err := urlWithQueryParams(client.constructURL("orders", strconv.Itoa(orderID), "shipments"), params)
	if err != nil {
		return nil, MetaData{}, fmt.Errorf("failed to construct URL for GetOrderShipments (order ID: %d): %w", orderID, err)
	}

	if err := client.Get(getOrdersURL, &response.Data); err != nil {
		return nil, MetaData{}, fmt.Errorf("failed to get shipments for order %d: %w", orderID, err)
	}

	return response.Data, response.Meta, nil
}

type OrderShipmentQueryParams struct {
	Page  int `url:"page,omitempty"`
	Limit int `url:"limit,omitempty"`
}

type Address struct {
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Company     string `json:"company"`
	Street1     string `json:"street_1"`
	Street2     string `json:"street_2"`
	City        string `json:"city"`
	State       string `json:"state"`
	Zip         string `json:"zip"`
	Country     string `json:"country"`
	CountryISO2 string `json:"country_iso2"`
	Phone       string `json:"phone"`
	Email       string `json:"email"`
}

type Item struct {
	OrderProductID int `json:"order_product_id"`
	ProductID      int `json:"product_id"`
	Quantity       int `json:"quantity"`
}

type OrderShipment struct {
	ID                          int     `json:"id"`
	OrderID                     int     `json:"order_id"`
	CustomerID                  int     `json:"customer_id"`
	OrderAddressID              int     `json:"order_address_id"`
	DateCreated                 string  `json:"date_created"`
	TrackingNumber              string  `json:"tracking_number"`
	ShippingMethod              string  `json:"shipping_method"`
	ShippingProvider            string  `json:"shipping_provider"`
	TrackingCarrier             string  `json:"tracking_carrier"`
	TrackingLink                string  `json:"tracking_link"`
	Comments                    string  `json:"comments"`
	BillingAddress              Address `json:"billing_address"`
	ShippingAddress             Address `json:"shipping_address"`
	Items                       []Item  `json:"items"`
	ShippingProviderDisplayName string  `json:"shipping_provider_display_name"`
	GeneratedTrackingLink       string  `json:"generated_tracking_link"`
}
