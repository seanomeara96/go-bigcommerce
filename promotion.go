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
