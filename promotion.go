package bigcommerce

import (
	"strconv"
)

type Promotion struct {
	ID                       int                      `json:"id"`
	RedemptionType           string                   `json:"redemption_type"`
	Name                     string                   `json:"name"`
	Channels                 []PromotionChannel       `json:"channels"`
	Customer                 PromotionCustomer        `json:"customer"`
	Segments                 PromotionSegments        `json:"segments"`
	Status                   string                   `json:"status"`
	StartDate                string                   `json:"start_date"`
	EndDate                  string                   `json:"end_date"`
	Stop                     bool                     `json:"stop"`
	CanBeUsedWithOther       bool                     `json:"can_be_used_with_other_promotions"`
	CurrencyCode             string                   `json:"currency_code"`
	Notifications            []PromotionNotification  `json:"notifications"`
	ShippingAddress          PromotionShippingAddress `json:"shipping_address"`
	Schedule                 PromotionSchedule        `json:"schedule"`
	CouponOverridesAutomatic bool                     `json:"coupon_overrides_automatic_when_offering_higher_discounts"`
}

type PromotionSchedule struct {
	WeekFrequency  int      `json:"week_frequency"`
	WeekDays       []string `json:"week_days"`
	DailyStartTime string   `json:"daily_start_time"`
	DailyEndTime   string   `json:"daily_end_time"`
}

type PromotionShippingAddress struct {
	Countries []struct{} `json:"countries"`
}

type PromotionNotification struct {
	Content   string   `json:"content"`
	Type      string   `json:"type"`
	Locations []string `json:"locations"`
}
type PromotionChannel struct {
	ID int `json:"id"`
}

type PromotionSegmentRuleCartValue struct {
	Discount  string `json:"discount"`
	ApplyOnce bool   `json:"apply_once"`
	Stop      bool   `json:"stop"`
}

type PromotionSegmentRuleConditionCartItems struct {
	Brands          []int  `json:"brands"`
	MinimumSpend    string `json:"minimum_spend"`
	MinimumQuantity int    `json:"minimum_quantity"`
}

type PromotionSegmentRuleConditionCart struct {
	Items PromotionSegmentRuleConditionCartItems `json:"items"`
}

type PromotionSegmentRuleCondition struct {
	Cart        PromotionSegmentRuleConditionCart `json:"cart"`
	CurrentUses int                               `json:"current_uses"`
	MaxUses     int                               `json:"max_uses"`
}

type PromotionSegmentRule struct {
	Action    string                        `json:"action"`
	CartValue PromotionSegmentRuleCartValue `json:"cart_value"`
	Condition PromotionSegmentRuleCondition `json:"condition"`
}

type PromotionSegments struct {
	IDs   []string               `json:"id"`
	Rules []PromotionSegmentRule `json:"rules"`
}

type PromotionCustomer struct {
	GroupIDs          []int `json:"group_ids"`
	MinimumOrderCount int   `json:"minimum_order_count"`
	ExcludedGroupIDs  []int `json:"excluded_group_ids"`
}

func (c *Client) GetPromotion(id int) (Promotion, error) {
	type Response struct {
		Data Promotion `json:"data"`
		Meta MetaData  `json:"meta"`
	}

	var response Response
	path := c.constructURL("promotions", strconv.Itoa(id))
	err := c.Get(path, &response)
	if err != nil {
		return response.Data, err
	}

	return response.Data, nil
}

