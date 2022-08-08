package dto

import "time"

type PresentmentPrices struct {
	Price          Price  `json:"price,omitempty"`
	CompareAtPrice string `json:"compare_at_price,omitempty"`
}

type Option struct {
	Option1 string `json:"option1"`
	Option2 string `json:"option2"`
	Option3 string `json:"option3"`
}

type Variant struct {
	Barcode                     string              `json:"barcode,omitempty"`
	CompareAtPrice              string              `json:"compare_at_price,omitempty"`
	CreatedAt                   *time.Time          `json:"created_at,omitempty"`
	FulfillmentService          string              `json:"fulfillment_service,omitempty"`
	Grams                       int                 `json:"grams,omitempty"`
	ID                          int64               `json:"id,omitempty"`
	ImageID                     int64               `json:"image_id,omitempty"`
	InventoryItemID             int64               `json:"inventory_item_id,omitempty"`
	InventoryManagement         string              `json:"inventory_management,omitempty"`
	InventoryPolicy             string              `json:"inventory_policy,omitempty"`
	InventoryQuantity           int                 `json:"inventory_quantity,omitempty"`
	OldInventoryQuantity        int                 `json:"old_inventory_quantity,omitempty"`
	InventoryQuantityAdjustment int                 `json:"inventory_quantity_adjustment,omitempty"`
	Option                      Option              `json:"option"`
	PresentmentPrices           []PresentmentPrices `json:"presentment_prices,omitempty"`
	Position                    int                 `json:"position,omitempty"`
	Price                       string              `json:"price,omitempty"`
	ProductID                   int64               `json:"product_id,omitempty"`
	RequiresShipping            bool                `json:"requires_shipping,omitempty"`
	Sku                         string              `json:"sku,omitempty"`
	Taxable                     bool                `json:"taxable,omitempty"`
	TaxCode                     string              `json:"tax_code,omitempty"`
	Title                       string              `json:"title,omitempty"`
	UpdatedAt                   string              `json:"updated_at,omitempty"`
	Weight                      float32             `json:"weight,omitempty"`
	WeightUnit                  string              `json:"weight_unit,omitempty"`
}

type VariantCollection struct {
	Variant Variant `json:"variant"`
}

type GetVariantRequest struct {
	Fields                string `url:"fields"`
	Limit                 int    `url:"limit"`
	PresentmentCurrencies string `url:"presentment_currencies"`
	SinceID               int64  `url:"since_id"`
}

type GetVariantResponse struct {
	Variants []Variant `json:"variants"`
}
