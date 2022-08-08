package dto

import "time"

type ApplicationCharge struct {
	ConfirmationURL string     `json:"confirmation_url,omitempty"`
	CreatedAt       *time.Time `json:"created_at,omitempty"`
	ID              int64      `json:"id,omitempty"`
	Name            string     `json:"name,omitempty"`
	Price           string     `json:"price,omitempty"`
	ReturnURL       string     `json:"return_url,omitempty"`
	Status          string     `json:"status,omitempty"`
	Test            bool       `json:"test,omitempty"`
	UpdatedAt       *time.Time `json:"updated_at,omitempty"`
}

type ApplicationChargeCollection struct {
	ApplicationCharge ApplicationCharge `json:"application_charge"`
}

type GetApplicationChargeRequest struct {
	Fields  string `json:"fields"`
	SinceID int64  `json:"since_id"`
}

type GetApplicationChargeResponse struct {
	ApplicationCharges []ApplicationCharge `json:"application_charges"`
}
