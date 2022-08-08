package dto

import (
	"time"

	"github.com/shopspring/decimal"
)

type Order struct {
	ID                    int64            `json:"id,omitempty"`
	Name                  string           `json:"name,omitempty"`
	Email                 string           `json:"email,omitempty"`
	CreatedAt             *time.Time       `json:"created_at,omitempty"`
	UpdatedAt             *time.Time       `json:"updated_at,omitempty"`
	CancelledAt           *time.Time       `json:"cancelled_at,omitempty"`
	ClosedAt              *time.Time       `json:"closed_at,omitempty"`
	ProcessedAt           *time.Time       `json:"processed_at,omitempty"`
	Customer              Customer         `json:"customer"`
	BillingAddress        Address          `json:"billing_address"`
	ShippingAddress       Address          `json:"shipping_address"`
	Currency              string           `json:"currency,omitempty"`
	TotalPrice            *decimal.Decimal `json:"total_price,omitempty"`
	SubtotalPrice         *decimal.Decimal `json:"subtotal_price,omitempty"`
	TotalDiscounts        *decimal.Decimal `json:"total_discounts,omitempty"`
	TotalLineItemsPrice   *decimal.Decimal `json:"total_line_items_price,omitempty"`
	TaxesIncluded         bool             `json:"taxes_included,omitempty"`
	TotalTax              *decimal.Decimal `json:"total_tax,omitempty"`
	TaxLines              []TaxLine        `json:"tax_lines,omitempty"`
	TotalWeight           int              `json:"total_weight,omitempty"`
	FinancialStatus       string           `json:"financial_status,omitempty"`
	Fulfillments          []Fulfillment    `json:"fulfillments,omitempty"`
	FulfillmentStatus     string           `json:"fulfillment_status,omitempty"`
	Token                 string           `json:"token,omitempty"`
	CartToken             string           `json:"cart_token,omitempty"`
	Number                int              `json:"number,omitempty"`
	OrderNumber           int64            `json:"order_number,omitempty"`
	Note                  string           `json:"note,omitempty"`
	Test                  bool             `json:"test,omitempty"`
	BrowserIp             string           `json:"browser_ip,omitempty"`
	BuyerAcceptsMarketing bool             `json:"buyer_accepts_marketing,omitempty"`
	CancelReason          string           `json:"cancel_reason,omitempty"`
	NoteAttributes        []NoteAttribute  `json:"note_attributes,omitempty"`
	DiscountCodes         []DiscountCode   `json:"discount_codes,omitempty"`
	LineItems             []LineItem       `json:"line_items,omitempty"`
	ShippingLines         []ShippingLine   `json:"shipping_lines,omitempty"`
	Transactions          []Transaction    `json:"transactions,omitempty"`
	AppID                 int64            `json:"app_id,omitempty"`
	CustomerLocale        string           `json:"customer_locale,omitempty"`
	LandingSite           string           `json:"landing_site,omitempty"`
	ReferringSite         string           `json:"referring_site,omitempty"`
	SourceName            string           `json:"source_name,omitempty"`
	ClientDetails         ClientDetails    `json:"client_details"`
	Tags                  string           `json:"tags,omitempty"`
	LocationID            int64            `json:"location_id,omitempty"`
	PaymentGatewayNames   []string         `json:"payment_gateway_names,omitempty"`
	ProcessingMethod      string           `json:"processing_method,omitempty"`
	Refunds               []Refund         `json:"refunds,omitempty"`
	UserID                int64            `json:"user_id,omitempty"`
	OrderStatusUrl        string           `json:"order_status_url,omitempty"`
	Gateway               string           `json:"gateway,omitempty"`
	Confirmed             bool             `json:"confirmed,omitempty"`
	TotalPriceUSD         *decimal.Decimal `json:"total_price_usd,omitempty"`
	CheckoutToken         string           `json:"checkout_token,omitempty"`
	Reference             string           `json:"reference,omitempty"`
	SourceIdentifier      string           `json:"source_identifier,omitempty"`
	SourceURL             string           `json:"source_url,omitempty"`
	DeviceID              int64            `json:"device_id,omitempty"`
	Phone                 string           `json:"phone,omitempty"`
	LandingSiteRef        string           `json:"landing_site_ref,omitempty"`
	CheckoutID            int64            `json:"checkout_id,omitempty"`
	ContactEmail          string           `json:"contact_email,omitempty"`
	Metafields            []Metafield      `json:"metafields,omitempty"`
}

