package dto

import (
	"time"
)

type SmsMarketingConsent struct {
	State                string     `json:"state,omitempty"`
	OptInLevel           string     `json:"opt_in_level,omitempty"`
	ConsentUpdatedAt     *time.Time `json:"consent_updated_at,omitempty"`
	ConsentCollectedFrom string     `json:"consent_collected_from,omitempty"`
}

type Customer struct {
	AcceptsMarketing          bool                `json:"accepts_marketing,omitempty"`
	AcceptsMarketingUpdatedAt *time.Time          `json:"accepts_marketing_updated_at,omitempty"`
	Addresses                 []Address           `json:"addresses,omitempty"`
	Currency                  string              `json:"currency,omitempty"`
	CreatedAt                 *time.Time          `json:"created_at,omitempty"`
	DefaultAddress            Address             `json:"default_address,omitempty"`
	Email                     string              `json:"email,omitempty"`
	FirstName                 string              `json:"first_name,omitempty"`
	ID                        int64               `json:"id,omitempty"`
	LastName                  string              `json:"last_name,omitempty"`
	LastOrderID               int64               `json:"last_order_id,omitempty"`
	LastOrderName             string              `json:"last_order_name,omitempty"`
	Metafield                 Metafield           `json:"metafield,omitempty"`
	MarketingOptInLevel       string              `json:"marketing_opt_in_level,omitempty"`
	MultipassIdentifier       string              `json:"multipass_identifier,omitempty"`
	Note                      string              `json:"note,omitempty"`
	OrdersCount               int                 `json:"orders_count,omitempty"`
	Password                  string              `json:"password,omitempty"`
	PasswordConfirmation      string              `json:"password_confirmation,omitempty"`
	Phone                     string              `json:"phone,omitempty"`
	SmsMarketingConsent       SmsMarketingConsent `json:"sms_marketing_consent,omitempty"`
	State                     string              `json:"state,omitempty"`
	Tags                      string              `json:"tags,omitempty"`
	TaxExempt                 bool                `json:"tax_exempt,omitempty"`
	TaxExemptions             []string            `json:"tax_exemptions,omitempty"`
	TotalSpent                string              `json:"total_spent,omitempty"`
	UpdatedAt                 *time.Time          `json:"updated_at,omitempty"`
	VerifiedEmail             bool                `json:"verified_email,omitempty"`
}

type CustomerCollection struct {
	Customer Customer `json:"customer"`
}

type CreateActivationURLResponse struct {
	AccountActivationURL string `json:"account_activation_url"`
}

type GetCustomerRequest struct {
	CreatedAtMax *time.Time `url:"created_at_max"`
	CreatedAtMin *time.Time `url:"created_at_min"`
	Fields       string     `url:"fields"`
	IDs          string     `url:"ids"`
	Limit        string     `url:"limit"`
	SinceID      int64      `url:"since_id"`
	UpdatedAtMax *time.Time `url:"updated_at_max"`
	UpdatedAtMin *time.Time `url:"updated_at_min"`
}

type GetCustomerResponse struct {
	Customers []Customer `json:"customers"`
}

type GetCustomerOrderRequest struct {
	Status string `url:"status,omitempty"`
}

type GetCustomerOrderResponse struct {
	Orders []Order `json:"orders"`
}

type SearchCustomerRequest struct {
	Query  string `url:"query"`
	Fields string `url:"fields"`
	Order  string `url:"order"`
	Limit  string `url:"limit"`
}

type SendCustomerInviteResponse struct {
	CustomerInvite Email `json:"customer_invite"`
}
