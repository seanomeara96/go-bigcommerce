package bigcommerce

import (
	"fmt"
)

// OrderStatus represents the structure of each order status.
type OrderStatus struct {
	ID                int    `json:"id"`
	Name              string `json:"name"`
	SystemLabel       string `json:"system_label"`
	CustomLabel       string `json:"custom_label"`
	SystemDescription string `json:"system_description"`
}

func (client *Client) GetOrderStatuses() ([]OrderStatus, error) {

	type ResponseObject struct {
		Data []OrderStatus `json:"data"`
		Meta MetaData      `json:"meta"`
	}
	var response ResponseObject

	err := client.Version2Required()
	if err != nil {
		return nil, fmt.Errorf("version 2 required for GetOrderStatuses: %w", err)
	}

	path := client.constructURL("order_statuses")

	if err := client.Get(path, &response); err != nil {
		return nil, fmt.Errorf("failed to get order statuses: %w", err)
	}

	return response.Data, nil
}
