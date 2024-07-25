package bigcommerce

import (
	"encoding/json"
)

// OrderStatus represents the structure of each order status.
type OrderStatus struct {
	ID                int    `json:"id"`
	Name              string `json:"name"`
	SystemLabel       string `json:"system_label"`
	CustomLabel       string `json:"custom_label"`
	SystemDescription string `json:"system_description"`
}

// GetAllOrderStatusResponse represents the response for getting all order statuses.
type GetAllOrderStatusResponse struct {
	OrderStatuses []OrderStatus `json:"order_statuses"`
}

func (client *Client) GetOrderStatuses() (GetAllOrderStatusResponse, error) {

	type ResponseObject struct {
		Data GetAllOrderStatusResponse `json:"data"`
		Meta MetaData                  `json:"meta"`
	}
	var response ResponseObject

	err := client.Version2Required()
	if err != nil {
		return GetAllOrderStatusResponse{}, nil
	}

	path := client.BaseURL().JoinPath("/storefront/order_statuses").String()

	resp, err := client.Get(path)
	if err != nil {
		return response.Data, err
	}
	defer resp.Body.Close()

	if err := json.NewDecoder(resp.Body).Decode(&response.Data); err != nil {
		return response.Data, err
	}

	return response.Data, nil
}
