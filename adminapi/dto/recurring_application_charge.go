package dto

import "time"

type RecurringApplicationCharge struct {
	ActivatedOn     *time.Time `json:"activated_on,omitempty"`
	BillingOn       *time.Time `json:"billing_on,omitempty"`
	CancelledOn     *time.Time `json:"cancelled_on,omitempty"`
	CappedAmount    string     `json:"capped_amount,omitempty"`
	ConfirmationURL string     `json:"confirmation_url,omitempty"`
	CreatedAt       string     `json:"created_at,omitempty"`
	ID              int64      `json:"id,omitempty"`
	Name            string     `json:"name,omitempty"`
	Price           string     `json:"price,omitempty"`
	ReturnURL       string     `json:"return_url,omitempty"`
	Status          string     `json:"status,omitempty"`
	Terms           string     `json:"terms,omitempty"`
	Test            bool       `json:"test,omitempty"`
	TrialDays       int64      `json:"trial_days,omitempty"`
	TrialEndsOn     *time.Time `json:"trial_ends_on,omitempty"`
	UpdatedAt       time.Time  `json:"updated_at,omitempty"`
}

type RecurringApplicationChargeCollection struct {
	RecurringApplicationCharge RecurringApplicationCharge `json:"recurring_application_charge"`
}

type GetRecurringApplicationChargeRequest struct {
	Fields  string `json:"fields"`
	SinceID int64  `json:"since_id"`
}

type GetRecurringApplicationChargeResponse struct {
	RecurringApplicationCharges []RecurringApplicationCharge `json:"recurring_application_charges"`
}
