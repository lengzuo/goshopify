package dto

import "time"

type Report struct {
	Category  string     `json:"category,omitempty"`
	ID        int64      `json:"id,omitempty"`
	Name      string     `json:"name,omitempty"`
	ShopifyQL string     `json:"shopify_ql,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}

type ReportCollection struct {
	Report Report `json:"report"`
}

type GetReportRequest struct {
	Fields       string     `url:"fields"`
	IDs          string     `url:"ids"`
	Limit        string     `url:"limit"`
	SinceID      int64      `url:"since_id"`
	UpdatedAtMax *time.Time `url:"updated_at_max"`
	UpdatedAtMin *time.Time `url:"updated_at_min"`
}

type GetReportResponse struct {
	Reports []Report `json:"reports"`
}

type UpdateReportRequest struct {
	Report Report `json:"report"`
}
