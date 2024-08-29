package bigcommerce

import (
	"fmt"
	"strconv"
)

type CouponQueryParams struct {
	ID          string `url:"id,omitempty"`
	Code        string `url:"code,omitempty"`
	Name        string `url:"name,omitempty"`
	Type        string `url:"type,omitempty"`
	MinID       int    `url:"min_id,omitempty"`
	MaxID       int    `url:"max_id,omitempty"`
	Page        int    `url:"page,omitempty"`
	Limit       int    `url:"limit,omitempty"`
	ExcludeType string `url:"exclude_type,omitempty"`
}
type Coupon struct {
	ID                 int                  `json:"id"`
	DateCreated        string               `json:"date_created"`
	NumUses            int                  `json:"num_uses"`
	Name               string               `json:"name"`
	Type               string               `json:"type"`
	Amount             string               `json:"amount"`
	MinPurchase        string               `json:"min_purchase"`
	Expires            string               `json:"expires"`
	Enabled            bool                 `json:"enabled"`
	Code               string               `json:"code"`
	AppliesTo          CouponAppliesTo      `json:"applies_to"`
	MaxUses            int                  `json:"max_uses"`
	MaxUsesPerCustomer int                  `json:"max_uses_per_customer"`
	RestrictedTo       []CouponRestrictedTo `json:"restricted_to"`
}

type CouponAppliesTo struct {
	IDs    []int  `json:"ids"`
	Entity string `json:"entity"`
}

type CouponRestrictedTo struct {
	Countries       string   `json:"countries,omitempty"`
	Null            string   `json:"null,omitempty"`
	ShippingMethods []string `json:"shipping_methods,omitempty"`
}

type CouponResponseObject struct {
	Data Coupon   `json:"data"`
	Meta MetaData `json:"meta"`
}

type CouponsResponseObject struct {
	Data []Coupon `json:"data"`
	Meta MetaData `json:"meta"`
}
type CreateCouponParams struct {
	Name               string        `json:"name"`
	Type               string        `json:"type"`
	Amount             string        `json:"amount"`
	MinPurchase        string        `json:"min_purchase,omitempty"`
	Expires            string        `json:"expires,omitempty"`
	Enabled            bool          `json:"enabled"`
	Code               string        `json:"code"`
	AppliesTo          *AppliesTo    `json:"applies_to"`
	MaxUses            int           `json:"max_uses,omitempty"`
	MaxUsesPerCustomer int           `json:"max_uses_per_customer,omitempty"`
	RestrictedTo       *RestrictedTo `json:"restricted_to,omitempty"`
}

type UpdateCouponParams struct {
	Name               string        `json:"name,omitempty"`
	Type               string        `json:"type,omitempty"`
	Amount             string        `json:"amount,omitempty"`
	MinPurchase        string        `json:"min_purchase,omitempty"`
	Expires            string        `json:"expires,omitempty"`
	Enabled            bool          `json:"enabled,omitempty"`
	Code               string        `json:"code,omitempty"`
	AppliesTo          *AppliesTo    `json:"applies_to,omitempty"`
	MaxUses            int           `json:"max_uses,omitempty"`
	MaxUsesPerCustomer int           `json:"max_uses_per_customer,omitempty"`
	RestrictedTo       *RestrictedTo `json:"restricted_to,omitempty"`
}

type AppliesTo struct {
	IDs    []int  `json:"ids"`
	Entity string `json:"entity"`
}

type RestrictedTo struct {
	ShippingMethods []string `json:"shipping_methods,omitempty"`
}

func (client *V2Client) CreateCoupon(params CreateCouponParams) (Coupon, error) {
	var response CouponResponseObject

	path := client.constructURL("coupons")
	if err := client.Post(path, params, &response.Data); err != nil {
		return response.Data, fmt.Errorf("failed to create coupon: %w", err)
	}

	return response.Data, nil
}

func (client *V2Client) UpdateCoupon(couponID int, params UpdateCouponParams) (Coupon, error) {
	var response CouponResponseObject

	path := client.constructURL("coupons", strconv.Itoa(couponID))
	if err := client.Put(path, params, &response.Data); err != nil {
		return response.Data, fmt.Errorf("failed to update coupon with ID %d: %w", couponID, err)
	}

	return response.Data, nil
}

func (client *V2Client) GetCoupons(params CouponQueryParams) ([]Coupon, error) {
	var response CouponsResponseObject

	path, err := urlWithQueryParams(client.constructURL("coupons"), params)
	if err != nil {
		return response.Data, fmt.Errorf("failed to construct URL with query params: %w", err)
	}

	if err := client.Get(path, &response.Data); err != nil {
		return response.Data, fmt.Errorf("failed to get coupons: %w", err)
	}

	return response.Data, nil
}

func (client *V2Client) GetCoupon(couponID int) (Coupon, error) {
	var response CouponResponseObject

	path := client.constructURL("coupons", strconv.Itoa(couponID))
	if err := client.Get(path, &response.Data); err != nil {
		return response.Data, fmt.Errorf("failed to get coupon with ID %d: %w", couponID, err)
	}

	return response.Data, nil
}

func (client *V2Client) DeleteCoupon(couponID int) error {

	path := client.constructURL("coupons", strconv.Itoa(couponID))

	if err := client.Delete(path, nil); err != nil {
		return fmt.Errorf("failed to delete coupon with ID %d: %w", couponID, err)
	}

	return nil
}
