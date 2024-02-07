package bigcommerce

import (
	"encoding/json"
	"strconv"
)

type Promotion struct {
	ID             int    `json:"id"`
	RedemptionType string `json:"redemption_type"`
	Name           string `json:"name"`
	Channels       []struct {
		ID int `json:"id"`
	} `json:"channels"`
	Customer struct {
		GroupIDs          []int `json:"group_ids"`
		MinimumOrderCount int   `json:"minimum_order_count"`
		ExcludedGroupIDs  []int `json:"excluded_group_ids"`
	} `json:"customer"`
	Segments struct {
		IDs   []string `json:"id"`
		Rules []struct {
			Action    string `json:"action"`
			CartValue struct {
				Discount  string `json:"discount"`
				ApplyOnce bool   `json:"apply_once"`
				Stop      bool   `json:"stop"`
			} `json:"cart_value"`
			Condition struct {
				Cart struct {
					Items struct {
						Brands          []int  `json:"brands"`
						MinimumSpend    string `json:"minimum_spend"`
						MinimumQuantity int    `json:"minimum_quantity"`
					} `json:"items"`
				} `json:"cart"`
				CurrentUses int `json:"current_uses"`
				MaxUses     int `json:"max_uses"`
			} `json:"condition"`
		} `json:"rules"`
	} `json:"segments"`
	Status             string `json:"status"`
	StartDate          string `json:"start_date"`
	EndDate            string `json:"end_date"`
	Stop               bool   `json:"stop"`
	CanBeUsedWithOther bool   `json:"can_be_used_with_other_promotions"`
	CurrencyCode       string `json:"currency_code"`
	Notifications      []struct {
		Content   string   `json:"content"`
		Type      string   `json:"type"`
		Locations []string `json:"locations"`
	} `json:"notifications"`
	ShippingAddress struct {
		Countries []struct{} `json:"countries"`
	} `json:"shipping_address"`
	Schedule struct {
		WeekFrequency  int      `json:"week_frequency"`
		WeekDays       []string `json:"week_days"`
		DailyStartTime string   `json:"daily_start_time"`
		DailyEndTime   string   `json:"daily_end_time"`
	} `json:"schedule"`
	CouponOverridesAutomatic bool `json:"coupon_overrides_automatic_when_offering_higher_discounts"`
}

func (c *Client) GetPromotion(id int) (Promotion, error) {
	type Response struct {
		Data Promotion `json:"data"`
		Meta MetaData  `json:"meta"`
	}
	var response Response
	path := c.BaseURL.JoinPath("/promotions/" + strconv.Itoa(id)).String()
	resp, err := c.Get(path)
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

type PromotionUpdateParams struct {
	Name     string `json:"name,omitempty" url:"name,omitempty"`
	Channels []struct {
		ID int `json:"id,omitempty" url:"id,omitempty"`
	} `json:"channels,omitempty" url:"channels,omitempty"`

	Customer struct {
		GroupIDs          []int `json:"group_ids,omitempty" url:"group_ids,omitempty"`
		MinimumOrderCount int   `json:"minimum_order_count,omitempty" url:"minimum_order_count,omitempty"`
		ExcludedGroupIDs  []int `json:"excluded_group_ids,omitempty" url:"excluded_group_ids,omitempty"`
		Segments          struct {
			ID []string `json:"id,omitempty" url:"id,omitempty"`
		} `json:"segments,omitempty" url:"segments,omitempty"`
	} `json:"customer,omitempty" url:"customer,omitempty"`

	Rules []struct {
		Action string `json:"action,omitempty" url:"action,omitempty"`

		CartValue struct {
			Discount struct {
				FixedAmount string `json:"fixed_amount,omitempty" url:"fixed_amount,omitempty"`
			} `json:"discount,omitempty" url:"discount,omitempty"`
		} `json:"cart_value,omitempty" url:"cart_value,omitempty"`

		ApplyOnce bool `json:"apply_once,omitempty" url:"apply_once,omitempty"`
		Stop      bool `json:"stop,omitempty" url:"stop,omitempty"`

		Condition struct {
			Cart struct {
				Items struct {
					Brands []int `json:"brands,omitempty" url:"brands,omitempty"`
				} `json:"items,omitempty" url:"items,omitempty"`

				MinimumSpend    string `json:"minimum_spend,omitempty" url:"minimum_spend,omitempty"`
				MinimumQuantity int    `json:"minimum_quantity,omitempty" url:"minimum_quantity,omitempty"`
			} `json:"cart,omitempty" url:"cart,omitempty"`
		} `json:"condition,omitempty" url:"condition,omitempty"`
	} `json:"rules,omitempty" url:"rules,omitempty"`

	MaxUses                      int    `json:"max_uses,omitempty" url:"max_uses,omitempty"`
	Status                       string `json:"status,omitempty" url:"status,omitempty"`
	StartDate                    string `json:"start_date,omitempty" url:"start_date,omitempty"`
	EndDate                      string `json:"end_date,omitempty" url:"end_date,omitempty"`
	Stop                         bool   `json:"stop,omitempty" url:"stop,omitempty"`
	CanBeUsedWithOtherPromotions bool   `json:"can_be_used_with_other_promotions,omitempty" url:"can_be_used_with_other_promotions,omitempty"`
	CurrencyCode                 string `json:"currency_code,omitempty" url:"currency_code,omitempty"`

	Notifications []struct {
		Content   string   `json:"content,omitempty" url:"content,omitempty"`
		Type      string   `json:"type,omitempty" url:"type,omitempty"`
		Locations []string `json:"locations,omitempty" url:"locations,omitempty"`
	} `json:"notifications,omitempty" url:"notifications,omitempty"`

	ShippingAddress struct {
		Countries []struct {
			ISO2CountryCode string `json:"iso2_country_code,omitempty" url:"iso2_country_code,omitempty"`
		} `json:"countries,omitempty" url:"countries,omitempty"`
	} `json:"shipping_address,omitempty" url:"shipping_address,omitempty"`

	Schedule struct {
		WeekFrequency  int      `json:"week_frequency,omitempty" url:"week_frequency,omitempty"`
		WeekDays       []string `json:"week_days,omitempty" url:"week_days,omitempty"`
		DailyStartTime string   `json:"daily_start_time,omitempty" url:"daily_start_time,omitempty"`
		DailyEndTime   string   `json:"daily_end_time,omitempty" url:"daily_end_time,omitempty"`
	} `json:"schedule,omitempty" url:"schedule,omitempty"`

	CouponOverridesAutomaticWhenOfferingHigherDiscounts bool `json:"coupon_overrides_automatic_when_offering_higher_discounts,omitempty" url:"coupon_overrides_automatic_when_offering_higher_discounts,omitempty"`
}

func (c *Client) UpdatePromotion(id int, params PromotionUpdateParams) (Promotion, error) {
	type Response struct {
		Data Promotion `json:"data"`
		Meta MetaData  `json:"meta"`
	}
	var response Response

	path := c.BaseURL.JoinPath("/promotions/" + strconv.Itoa(id)).String()
	payloadBytes, err := json.Marshal(params)
	if err != nil {
		return response.Data, err
	}

	resp, err := c.Put(path, payloadBytes)
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
