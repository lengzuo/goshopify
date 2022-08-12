package dto

import (
	"github.com/shopspring/decimal"
	"time"
)

type CountryHarmonizedSystemCodes struct {
	HarmonizedSystemCode string `json:"harmonized_system_code,omitempty"`
	CountryCode          string `json:"country_code,omitempty"`
}

type InventoryItem struct {
	Cost                         *decimal.Decimal               `json:"cost,omitempty"`
	CountryCodeOfOrigin          string                         `json:"country_code_of_origin,omitempty"`
	CountryHarmonizedSystemCodes []CountryHarmonizedSystemCodes `json:"country_harmonized_system_codes,omitempty"`
	CreatedAt                    *time.Time                     `json:"created_at,omitempty"`
	HarmonizedSystemCode         int                            `json:"harmonized_system_code,omitempty"`
	ID                           int64                          `json:"id,omitempty"`
	ProvinceCodeOfOrigin         string                         `json:"province_code_of_origin,omitempty"`
	Sku                          string                         `json:"sku,omitempty"`
	Tracked                      bool                           `json:"tracked,omitempty"`
	UpdatedAt                    *time.Time                     `json:"updated_at,omitempty"`
	RequiresShipping             bool                           `json:"requires_shipping,omitempty"`
}

type InventoryItemCollection struct {
	InventoryItem InventoryItem `json:"inventory_item"`
}

type GetInventoryItemRequest struct {
	IDs   string `url:"ids,omitempty"`
	Limit string `url:"limit,omitempty"`
}

type GetInventoryItemResponse struct {
	InventoryItems []InventoryItem `json:"inventory_items"`
}
