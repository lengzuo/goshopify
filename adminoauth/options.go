package adminoauth

import (
	"net/http"

	"github.com/lengzuo/goshopify/common"
)

type Option func(c *impl)

// WithToken is a function to set after u got the shopify token
func WithToken(token string) Option {
	return func(c *impl) {
		c.token = token
	}
}

// WithCredentials is a function to set shopify credentials
func WithCredentials(credentials common.Credentials) Option {
	return func(c *impl) {
		c.credentials = credentials
	}
}

// WithClient is a function to override shopify credentials
func WithClient(httpClient *http.Client) Option {
	return func(c *impl) {
		c.client = httpClient
	}
}
