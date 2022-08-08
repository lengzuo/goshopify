package dto

import (
	"time"

	"github.com/shopspring/decimal"
)

type Refund struct {
	CreatedAt        *time.Time         `json:"created_at,omitempty"`
	Duties           Duties             `json:"duties,omitempty"`
	ID               int64              `json:"id,omitempty"`
	Note             string             `json:"note,omitempty"`
	OrderAdjustments []OrderAdjustments `json:"order_adjustments,omitempty"`
	ProcessedAt      *time.Time         `json:"processed_at,omitempty"`
	RefundDuties     []RefundDuties     `json:"refund_duties,omitempty"`
	RefundLineItems  []RefundLineItem   `json:"refund_line_items,omitempty"`
	Restock          bool               `json:"restock,omitempty"`
	Transactions     []Transaction      `json:"transactions,omitempty"`
	UserID           int64              `json:"user_id,omitempty"`
}

type Duty struct {
	DutyID    int64     `json:"duty_id,omitempty"`
	AmountSet AmountSet `json:"amount_set,omitempty"`
}

type Duties struct {
	Duties []Duty `json:"duties,omitempty"`
}

type OrderAdjustments struct {
	ID           int64            `json:"id,omitempty"`
	OrderID      int64            `json:"order_id,omitempty"`
	RefundID     int64            `json:"refund_id,omitempty"`
	Amount       *decimal.Decimal `json:"amount,omitempty"`
	TaxAmount    *decimal.Decimal `json:"tax_amount,omitempty"`
	Kind         string           `json:"kind,omitempty"`
	Reason       string           `json:"reason,omitempty"`
	AmountSet    AmountSet        `json:"amount_set,omitempty"`
	TaxAmountSet AmountSet        `json:"tax_amount_set,omitempty"`
}

type RefundDuties struct {
	DutyID     int    `json:"duty_id,omitempty"`
	RefundType string `json:"refund_type,omitempty"`
}

type RefundLineItem struct {
	ID          int64            `json:"id,omitempty"`
	Quantity    int              `json:"quantity,omitempty"`
	LineItemID  int64            `json:"line_item_id,omitempty"`
	LineItem    LineItem         `json:"line_item"`
	Subtotal    *decimal.Decimal `json:"subtotal,omitempty"`
	TotalTax    *decimal.Decimal `json:"total_tax,omitempty"`
	LocationID  int              `json:"location_id,omitempty"`
	RestockType string           `json:"restock_type,omitempty"`
	SubtotalSet AmountSet        `json:"subtotal_set,omitempty"`
	TotalTaxSet AmountSet        `json:"total_tax_set,omitempty"`
}
