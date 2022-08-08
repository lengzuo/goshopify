package dto

import "time"

type StorefrontAccessToken struct {
	AccessToken       string     `json:"access_token,omitempty"`
	AccessScope       string     `json:"access_scope,omitempty"`
	CreatedAt         *time.Time `json:"created_at,omitempty"`
	ID                int64      `json:"id,omitempty"`
	AdminGraphqlAPIID string     `json:"admin_graphql_api_id,omitempty"`
	Title             string     `json:"title,omitempty"`
}

type StorefrontAccessTokenCollection struct {
	StorefrontAccessToken StorefrontAccessToken `json:"storefront_access_token"`
}

type GetStorefrontAccessTokenResponse struct {
	StorefrontAccessTokens []StorefrontAccessToken `json:"storefront_access_tokens"`
}
