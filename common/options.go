package common

import (
	"net/http"
)

type Option func(r *http.Request)

// WithAdminAuthHeader optionally sets the api-version if the passed string is valid
func WithAdminAuthHeader(token string) Option {
	return func(r *http.Request) {
		r.Header.Set(HeaderXShopifyAccessToken, token)
	}
}

// WithCredentials is a function to set shopify credentials
func WithCredentials(credentials Credentials) Option {
	return func(r *http.Request) {
		r.SetBasicAuth(credentials.APIKey, credentials.Password)
	}
}
