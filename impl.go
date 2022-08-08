package goshopify

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"net/url"

	"github.com/google/go-querystring/query"
	"github.com/lengzuo/goshopify/adminoauth"
	"github.com/lengzuo/goshopify/common"
)

const (
	authorizePath = "admin/oauth/authorize"
	qsHMAC        = "hmac"
	qsSignature   = "signature"
)

// App represents basic App settings such as Api key, secret, scope, and redirect url.
// See oauth.go for OAuth related helper functions.
type App struct {
	Credentials common.Credentials
	RedirectURL string
	Scope       string
}

func (a App) AuthorizeURL(shopName string, state string) string {
	shopUrl, _ := url.Parse(common.ShopBaseUrl(shopName))
	shopUrl.Path = authorizePath
	qs := Authorize{
		ClientID:    a.Credentials.APIKey,
		RedirectURI: a.RedirectURL,
		Scope:       a.Scope,
		State:       state,
	}
	q, _ := query.Values(qs)
	shopUrl.RawQuery = q.Encode()
	return shopUrl.String()
}

// VerifyMessage uses to verify message against a message HMAC
func (a App) VerifyMessage(message, messageMAC string) bool {
	mac := hmac.New(sha256.New, []byte(a.Credentials.APISecret))
	mac.Write([]byte(message))
	expectedMAC := mac.Sum(nil)

	actualMac, err := hex.DecodeString(messageMAC)
	if err != nil {
		return false
	}

	return hmac.Equal(actualMac, expectedMAC)
}

// VerifyAuthorizationURL use to check callback parameters from url.
func (a App) VerifyAuthorizationURL(u *url.URL) (bool, error) {
	q := u.Query()
	messageMAC := q.Get(qsHMAC)

	q.Del(qsHMAC)
	q.Del(qsSignature)

	message, err := url.QueryUnescape(q.Encode())

	return a.VerifyMessage(message, messageMAC), err
}

func (a App) AdminClient(shopName string) adminoauth.API {
	return adminoauth.New(shopName, adminoauth.WithCredentials(a.Credentials))
}
