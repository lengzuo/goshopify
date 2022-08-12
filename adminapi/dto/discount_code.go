package dto

import "time"

type DiscountCode struct {
	ID          int64      `json:"id,omitempty"`
	PriceRuleID int64      `json:"price_rule_id,omitempty"`
	Code        string     `json:"code,omitempty"`
	UsageCount  int        `json:"usage_count,omitempty"`
	CreatedAt   *time.Time `json:"created_at,omitempty"`
	UpdatedAt   *time.Time `json:"updated_at,omitempty"`
}

type DiscountCodeCollection struct {
	DiscountCode DiscountCode `json:"discount_code"`
}

type PriceRuleDiscountCode struct {
	Code string `json:"code"`
}

// PriceRuleDiscountCodeCollection use in Batch API request
type PriceRuleDiscountCodeCollection struct {
	DiscountCodes []PriceRuleDiscountCode `json:"discount_codes"`
}

type PriceRuleDiscountCodeCreation struct {
	ID            int64      `json:"id"`
	PriceRuleID   int64      `json:"price_rule_id"`
	StartedAt     *time.Time `json:"started_at"`
	CompletedAt   *time.Time `json:"completed_at"`
	CreatedAt     *time.Time `json:"created_at"`
	UpdatedAt     *time.Time `json:"updated_at"`
	Status        string     `json:"status"`
	CodesCount    int        `json:"codes_count"`
	ImportedCount int        `json:"imported_count"`
	FailedCount   int        `json:"failed_count"`
	Logs          []string   `json:"logs"`
}

// CreatePriceRuleDiscountCodeResponse use in Batch API response
type CreatePriceRuleDiscountCodeResponse struct {
	DiscountCodeCreation PriceRuleDiscountCodeCreation `json:"discount_code_creation"`
}

// CreateDiscountCodeRequest use for create single discount code request
type CreateDiscountCodeRequest struct {
	DiscountCode PriceRuleDiscountCode `json:"discount_code"`
}

type DiscountCodeCountRequest struct {
	TimesUsed    bool `url:"times_used"`
	TimesUsedMax int  `url:"times_used_max,omitempty"`
	TimesUsedMin int  `url:"times_used_min,omitempty"`
}

type DiscountLookCountRequest struct {
	Code bool `url:"code"`
}

type GetDiscountCodeResponse struct {
	DiscountCodes []DiscountCode `json:"discount_codes"`
}
