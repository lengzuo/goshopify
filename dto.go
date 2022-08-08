package goshopify

type Authorize struct {
	ClientID    string `url:"client_id"`
	RedirectURI string `url:"redirect_uri"`
	Scope       string `url:"scope"`
	State       string `url:"state"`
}
