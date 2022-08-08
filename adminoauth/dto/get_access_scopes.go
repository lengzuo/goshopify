package dto

type AccessScope struct {
	Handle string `json:"handle,omitempty"`
}

type GetAccessScopeResponse struct {
	AccessScopes []AccessScope `json:"access_scopes"`
}
