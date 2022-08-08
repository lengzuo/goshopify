package dto

import "time"

type Theme struct {
	CreatedAt    *time.Time `json:"created_at,omitempty"`
	ID           int64      `json:"id,omitempty"`
	Name         string     `json:"name,omitempty"`
	Previewable  bool       `json:"previewable,omitempty"`
	Processing   bool       `json:"processing,omitempty"`
	Role         string     `json:"role,omitempty"`
	ThemeStoreID int        `json:"theme_store_id,omitempty"`
	UpdatedAt    *time.Time `json:"updated_at,omitempty"`
}

type ThemeCollection struct {
	Theme Theme `json:"theme"`
}

type GetThemeRequest struct {
	Limit int `url:"limit"`
}

type GetThemeResponse struct {
	Themes []Theme `json:"themes"`
}
