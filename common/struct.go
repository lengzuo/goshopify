package common

type Credentials struct {
	Password  string
	APIKey    string
	APISecret string
}

type Endpoint struct {
	Path   string
	Method string
}
