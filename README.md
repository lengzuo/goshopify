# goshopify

Open source Shopify Golang library. I will be very happy and welcome if you would like to provide your contribution to this library.  

**Note**: This library only supports below resource right now:
- [x] Access Scopes
- [x] Access Token
- [x] Address
- [x] Application Charge
- [x] Application Credit
- [x] Customer
- [x] Discount Code 
- [x] Draft Order
- [x] Event
- [x] Inventory Item 
- [x] Metafield
- [x] Order
- [x] Price Rule 
- [x] Product
- [x] Product Image
- [x] Recurring Application Charge
- [x] Report
- [x] Script Tag
- [x] Storefront Access Token
- [x] Theme
- [x] Variant 

## Install

```console
$ go get github.com/lengzuo/goshopify
```

## Use

```go
import "github.com/lengzuo/goshopify"
```

This gives you access to the `goshopify` package.

#### Oauth

If you don't have an access token yet, you can obtain one with the oauth flow.
Something like this will work:

```go
// Create an app to init OAuth.
app := goshopify.App{
    Credentials: common.Credentials{
        APIKey:    "abcd",
        APISecret: "abcd_1234",
    },
    RedirectURL: "https://authorized-redirect-url.com",
    Scope:       "read_products", 
}

// Perform OAuth Authentication for your app and redirect client to the authorizedURL 
func ShopOAuth(w http.ResponseWriter, r *http.Request) {
    shopName := r.URL.Query().Get("shop")
    state := "nonce"
    authUrl := app.AuthorizeURL(shopName, state)
    http.Redirect(w, r, authUrl, http.StatusFound)
}

// Retrieve a permanent access token from OAuth callback
func OAuthCallback(w http.ResponseWriter, r *http.Request) {
    // Validate callback URL 
    if ok, _ := app.VerifyAuthorizationURL(r.URL); !ok {
        http.Error(w, "invalid signature", http.StatusUnauthorized)
        return
    }

    query := r.URL.Query()
    shopName := query.Get("shop")
    code := query.Get("code")
    
	// Init /admin/oauth shopify client with your shopName
    adminOAuthClient := app.AdminClient(shopName)

    ctx := context.TODO()
	// Request admin access token from shopify
    resp, err := app.GetAccessToken(shopName, code)
    if err != nil {
        http.Error(w, fmt.Sprintf("error while calling app.GetAccessToken: %s", err), http.StatusInternalServerError)
		return
	}
	
    // Do something with the token, like store it in a DB.
	fmt.Printf("token is: %s", resp.AccessToken)

	// Access admin/oauth get access scopes API 
    _, err = adminOAuthClient.GetAccessScopes(ctx)
    if err != nil {
        http.Error(w, fmt.Sprintf("err oauthClient.GetAccessScopes(ctx) is : %s", err), http.StatusInternalServerError)
        return
    }
}
```

#### API calls with the access token
With a permanent access token, you can make API calls like this:

```go
    ctx := context.TODO()
    client := adminapi.New("ShopName",
        adminapi.WithToken("Token"),
        adminapi.WithVersion("2022-01"),
    )

    productsResp, err := client.Product.Get(ctx, dto.GetProductRequest{})
    if err != nil {
        fmt.Printf("err client.GetProducts is : %s", err)
        return
    }

    fmt.Printf("Get products response is : %+v", productsResp.Products)

    productCount, err := client.Product.Count(ctx, dto.CountProductRequest{})
    if err != nil {
        fmt.Printf("err client.Product is : %s", err)
        return
    }

    fmt.Printf("Get product count response is : %d", productCount.Count)
```

#### Private App Auth
Private Shopify apps use basic authentication. 

```go
// Create an app somewhere.
ctx := context.TODO()
client := adminapi.New("ShopName",
    adminapi.WithCredentials(common.Credentials{
        Password:  "password",
        APIKey:    "APIKey",
	}),
)

// Fetch the products.
productsResp, err := client.Product.Get(ctx, dto.GetProductRequest{}))
```
#### WithVersion
Read more details on the [Shopify API Versioning](https://shopify.dev/concepts/about-apis/versioning)
to understand the format and release schedules. You can use `WithVersion` to specify a specific version 
of the API. If you do not use this option you will be defaulted to the oldest stable API.

```go
client := adminapi.New("ShopName",
    adminapi.WithToken("Token"),
    adminapi.WithVersion("2022-01"),
)
```

#### Define you own shopify resources API 

Not all endpoints are implemented right now. In those case, feel free to
implement them and make a PR, or you can create your own struct for the data
and use `NewRequest` with the API client. This is how the existing endpoints
are implemented.

For example, let's say you want to fetch webhooks. There's a helper function
`Get` specifically for fetching stuff so this will work:

```go
// Declare a model for the webhook
type Webhook struct {
    ID int         `json:"id"`
    Address string `json:"address"`
}

// Declare a model for the resource root.
type WebhooksResource struct {
    Webhooks []Webhook `json:"webhooks"`
}

client := adminapi.New("ShopName",
    adminapi.WithToken("Token"),
)
ctx := context.TODO()
path := "admin/webhooks.json"
resp := new(WebhooksResource)
err := client.Call(ctx, http.MethodGet, path, nil, resp)
if err != nil {
    fmt.Printf("error is: %s\n", err)
    return
}

bytes, _ := json.Marshal(resp)
fmt.Printf("resp: %s", bytes)
```

