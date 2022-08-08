package adminoauth

import (
	"context"
	"net/http"
	"net/url"
	"path"

	"github.com/lengzuo/goshopify/adminapi/common/dto"
	"github.com/lengzuo/goshopify/common"
)

const (
	defaultAPIPath = "/admin/oauth/"
)

// impl manages communication with the Shopify API.
type impl struct {
	// client used to communicate with the Shopify API.
	client *http.Client

	// baseURL for API requests.
	baseURL url.URL

	// token used to access shopify APIs from token
	token string
	// credentials used to access shopify APIs from credentials
	credentials common.Credentials

	// Setup all APIs
	accessScopesEndpoint common.Endpoint
	accessTokenEndpoint  common.Endpoint
}

func New(shopName string, opts ...Option) *impl {
	baseURL, _ := url.Parse(common.ShopBaseUrl(shopName))
	c := &impl{
		client:               common.DefaultHTTPClient(),
		baseURL:              *baseURL,
		accessScopesEndpoint: setUpAccessScopesEndpoint(),
		accessTokenEndpoint:  setUpAccessTokenEndpoint(),
	}
	// client custom configuration
	for _, opt := range opts {
		opt(c)
	}

	c.baseURL.Path = path.Join(c.baseURL.Path, defaultAPIPath)

	return c
}

func (c *impl) Call(ctx context.Context, method, path string, req, resp interface{}) (err error) {
	return common.Send(ctx, c, method, path, req, resp, func(r *http.Request) {
		if c.token != "" {
			r.Header.Add(common.HeaderXShopifyAccessToken, c.token)
		} else if c.credentials.Password != "" {
			r.SetBasicAuth(c.credentials.APIKey, c.credentials.Password)
		}
	})
}

func (c *impl) Client() *http.Client {
	return c.client
}

func (c *impl) ErrorDTO() error {
	return new(dto.ErrResp)
}

func (c *impl) BaseURL() url.URL {
	return c.baseURL
}
