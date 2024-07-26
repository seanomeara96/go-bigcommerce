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

func (client *Client) GetOrderStatuses() ([]OrderStatus, error) {

	type ResponseObject struct {
		Data []OrderStatus `json:"data"`
		Meta MetaData      `json:"meta"`
	}
	var response ResponseObject

	err := client.Version2Required()
	if err != nil {
		return []OrderStatus{}, nil
	}

	path := client.BaseURL().JoinPath("/order_statuses").String()

	resp, err := client.Get(path)
	if err != nil {
		return response.Data, err
	}
	defer resp.Body.Close()

	if err := expectStatusCode(200, resp); err != nil {
		return response.Data, err
	}

	if err := json.NewDecoder(resp.Body).Decode(&response.Data); err != nil {
		return response.Data, err
	}

	return response.Data, nil
}
