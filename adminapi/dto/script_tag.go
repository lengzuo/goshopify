package dto

import "time"

type ScriptTag struct {
	CreatedAt    *time.Time `json:"created_at,omitempty"`
	Event        string     `json:"event,omitempty"`
	ID           int64      `json:"id,omitempty"`
	Src          string     `json:"src,omitempty"`
	DisplayScope string     `json:"display_scope,omitempty"`
	Cache        bool       `json:"cache,omitempty"`
	UpdatedAt    *time.Time `json:"updated_at,omitempty"`
}

type ScriptTagCollection struct {
	ScriptTag ScriptTag `json:"script_tag"`
}

type GetScriptTagRequest struct {
	CreatedAtMax *time.Time `url:"created_at_max"`
	CreatedAtMin *time.Time `url:"created_at_min"`
	Fields       string     `url:"fields"`
	Limit        string     `url:"limit"`
	SinceID      int64      `url:"since_id"`
	Src          string     `url:"src"`
	UpdatedAtMax *time.Time `url:"updated_at_max"`
	UpdatedAtMin *time.Time `url:"updated_at_min"`
}

type GetScriptTagResponse struct {
	ScriptTags []ScriptTag `json:"script_tags"`
}
