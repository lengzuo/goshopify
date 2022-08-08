package dto

import (
	"time"

	"github.com/shopspring/decimal"
)

type ShippingLines struct {
	Custom string           `json:"custom,omitempty"`
	Handle string           `json:"handle,omitempty"`
	Title  string           `json:"title,omitempty"`
	Price  *decimal.Decimal `json:"price,omitempty"`
}

type PaymentTerms struct {
	Amount           *decimal.Decimal   `json:"amount,omitempty"`
	Currency         string             `json:"currency,omitempty"`
	PaymentTermsName string             `json:"payment_terms_name,omitempty"`
	PaymentTermsType string             `json:"payment_terms_type,omitempty"`
	DueInDays        int                `json:"due_in_days,omitempty"`
	PaymentSchedules []PaymentSchedules `json:"payment_schedules,omitempty"`
}

type PaymentSchedules struct {
	Amount                *decimal.Decimal `json:"amount,omitempty"`
	Currency              string           `json:"currency,omitempty"`
	IssuedAt              *time.Time       `json:"issued_at,omitempty"`
	DueAt                 *time.Time       `json:"due_at,omitempty"`
	CompletedAt           *time.Time       `json:"completed_at,omitempty"`
	ExpectedPaymentMethod string           `json:"expected_payment_method,omitempty"`
}

type DraftOrder struct {
	ID              int64            `json:"id,omitempty"`
	OrderID         int64            `json:"order_id,omitempty"`
	Name            string           `json:"name,omitempty"`
	Customer        Customer         `json:"customer,omitempty"`
	ShippingAddress Address          `json:"shipping_address,omitempty"`
	BillingAddress  Address          `json:"billing_address,omitempty"`
	Note            string           `json:"note,omitempty"`
	NoteAttributes  []NoteAttribute  `json:"note_attribute,omitempty"`
	Email           string           `json:"email,omitempty"`
	Currency        string           `json:"currency,omitempty"`
	InvoiceSentAt   *time.Time       `json:"invoice_sent_at,omitempty"`
	InvoiceURL      string           `json:"invoice_url,omitempty"`
	LineItems       []LineItem       `json:"line_items,omitempty"`
	PaymentTerms    PaymentTerms     `json:"payment_terms,omitempty"`
	ShippingLine    *ShippingLines   `json:"shipping_line,omitempty"`
	SourceName      string           `json:"source_name"`
	Tags            string           `json:"tags,omitempty"`
	TaxExempt       bool             `json:"tax_exempt"`
	TaxExemptions   []string         `json:"tax_exemptions"`
	TaxLines        []TaxLine        `json:"tax_lines,omitempty"`
	AppliedDiscount *AppliedDiscount `json:"applied_discount,omitempty"`
	TaxesIncluded   bool             `json:"taxes_included,omitempty"`
	TotalTax        *decimal.Decimal `json:"total_tax,omitempty"`
	TotalPrice      *decimal.Decimal `json:"total_price,omitempty"`
	SubtotalPrice   *decimal.Decimal `json:"subtotal_price,omitempty"`
	CompletedAt     *time.Time       `json:"completed_at,omitempty"`
	CreatedAt       *time.Time       `json:"created_at,omitempty"`
	UpdatedAt       *time.Time       `json:"updated_at,omitempty"`
	Status          string           `json:"status,omitempty"`
}

type DraftOrderCollection struct {
	DraftOrder DraftOrder `json:"draft_order"`
}

type DraftOrderInvoice struct {
	DraftOrderInvoice Email `json:"draft_order_invoice"`
}

type GetDraftOrderRequest struct {
	Fields       string     `url:"fields,omitempty"`
	IDs          string     `url:"ids,omitempty"`
	Limit        string     `url:"limit,omitempty"`
	SinceID      int64      `url:"since_id,omitempty"`
	Status       string     `url:"status,omitempty"`
	UpdatedAtMax *time.Time `url:"updated_at_max,omitempty"`
	UpdatedAtMin *time.Time `url:"updated_at_min,omitempty"`
}

type GetDraftOrderResponse struct {
	DraftOrders []DraftOrder `json:"draft_orders"`
}

type CountDraftOrderRequest struct {
	SinceID      int64      `url:"since_id,omitempty"`
	Status       string     `url:"status,omitempty"`
	UpdatedAtMax *time.Time `url:"updated_at_max,omitempty"`
	UpdatedAtMin *time.Time `url:"updated_at_min,omitempty"`
}

type CompleteDraftOrderRequest struct {
	PaymentGatewayID string `json:"payment_gateway_id"`
	PaymentPending   *bool  `json:"payment_pending,omitempty"`
}
