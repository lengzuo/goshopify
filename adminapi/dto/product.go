package dto

import (
	"time"
)

type Options struct {
	ID        int64    `json:"id,omitempty"`
	ProductID int64    `json:"product_id,omitempty"`
	Name      string   `json:"name,omitempty"`
	Position  int      `json:"position,omitempty"`
	Values    []string `json:"values,omitempty"`
}

type Product struct {
	ID                int64     `json:"id,omitempty"`
	Title             string    `json:"title,omitempty"`
	BodyHTML          string    `json:"body_html,omitempty"`
	Vendor            string    `json:"vendor,omitempty"`
	ProductType       string    `json:"product_type,omitempty"`
	CreatedAt         string    `json:"created_at,omitempty"`
	Handle            string    `json:"handle,omitempty"`
	UpdatedAt         string    `json:"updated_at,omitempty"`
	PublishedAt       string    `json:"published_at,omitempty"`
	TemplateSuffix    string    `json:"template_suffix,omitempty"`
	Status            string    `json:"status,omitempty"`
	PublishedScope    string    `json:"published_scope,omitempty"`
	Tags              string    `json:"tags,omitempty"`
	AdminGraphqlAPIID string    `json:"admin_graphql_api_id,omitempty"`
	Variants          []Variant `json:"variants,omitempty"`
	Options           []Options `json:"options,omitempty"`
	Images            []Image   `json:"images,omitempty"`
	Image             *Image    `json:"image,omitempty"`
}

type ProductCollection struct {
	Product Product `json:"product"`
}

type GetProductRequest struct {
	CollectionID          int64      `url:"collection_id,omitempty"`
	CreatedAtMax          *time.Time `url:"created_at_max,omitempty"`
	CreatedAtMin          *time.Time `url:"created_at_min,omitempty"`
	Fields                string     `url:"fields,omitempty"`
	Handle                string     `url:"handle,omitempty"`
	IDs                   string     `url:"ids,omitempty"`
	Limit                 string     `url:"limit,omitempty"`
	PresentmentCurrencies string     `url:"presentment_currencies,omitempty"`
	ProductType           string     `url:"product_type,omitempty"`
	PublishedAtMax        *time.Time `url:"published_at_max,omitempty"`
	PublishedAtMin        *time.Time `url:"published_at_min,omitempty"`
	PublishedStatus       string     `url:"published_status,omitempty"`
	SinceID               int64      `url:"since_id,omitempty"`
	Status                string     `url:"status,omitempty"`
	Title                 string     `url:"title,omitempty"`
	UpdatedAtMax          *time.Time `url:"updated_at_max,omitempty"`
	UpdatedAtMin          *time.Time `url:"updated_at_min,omitempty"`
	Vendor                string     `url:"vendor,omitempty"`
}

type GetProductResponse struct {
	Products []Product `json:"products"`
}

type CountProductRequest struct {
	CollectionID    int64      `url:"collection_id,omitempty"`
	CreatedAtMax    *time.Time `url:"created_at_max,omitempty"`
	CreatedAtMin    *time.Time `url:"created_at_min,omitempty"`
	ProductType     string     `url:"product_type,omitempty"`
	PublishedAtMax  *time.Time `url:"published_at_max,omitempty"`
	PublishedAtMin  *time.Time `url:"published_at_min,omitempty"`
	PublishedStatus string     `url:"published_status,omitempty"`
	UpdatedAtMax    *time.Time `url:"updated_at_max,omitempty"`
	UpdatedAtMin    *time.Time `url:"updated_at_min,omitempty"`
	Vendor          string     `url:"vendor,omitempty"`
}
