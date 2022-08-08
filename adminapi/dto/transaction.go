package dto

import (
	"time"

	"github.com/shopspring/decimal"
)

type PaymentDetails struct {
	AVSResultCode     string `json:"avs_result_code,omitempty"`
	CreditCardBin     string `json:"credit_card_bin,omitempty"`
	CVVResultCode     string `json:"cvv_result_code,omitempty"`
	CreditCardNumber  string `json:"credit_card_number,omitempty"`
	CreditCardCompany string `json:"credit_card_company,omitempty"`
}

type Transaction struct {
	ID                              int64                           `json:"id,omitempty"`
	OrderID                         int64                           `json:"order_id,omitempty"`
	Amount                          *decimal.Decimal                `json:"amount,omitempty"`
	Kind                            string                          `json:"kind,omitempty"`
	Gateway                         string                          `json:"gateway,omitempty"`
	Status                          string                          `json:"status,omitempty"`
	Message                         string                          `json:"message,omitempty"`
	CreatedAt                       *time.Time                      `json:"created_at,omitempty"`
	Test                            bool                            `json:"test,omitempty"`
	Authorization                   string                          `json:"authorization,omitempty"`
	AuthorizationExpiresAt          *time.Time                      `json:"authorization_expires_at,omitempty"`
	Currency                        string                          `json:"currency,omitempty"`
	LocationID                      int64                           `json:"location_id,omitempty"`
	UserID                          int64                           `json:"user_id,omitempty"`
	ParentID                        int64                           `json:"parent_id,omitempty"`
	DeviceID                        int64                           `json:"device_id,omitempty"`
	ErrorCode                       string                          `json:"error_code,omitempty"`
	ExtendedAuthorizationAttributes ExtendedAuthorizationAttributes `json:"extended_authorization_attributes,omitempty"`
	SourceName                      string                          `json:"source_name,omitempty"`
	Source                          string                          `json:"source,omitempty,omitempty"`
	PaymentDetails                  PaymentDetails                  `json:"payment_details,omitempty"`
	PaymentsRefundAttributes        PaymentsRefundAttributes        `json:"payments_refund_attributes,omitempty"`
	CurrencyExchangeAdjustment      CurrencyExchangeAdjustment      `json:"currency_exchange_adjustment,omitempty"`
}

type ExtendedAuthorizationAttributes struct {
	StandardAuthorizationExpiresAt string `json:"standard_authorization_expires_at,omitempty"`
	ExtendedAuthorizationExpiresAt string `json:"extended_authorization_expires_at,omitempty"`
}

type PaymentsRefundAttributes struct {
	Status                  string `json:"status,omitempty"`
	AcquirerReferenceNumber string `json:"acquirer_reference_number,omitempty"`
}

type CurrencyExchangeAdjustment struct {
	ID             int64  `json:"id,omitempty"`
	Adjustment     string `json:"adjustment,omitempty"`
	OriginalAmount string `json:"original_amount,omitempty"`
	FinalAmount    string `json:"final_amount,omitempty"`
	Currency       string `json:"currency,omitempty"`
}
