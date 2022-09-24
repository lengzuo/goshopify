package adminapi

import (
	"context"
	"github.com/lengzuo/goshopify/adminapi/discountcode"
	"github.com/lengzuo/goshopify/adminapi/event"
	"github.com/lengzuo/goshopify/adminapi/inventoryitem"
	"github.com/lengzuo/goshopify/adminapi/pricerule"
	"github.com/lengzuo/goshopify/adminapi/recurringapplicationcharge"
	"net/http"
	"net/url"
	"path"

	"github.com/lengzuo/goshopify/adminapi/address"
	"github.com/lengzuo/goshopify/adminapi/applicationcharge"
	"github.com/lengzuo/goshopify/adminapi/applicationcredit"
	"github.com/lengzuo/goshopify/adminapi/common/dto"
	"github.com/lengzuo/goshopify/adminapi/customer"
	"github.com/lengzuo/goshopify/adminapi/draftorder"
	"github.com/lengzuo/goshopify/adminapi/image"
	"github.com/lengzuo/goshopify/adminapi/metafield"
	"github.com/lengzuo/goshopify/adminapi/order"
	"github.com/lengzuo/goshopify/adminapi/product"
	"github.com/lengzuo/goshopify/adminapi/report"
	"github.com/lengzuo/goshopify/adminapi/scripttag"
	"github.com/lengzuo/goshopify/adminapi/storefrontacesstoken"
	"github.com/lengzuo/goshopify/adminapi/theme"
	"github.com/lengzuo/goshopify/adminapi/variant"
	"github.com/lengzuo/goshopify/adminapi/webhook"
	"github.com/lengzuo/goshopify/common"
)

const (
	defaultAPIPath    = "admin/api"
	defaultApiVersion = "2021-10"
)

// impl manages communication with the Shopify API.
type impl struct {
	// client used to communicate with the Shopify API.
	client *http.Client

	// baseURL for API requests.
	baseURL url.URL

	// version you're currently using of the api, defaults to "stable"
	apiVersion string

	// token used to access shopify APIs from token
	token string
	// credentials used to access shopify APIs from credentials
	credentials common.Credentials

	// All APIs
	Address                    address.API
	ApplicationCharge          applicationcharge.API
	ApplicationCredit          applicationcredit.API
	Customer                   customer.API
	DiscountCode               discountcode.API
	DraftOrder                 draftorder.API
	Event                      event.API
	InventoryItem              inventoryitem.API
	Metafield                  metafield.API
	Order                      order.API
	PriceRule                  pricerule.API
	Product                    product.API
	ProductImage               image.API
	RecurringApplicationCharge recurringapplicationcharge.API
	Report                     report.API
	ScriptTag                  scripttag.API
	StorefrontAccessToken      storefrontacesstoken.API
	Theme                      theme.API
	Variant                    variant.API
	Webhook                    webhook.API
}

func setupAPI(c *impl) {
	c.Address = address.New(c.Call)
	c.ApplicationCharge = applicationcharge.New(c.Call)
	c.ApplicationCredit = applicationcredit.New(c.Call)
	c.Customer = customer.New(c.Call)
	c.DiscountCode = discountcode.New(c.Call)
	c.DraftOrder = draftorder.New(c.Call)
	c.Event = event.New(c.Call)
	c.InventoryItem = inventoryitem.New(c.Call)
	c.Metafield = metafield.New(c.Call)
	c.Order = order.New(c.Call)
	c.PriceRule = pricerule.New(c.Call)
	c.Product = product.New(c.Call)
	c.ProductImage = image.New(c.Call)
	c.RecurringApplicationCharge = recurringapplicationcharge.New(c.Call)
	c.Report = report.New(c.Call)
	c.ScriptTag = scripttag.New(c.Call)
	c.StorefrontAccessToken = storefrontacesstoken.New(c.Call)
	c.Theme = theme.New(c.Call)
	c.Variant = variant.New(c.Call)
	c.Webhook = webhook.New(c.Call)
}

func New(shopName string, opts ...Option) *impl {
	baseURL, _ := url.Parse(common.ShopBaseUrl(shopName))
	c := &impl{
		baseURL:    *baseURL,
		client:     common.DefaultHTTPClient(),
		apiVersion: defaultApiVersion,
	}

	// client custom configuration
	for _, opt := range opts {
		opt(c)
	}

	c.baseURL.Path = path.Join(c.baseURL.Path, defaultAPIPath, c.apiVersion)

	setupAPI(c)

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
