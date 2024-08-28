package bigcommerce

import (
	"fmt"
	"strconv"
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

	listOrderCouponsPath := client.constructURL("orders", strconv.Itoa(orderID), "coupons")

	err := client.Get(listOrderCouponsPath, &response)
	if err != nil {
		return nil, fmt.Errorf("failed to list coupons for order %d: %w", orderID, err)
	}

	return response.Data, nil
}
