package bigcommerce

import (
	"encoding/json"
	"strconv"
)

type ShippingAddressQueryParams struct {
	OrderID int `url:"order_id" validate:"required"`
	Page    int `url:"page,omitempty"`
	Limit   int `url:"limit,omitempty"`
}

type ShippingAddressFormField struct {
	Name  string      `json:"name"`
	Value interface{} `json:"value"` // Value can be number, string, or array
}

type ShippingQuotes struct {
	URL      string `json:"url"`      // Read-only
	Resource string `json:"resource"` // Read-only
}

type ShippingAddress struct {
	ID                     int                        `json:"id"`
	OrderID                int                        `json:"order_id"`
	ItemsTotal             float64                    `json:"items_total"`
	ItemsShipped           float64                    `json:"items_shipped"`
	BaseCost               string                     `json:"base_cost"`
	CostExTax              string                     `json:"cost_ex_tax"`
	CostIncTax             string                     `json:"cost_inc_tax"`
	CostTax                string                     `json:"cost_tax"`
	CostTaxClassID         int                        `json:"cost_tax_class_id"`
	BaseHandlingCost       string                     `json:"base_handling_cost"`
	HandlingCostExTax      string                     `json:"handling_cost_ex_tax"`
	HandlingCostIncTax     string                     `json:"handling_cost_inc_tax"`
	HandlingCostTax        string                     `json:"handling_cost_tax"`
	HandlingCostTaxClassID int                        `json:"handling_cost_tax_class_id"`
	ShippingZoneID         float64                    `json:"shipping_zone_id"`
	ShippingZoneName       string                     `json:"shipping_zone_name"`
	FormFields             []ShippingAddressFormField `json:"form_fields"`
	ShippingQuotes         ShippingQuotes             `json:"shipping_quotes"` // Read-only
	FirstName              string                     `json:"first_name"`
	LastName               string                     `json:"last_name"`
	Company                string                     `json:"company"`
	Street1                string                     `json:"street_1"`
	Street2                string                     `json:"street_2"`
	City                   string                     `json:"city"`
	State                  string                     `json:"state"`
	Zip                    string                     `json:"zip"`
	Country                string                     `json:"country"`
	CountryISO2            string                     `json:"country_iso2"`
	Phone                  string                     `json:"phone"`
	Email                  string                     `json:"email"`
	ShippingMethod         string                     `json:"shipping_method"`
}

func (client *Client) GetOrderShippingAddress(orderID int, params ShippingAddressQueryParams) ([]ShippingAddress, MetaData, error) {
	type ResponseData struct {
		Data []ShippingAddress `json:"data"`
		Meta MetaData          `json:"meta"`
	}
	var response ResponseData

	err := client.Version2Required()
	if err != nil {
		return response.Data, response.Meta, err
	}

	queryParams, err := paramString(params)
	if err != nil {
		return response.Data, response.Meta, err
	}

	getOrdersURL, err := urlWithQueryParams(client.constructURL("orders", strconv.Itoa(orderID), "shipping_addresses"), queryParams)
	if err != nil {
		return response.Data, response.Meta, err
	}

	resp, err := client.Get(getOrdersURL)
	if err != nil {
		return response.Data, response.Meta, err
	}
	defer resp.Body.Close()

	if err := expectStatusCode(200, resp); err != nil {
		return response.Data, response.Meta, err
	}

	if err := json.NewDecoder(resp.Body).Decode(&response.Data); err != nil {
		return response.Data, response.Meta, err
	}

	return response.Data, response.Meta, nil
}