type PromotionUpdateParams struct {
	Name                                                string                `json:"name,omitempty" url:"name,omitempty"`
	Channels                                            []PromotionChannel    `json:"channels,omitempty" url:"channels,omitempty"`
	Customer                                            CustomerParams        `json:"customer,omitempty" url:"customer,omitempty"`
	Rules                                               []RuleParams          `json:"rules,omitempty" url:"rules,omitempty"`
	MaxUses                                             int                   `json:"max_uses,omitempty" url:"max_uses,omitempty"`
	Status                                              string                `json:"status,omitempty" url:"status,omitempty"`
	StartDate                                           string                `json:"start_date,omitempty" url:"start_date,omitempty"`
	EndDate                                             string                `json:"end_date,omitempty" url:"end_date,omitempty"`
	Stop                                                bool                  `json:"stop,omitempty" url:"stop,omitempty"`
	CanBeUsedWithOtherPromotions                        bool                  `json:"can_be_used_with_other_promotions,omitempty" url:"can_be_used_with_other_promotions,omitempty"`
	CurrencyCode                                        string                `json:"currency_code,omitempty" url:"currency_code,omitempty"`
	Notifications                                       []NotificationParams  `json:"notifications,omitempty" url:"notifications,omitempty"`
	ShippingAddress                                     ShippingAddressParams `json:"shipping_address,omitempty" url:"shipping_address,omitempty"`
	Schedule                                            ScheduleParams        `json:"schedule,omitempty" url:"schedule,omitempty"`
	CouponOverridesAutomaticWhenOfferingHigherDiscounts bool                  `json:"coupon_overrides_automatic_when_offering_higher_discounts,omitempty" url:"coupon_overrides_automatic_when_offering_higher_discounts,omitempty"`
}

type CustomerParams struct {
	GroupIDs          []int         `json:"group_ids,omitempty" url:"group_ids,omitempty"`
	MinimumOrderCount int           `json:"minimum_order_count,omitempty" url:"minimum_order_count,omitempty"`
	ExcludedGroupIDs  []int         `json:"excluded_group_ids,omitempty" url:"excluded_group_ids,omitempty"`
	Segments          SegmentParams `json:"segments,omitempty" url:"segments,omitempty"`
}

type SegmentParams struct {
	ID []string `json:"id,omitempty" url:"id,omitempty"`
}

type RuleParams struct {
	Action    string          `json:"action,omitempty" url:"action,omitempty"`
	ApplyOnce bool            `json:"apply_once,omitempty" url:"apply_once,omitempty"`
	Stop      bool            `json:"stop,omitempty" url:"stop,omitempty"`
	CartValue CartValueParams `json:"cart_value,omitempty" url:"cart_value,omitempty"`
	Condition ConditionParams `json:"condition,omitempty" url:"condition,omitempty"`
}

type CartValueParams struct {
	Discount DiscountParams `json:"discount,omitempty" url:"discount,omitempty"`
}

type DiscountParams struct {
	FixedAmount string `json:"fixed_amount,omitempty" url:"fixed_amount,omitempty"`
}

type ConditionParams struct {
	Cart CartParams `json:"cart,omitempty" url:"cart,omitempty"`
}

type CartParams struct {
	Items           CartItemParams `json:"items,omitempty" url:"items,omitempty"`
	MinimumSpend    string         `json:"minimum_spend,omitempty" url:"minimum_spend,omitempty"`
	MinimumQuantity int            `json:"minimum_quantity,omitempty" url:"minimum_quantity,omitempty"`
}

type CartItemParams struct {
	Brands []int `json:"brands,omitempty" url:"brands,omitempty"`
}

type NotificationParams struct {
	Content   string   `json:"content,omitempty" url:"content,omitempty"`
	Type      string   `json:"type,omitempty" url:"type,omitempty"`
	Locations []string `json:"locations,omitempty" url:"locations,omitempty"`
}

type ShippingAddressParams struct {
	Countries []CountryParams `json:"countries,omitempty" url:"countries,omitempty"`
}

type CountryParams struct {
	ISO2CountryCode string `json:"iso2_country_code,omitempty" url:"iso2_country_code,omitempty"`
}

type ScheduleParams struct {
	WeekFrequency  int      `json:"week_frequency,omitempty" url:"week_frequency,omitempty"`
	WeekDays       []string `json:"week_days,omitempty" url:"week_days,omitempty"`
	DailyStartTime string   `json:"daily_start_time,omitempty" url:"daily_start_time,omitempty"`
	DailyEndTime   string   `json:"daily_end_time,omitempty" url:"daily_end_time,omitempty"`
}

func (c *Client) UpdatePromotion(id int, params PromotionUpdateParams) (Promotion, error) {
	type Response struct {
		Data Promotion `json:"data"`
		Meta MetaData  `json:"meta"`
	}
	var response Response

	path := c.constructURL("promotions", strconv.Itoa(id))

	err := c.Put(path, params, &response)
	if err != nil {
		return response.Data, err
	}

	return response.Data, nil
}