// Fulfillment represents a Shopify fulfillment.
type Fulfillment struct {
	ID              int64      `json:"id,omitempty"`
	OrderID         int64      `json:"order_id,omitempty"`
	LocationID      int64      `json:"location_id,omitempty"`
	Status          string     `json:"status,omitempty"`
	CreatedAt       *time.Time `json:"created_at,omitempty"`
	Service         string     `json:"service,omitempty"`
	UpdatedAt       *time.Time `json:"updated_at,omitempty"`
	TrackingCompany string     `json:"tracking_company,omitempty"`
	ShipmentStatus  string     `json:"shipment_status,omitempty"`
	TrackingNumber  string     `json:"tracking_number,omitempty"`
	TrackingNumbers []string   `json:"tracking_numbers,omitempty"`
	TrackingUrl     string     `json:"tracking_url,omitempty"`
	TrackingUrls    []string   `json:"tracking_urls,omitempty"`
	Receipt         Receipt    `json:"receipt,omitempty"`
	LineItems       []LineItem `json:"line_items,omitempty"`
	NotifyCustomer  bool       `json:"notify_customer"`
}

type Receipt struct {
	TestCase      bool   `json:"testcase,omitempty"`
	Authorization string `json:"authorization,omitempty"`
}

type ShippingLine struct {
	ID                            int64            `json:"id,omitempty"`
	Title                         string           `json:"title,omitempty"`
	Price                         *decimal.Decimal `json:"price,omitempty"`
	Code                          string           `json:"code,omitempty"`
	Source                        string           `json:"source,omitempty"`
	Phone                         string           `json:"phone,omitempty"`
	RequestedFulfillmentServiceID string           `json:"requested_fulfillment_service_id,omitempty"`
	DeliveryCategory              string           `json:"delivery_category,omitempty"`
	CarrierIdentifier             string           `json:"carrier_identifier,omitempty"`
	TaxLines                      []TaxLine        `json:"tax_lines,omitempty"`
}

type DiscountCode struct {
	Amount *decimal.Decimal `json:"amount,omitempty"`
	Code   string           `json:"code,omitempty"`
	Type   string           `json:"type,omitempty"`
}

type LineItem struct {
	ID                         int64                 `json:"id,omitempty"`
	ProductID                  int64                 `json:"product_id,omitempty"`
	VariantID                  int64                 `json:"variant_id,omitempty"`
	Quantity                   int                   `json:"quantity,omitempty"`
	Price                      *decimal.Decimal      `json:"price,omitempty"`
	TotalDiscount              *decimal.Decimal      `json:"total_discount,omitempty"`
	Title                      string                `json:"title,omitempty"`
	VariantTitle               string                `json:"variant_title,omitempty"`
	Name                       string                `json:"name,omitempty"`
	SKU                        string                `json:"sku,omitempty"`
	Vendor                     string                `json:"vendor,omitempty"`
	GiftCard                   bool                  `json:"gift_card,omitempty"`
	Taxable                    bool                  `json:"taxable,omitempty"`
	FulfillmentService         string                `json:"fulfillment_service,omitempty"`
	RequiresShipping           bool                  `json:"requires_shipping,omitempty"`
	VariantInventoryManagement string                `json:"variant_inventory_management,omitempty"`
	PreTaxPrice                *decimal.Decimal      `json:"pre_tax_price,omitempty"`
	Properties                 []NoteAttribute       `json:"properties,omitempty"`
	ProductExists              bool                  `json:"product_exists,omitempty"`
	FulfillableQuantity        int                   `json:"fulfillable_quantity,omitempty"`
	Grams                      int                   `json:"grams,omitempty"`
	FulfillmentStatus          string                `json:"fulfillment_status,omitempty"`
	TaxLines                   []TaxLine             `json:"tax_lines,omitempty"`
	OriginLocation             Address               `json:"origin_location,omitempty"`
	DestinationLocation        Address               `json:"destination_location,omitempty"`
	AppliedDiscount            AppliedDiscount       `json:"applied_discount,omitempty"`
	DiscountAllocations        []DiscountAllocations `json:"discount_allocations,omitempty"`
}

