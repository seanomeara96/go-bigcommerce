package bigcommerce

import (
	"errors"
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
	ID                 int                 `json:"id"`
	DateCreated        string              `json:"date_created"`
	NumUses            int                 `json:"num_uses"`
	Name               string              `json:"name"`
	Type               string              `json:"type"`
	Amount             string              `json:"amount"`
	MinPurchase        string              `json:"min_purchase"`
	Expires            string              `json:"expires"`
	Enabled            bool                `json:"enabled"`
	Code               string              `json:"code"`
	AppliesTo          CouponAppliesTo     `json:"applies_to"`
	MaxUses            int                 `json:"max_uses"`
	MaxUsesPerCustomer int                 `json:"max_uses_per_customer"`
	RestrictedTo       *CouponRestrictedTo `json:"restricted_to,omitempty"`
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
type CreateUpdateCouponParams struct {
	Name               string        `json:"name"`
	Type               string        `json:"type"`
	Amount             string        `json:"amount"`
	MinPurchase        string        `json:"min_purchase,omitempty"`
	Expires            string        `json:"expires,omitempty"`
	Enabled            bool          `json:"enabled"`
	Code               string        `json:"code"`
	AppliesTo          AppliesTo     `json:"applies_to"`
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

func validateCreateUpdateCoupon(coupon CreateUpdateCouponParams) error {
	if coupon.Name == "" {
		return errors.New("name is required")
	}
	if coupon.Type == "" {
		return errors.New("type is required")
	}
	if coupon.Amount == "" {
		return errors.New("amount is required")
	}
	if coupon.Enabled && coupon.Code == "" {
		return errors.New("code is required when the coupon is enabled")
	}
	if len(coupon.AppliesTo.IDs) == 0 {
		return errors.New("at least one ID is required in AppliesTo")
	}
	if coupon.MaxUses < 0 {
		return errors.New("maxUses must be a non-negative value")
	}
	if coupon.MaxUsesPerCustomer < 0 {
		return errors.New("maxUsesPerCustomer must be a non-negative value")
	}

	return nil
}

func (client *Client) CreateCoupon(params CreateUpdateCouponParams) (Coupon, error) {
	var response CouponResponseObject
	err := client.Version2Required()
	if err != nil {
		return response.Data, fmt.Errorf("version 2 required: %w", err)
	}

	err = validateCreateUpdateCoupon(params)
	if err != nil {
		return response.Data, fmt.Errorf("coupon validation failed: %w", err)
	}

	path := client.constructURL("coupons")
	if err := client.Post(path, params, &response.Data); err != nil {
		return response.Data, fmt.Errorf("failed to create coupon: %w", err)
	}

	return response.Data, nil
}

func (client *Client) UpdateCoupon(couponID int, params CreateUpdateCouponParams) (Coupon, error) {
	var response CouponResponseObject

	if err := client.Version2Required(); err != nil {
		return response.Data, fmt.Errorf("version 2 required: %w", err)
	}

	err := validateCreateUpdateCoupon(params)
	if err != nil {
		return response.Data, fmt.Errorf("coupon validation failed: %w", err)
	}

	path := client.constructURL("coupons", strconv.Itoa(couponID))
	if err := client.Put(path, params, &response.Data); err != nil {
		return response.Data, fmt.Errorf("failed to update coupon with ID %d: %w", couponID, err)
	}

	return response.Data, nil
}

func (client *Client) GetCoupons(params CouponQueryParams) ([]Coupon, MetaData, error) {
	var response CouponsResponseObject

	err := client.Version2Required()
	if err != nil {
		return response.Data, response.Meta, fmt.Errorf("version 2 required: %w", err)
	}

	path, err := urlWithQueryParams(client.constructURL("coupons"), params)
	if err != nil {
		return response.Data, response.Meta, fmt.Errorf("failed to construct URL with query params: %w", err)
	}

	if err := client.Get(path, &response.Data); err != nil {
		return response.Data, response.Meta, fmt.Errorf("failed to get coupons: %w", err)
	}

	return response.Data, response.Meta, nil
}

func (client *Client) GetCoupon(couponID int) (Coupon, error) {
	var response CouponResponseObject

	err := client.Version2Required()
	if err != nil {
		return response.Data, fmt.Errorf("version 2 required: %w", err)
	}

	path := client.constructURL("coupons", strconv.Itoa(couponID))
	if err := client.Get(path, &response.Data); err != nil {
		return response.Data, fmt.Errorf("failed to get coupon with ID %d: %w", couponID, err)
	}

	return response.Data, nil
}

func (client *Client) DeleteCoupon(couponID int) error {

	err := client.Version2Required()
	if err != nil {
		return fmt.Errorf("version 2 required: %w", err)
	}

	path := client.constructURL("coupons", strconv.Itoa(couponID))

	if err := client.Delete(path, nil); err != nil {
		return fmt.Errorf("failed to delete coupon with ID %d: %w", couponID, err)
	}

	return nil
}
