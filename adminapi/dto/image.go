package dto

import "time"

type Image struct {
	ID         int64      `json:"id,omitempty"`
	ProductID  int64      `json:"product_id,omitempty"`
	Position   int        `json:"position,omitempty"`
	CreatedAt  *time.Time `json:"created_at,omitempty"`
	UpdatedAt  *time.Time `json:"updated_at,omitempty"`
	Width      int        `json:"width,omitempty"`
	Height     int        `json:"height,omitempty"`
	Src        string     `json:"src,omitempty"`
	VariantIds []int64    `json:"variant_ids,omitempty"`
}

type ImageCollection struct {
	Image Image `json:"image"`
}

type GetImageRequest struct {
	Fields  string `url:"fields"`
	SinceID int64  `url:"since_id"`
}

type GetImageResponse struct {
	Images []Image `json:"images"`
}
