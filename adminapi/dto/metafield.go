package dto

import "time"

type Metafield struct {
	CreatedAt     *time.Time `json:"created_at,omitempty"`
	Description   string     `json:"description,omitempty"`
	ID            int64      `json:"id,omitempty"`
	Key           string     `json:"key,omitempty"`
	Namespace     string     `json:"namespace,omitempty"`
	OwnerID       int        `json:"owner_id,omitempty"`
	OwnerResource string     `json:"owner_resource,omitempty"`
	UpdatedAt     *time.Time `json:"updated_at,omitempty"`
	Value         int        `json:"value,omitempty"`
	Type          string     `json:"type,omitempty"`
}

type MetafieldCollection struct {
	Metafield Metafield `json:"metafield"`
}

type GetMetafieldRequest struct {
	CreatedAtMax  *time.Time `url:"created_at_max"`
	CreatedAtMin  *time.Time `url:"created_at_min"`
	Fields        string     `url:"fields"`
	Key           string     `url:"key"`
	Namespace     string     `url:"namespace"`
	SinceID       int64      `url:"since_id"`
	Type          string     `url:"type"`
	Title         string     `url:"title"`
	UpdatedAtMax  *time.Time `url:"updated_at_max"`
	UpdatedAtMin  *time.Time `url:"updated_at_min"`
	OwnerID       string     `url:"metafield[owner_id]"`
	OwnerResource string     `url:"metafield[owner_resource]"`
}

type GetMetafieldResponse struct {
	Metafields []Metafield `json:"metafields"`
}

type GetMetafieldOneRequest struct {
	Fields string `url:"fields"`
}
