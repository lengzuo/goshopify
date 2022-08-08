package dto

import "time"

type Address struct {
	ID           int64      `json:"id,omitempty"`
	CustomerID   int64      `json:"customer_id,omitempty"`
	CreatedAt    *time.Time `json:"created_at,omitempty"`
	FirstName    string     `json:"first_name,omitempty"`
	LastName     string     `json:"last_name,omitempty"`
	Company      string     `json:"company,omitempty"`
	Address1     string     `json:"address1,omitempty"`
	Address2     string     `json:"address2,omitempty"`
	City         string     `json:"city,omitempty"`
	Province     string     `json:"province,omitempty"`
	Country      string     `json:"country,omitempty"`
	Zip          string     `json:"zip,omitempty"`
	Phone        string     `json:"phone,omitempty"`
	ProvinceCode string     `json:"province_code,omitempty"`
	CountryCode  string     `json:"country_code,omitempty"`
	CountryName  string     `json:"country_name,omitempty"`
	Default      bool       `json:"default,omitempty"`
	UpdatedAt    *time.Time `json:"updated_at,omitempty"`
}

type AddressCollection struct {
	CustomerAddress Address `json:"customer_address"`
}

// BulkUpdateAddressRequest is a request for bulk api
type BulkUpdateAddressRequest struct {
	AddressIDs string `url:"address_ids[]"`
	Operation  string `url:"operation"`
}

type GetAddressRequest struct {
	Fields string `url:"fields"`
}

type GetAddressResponse struct {
	Addresses []AddressCollection `json:"addresses"`
}

type UpdateAddressRequest struct {
	Address Address `json:"address"`
}
