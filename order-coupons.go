package bigcommerce

import (
	"encoding/json"
	"fmt"
)

type OrderCoupon struct {
	ID       int     `json:"id"`
	CouponID int     `json:"coupon_id"`
	OrderID  int     `json:"order_id"`
	Code     string  `json:"code"`
	Amount   int     `json:"amount"`
	Type     int     `json:"type"`
	Discount float64 `json:"discount"`
}

func (orderCoupon *OrderCoupon) TypeName() string {
	var typeName string
	switch orderCoupon.Type {
	case 0:
		typeName = "per_item_discount"
	case 1:
		typeName = "percentage_discount"
	case 2:
		typeName = "per_total_discount"
	case 3:
		typeName = "shipping_discount"
	case 4:
		typeName = "free_shipping"
	case 5:
		typeName = "promotion"
	default:
		typeName = "unknown"
	}
	return typeName
}

func (client *Client) ListOrderCoupons(orderID int) ([]OrderCoupon, error) {
	type ResponseObject struct {
		Data []OrderCoupon `json:"data"`
		Meta MetaData      `json:"meta"`
	}
	var response ResponseObject

	listOrderCouponsPath := client.BaseURL().JoinPath("/orders/", fmt.Sprint(orderID), "/coupons").String()

	resp, err := client.Get(listOrderCouponsPath)
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