// AppliedDiscount is the discount applied to the line item or the draft order object.
type AppliedDiscount struct {
	Title       string `json:"applied_discount,omitempty"`
	Description string `json:"description,omitempty"`
	Value       string `json:"value,omitempty"`
	ValueType   string `json:"value_type,omitempty"`
	Amount      string `json:"amount,omitempty"`
}

type NoteAttribute struct {
	Name  string `json:"name,omitempty"`
	Value string `json:"value,omitempty"`
}

type TaxLine struct {
	Title string           `json:"title,omitempty"`
	Price *decimal.Decimal `json:"price,omitempty"`
	Rate  *decimal.Decimal `json:"rate,omitempty"`
}

type DiscountAllocations struct {
	Amount                   *decimal.Decimal `json:"amount,omitempty"`
	DiscountApplicationIndex int              `json:"discount_application_index,omitempty"`
	AmountSet                AmountSet        `json:"amount_set,omitempty"`
}

type AmountSet struct {
	ShopMoney        Price `json:"shop_money,omitempty"`
	PresentmentMoney Price `json:"presentment_money,omitempty"`
}

type ClientDetails struct {
	AcceptLanguage string `json:"accept_language,omitempty"`
	BrowserHeight  int    `json:"browser_height,omitempty"`
	BrowserIp      string `json:"browser_ip,omitempty"`
	BrowserWidth   int    `json:"browser_width,omitempty"`
	SessionHash    string `json:"session_hash,omitempty"`
	UserAgent      string `json:"user_agent,omitempty"`
}

type OrderCollection struct {
	Order Order `json:"order"`
}

type CancelOrderRequest struct {
	Amount   *decimal.Decimal `json:"amount,omitempty"`
	Currency string           `json:"currency,omitempty"`
	Email    bool             `json:"email,omitempty"`
	Reason   string           `json:"reason,omitempty"`
	Refund   Refund           `json:"refund"`
	Restock  bool             `json:"restock"`
}

type GetOrderRequest struct {
	AttributionAppID  string     `url:"attribution_app_id,omitempty"`
	CreatedAtMax      *time.Time `url:"created_at_max,omitempty"`
	CreatedAtMin      *time.Time `url:"created_at_min,omitempty"`
	Fields            string     `url:"fields,omitempty"`
	FinancialStatus   string     `url:"financial_status,omitempty"`
	FulfillmentStatus string     `url:"fulfillment_status,omitempty"`
	IDs               string     `url:"ids,omitempty"`
	Limit             string     `url:"limit,omitempty"`
	ProcessedAtMax    *time.Time `url:"processed_at_max,omitempty"`
	ProcessedAtMin    *time.Time `url:"processed_at_min,omitempty"`
	SinceID           int64      `url:"since_id,omitempty"`
}

type GetOrderResponse struct {
	Orders []Order `json:"orders"`
}

type CountOrderRequest struct {
	CreatedAtMax      *time.Time `url:"created_at_max,omitempty"`
	CreatedAtMin      *time.Time `url:"created_at_min,omitempty"`
	FinancialStatus   string     `url:"financial_status,omitempty"`
	FulfillmentStatus string     `url:"fulfillment_status,omitempty"`
	Status            string     `url:"status,omitempty"`
	UpdatedAtMax      *time.Time `url:"updated_at_max,omitempty"`
	UpdatedAtMin      *time.Time `url:"updated_at_min,omitempty"`
}
