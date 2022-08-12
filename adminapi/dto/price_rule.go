package dto

import (
	"github.com/shopspring/decimal"
	"time"
)

type PrerequisiteSubtotalRange struct {
	GreaterThanOrEqualTo string `json:"greater_than_or_equal_to,omitempty"`
}

type PrerequisiteQuantityRange struct {
	GreaterThanOrEqualTo int `json:"greater_than_or_equal_to,omitempty"`
}

type PrerequisiteShippingPriceRange struct {
	LessThanOrEqualTo string `json:"less_than_or_equal_to,omitempty"`
}

type PrerequisiteToEntitlementQuantityRatio struct {
	PrerequisiteQuantity int `json:"prerequisite_quantity,omitempty"`
	EntitledQuantity     int `json:"entitled_quantity,omitempty"`
}

type PriceRule struct {
	ID                                     int64                                  `json:"id,omitempty"`
	Title                                  string                                 `json:"title,omitempty"`
	ValueType                              string                                 `json:"value_type,omitempty"`
	Value                                  *decimal.Decimal                       `json:"value,omitempty"`
	CustomerSelection                      string                                 `json:"customer_selection,omitempty"`
	TargetType                             string                                 `json:"target_type,omitempty"`
	TargetSelection                        string                                 `json:"target_selection,omitempty"`
	AllocationMethod                       string                                 `json:"allocation_method,omitempty"`
	AllocationLimit                        string                                 `json:"allocation_limit,omitempty"`
	OncePerCustomer                        bool                                   `json:"once_per_customer,omitempty"`
	UsageLimit                             int                                    `json:"usage_limit,omitempty"`
	StartsAt                               *time.Time                             `json:"starts_at,omitempty"`
	EndsAt                                 *time.Time                             `json:"ends_at,omitempty"`
	CreatedAt                              *time.Time                             `json:"created_at,omitempty"`
	UpdatedAt                              *time.Time                             `json:"updated_at,omitempty"`
	EntitledProductIds                     []int64                                `json:"entitled_product_ids,omitempty"`
	EntitledVariantIds                     []int64                                `json:"entitled_variant_ids,omitempty"`
	EntitledCollectionIds                  []int64                                `json:"entitled_collection_ids,omitempty"`
	EntitledCountryIds                     []int64                                `json:"entitled_country_ids,omitempty"`
	PrerequisiteProductIds                 []int64                                `json:"prerequisite_product_ids,omitempty"`
	PrerequisiteVariantIds                 []int64                                `json:"prerequisite_variant_ids,omitempty"`
	PrerequisiteCollectionIds              []int64                                `json:"prerequisite_collection_ids,omitempty"`
	PrerequisiteSavedSearchIds             []int64                                `json:"prerequisite_saved_search_ids,omitempty"`
	PrerequisiteCustomerIds                []int64                                `json:"prerequisite_customer_ids,omitempty"`
	PrerequisiteSubtotalRange              PrerequisiteSubtotalRange              `json:"prerequisite_subtotal_range,omitempty"`
	PrerequisiteQuantityRange              PrerequisiteQuantityRange              `json:"prerequisite_quantity_range,omitempty"`
	PrerequisiteShippingPriceRange         PrerequisiteShippingPriceRange         `json:"prerequisite_shipping_price_range,omitempty"`
	PrerequisiteToEntitlementQuantityRatio PrerequisiteToEntitlementQuantityRatio `json:"prerequisite_to_entitlement_quantity_ratio,omitempty"`
}

type PriceRuleCollection struct {
	PriceRule PriceRule `json:"price_rule"`
}

type GetPriceRuleRequest struct {
	CreatedAtMax *time.Time `url:"created_at_max,omitempty"`
	CreatedAtMin *time.Time `url:"created_at_min,omitempty"`
	EndsAtMax    *time.Time `url:"ends_at_max,omitempty"`
	EndsAtMin    *time.Time `url:"ends_at_min,omitempty"`
	Limit        string     `url:"limit,omitempty"`
	SinceID      int64      `url:"since_id,omitempty"`
	StartsAtMax  *time.Time `url:"starts_at_max,omitempty"`
	StartsAtMin  *time.Time `url:"starts_at_min,omitempty"`
	TimesUsed    bool       `url:"times_used"`
	UpdatedAtMax *time.Time `url:"updated_at_max,omitempty"`
	UpdatedAtMin *time.Time `url:"updated_at_min,omitempty"`
}

type GetPriceRuleResponse struct {
	PriceRules []PriceRule `json:"price_rules"`
}
