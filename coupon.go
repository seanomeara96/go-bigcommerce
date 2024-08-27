package bigcommerce

import (
	"encoding/json"
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
	// Validate required fields
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

	// Additional validation logic for specific fields

	return nil // No validation errors
}

func (client *Client) CreateCoupon(params CreateUpdateCouponParams) (Coupon, error) {
	var response CouponResponseObject
	err := client.Version2Required()
	if err != nil {
		return response.Data, err
	}
	err = validateCreateUpdateCoupon(params)
	if err != nil {
		return response.Data, err
	}
	paramBytes, err := json.Marshal(params)
	if err != nil {
		return response.Data, err
	}
	path := client.constructURL("coupons")
	resp, err := client.Post(path, paramBytes)
	if err != nil {
		return response.Data, err
	}
	defer resp.Body.Close()
	err = expectStatusCode(201, resp)
	if err != nil {
		return response.Data, err
	}
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return response.Data, err
	}
	return response.Data, nil
}

func (client *Client) UpdateCoupon(couponID int, params CreateUpdateCouponParams) (Coupon, error) {
	var response CouponResponseObject
	if err := client.Version2Required(); err != nil {
		return response.Data, err
	}
	err := validateCreateUpdateCoupon(params)
	if err != nil {
		return response.Data, err
	}

	path := client.constructURL("coupons", strconv.Itoa(couponID))
	resp, err := client.Put(path, params)
	if err != nil {
		return response.Data, err
	}
	defer resp.Body.Close()
	err = expectStatusCode(200, resp)
	if err != nil {
		return response.Data, err
	}
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return response.Data, err
	}
	return response.Data, nil
}

func (client *Client) GetCoupons(params CouponQueryParams) ([]Coupon, MetaData, error) {
	var response CouponsResponseObject
	err := client.Version2Required()
	if err != nil {
		return response.Data, response.Meta, err
	}

	path, err := urlWithQueryParams(client.constructURL("coupons"), params)
	if err != nil {
		return response.Data, response.Meta, err
	}

	resp, err := client.Get(path)
	if err != nil {
		return response.Data, response.Meta, err
	}
	defer resp.Body.Close()
	err = expectStatusCode(200, resp)
	if err != nil {
		return response.Data, response.Meta, err
	}
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return response.Data, response.Meta, err
	}
	return response.Data, response.Meta, nil
}

func (client *Client) GetCoupon(couponID int) (Coupon, error) {
	var response CouponResponseObject
	err := client.Version2Required()
	if err != nil {
		return response.Data, err
	}
	path := client.constructURL("coupons", fmt.Sprint(couponID))
	resp, err := client.Get(path)
	if err != nil {
		return response.Data, err
	}
	defer resp.Body.Close()
	err = expectStatusCode(200, resp)
	if err != nil {
		return response.Data, err
	}
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return response.Data, err
	}
	return response.Data, nil
}
func (client *Client) DeleteCoupon(couponID int) error {
	path := client.constructURL("coupons", fmt.Sprint(couponID))

	resp, err := client.Delete(path)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	err = expectStatusCode(204, resp)
	if err != nil {
		return err
	}
	return nil
}
