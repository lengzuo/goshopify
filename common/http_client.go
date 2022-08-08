package common

import (
	"bytes"
	"context"
	"encoding/json"
	"net"
	"net/http"
	"net/url"
	"time"

	"github.com/google/go-querystring/query"
)

const defaultHTTPTimeout = 5 * time.Second

func DefaultHTTPClient() *http.Client {
	return &http.Client{
		Transport: &http.Transport{
			MaxIdleConnsPerHost: 1000,
			DialContext: (&net.Dialer{
				Timeout:   defaultHTTPTimeout,
				KeepAlive: 30 * time.Second,
			}).DialContext,
			TLSHandshakeTimeout: defaultHTTPTimeout,
			DisableKeepAlives:   false,
		},
		Timeout: defaultHTTPTimeout,
	}
}

type Call func(ctx context.Context, method, path string, req, resp interface{}) (err error)
type AuthSetter func(r *http.Request)

type Sender interface {
	Call(ctx context.Context, method, path string, req, resp interface{}) (err error)
	Client() *http.Client
	ErrorDTO() error
	BaseURL() url.URL
}

// Send use to send HTTP request to shopify
func Send(ctx context.Context, sender Sender, method, path string, req, resp interface{}, authSetter AuthSetter) (err error) {
	var httpReq *http.Request
	httpReq, err = MakeRequest(ctx, sender.BaseURL(), method, path, req)
	if err != nil {
		return err
	}

	authSetter(httpReq)

	var httpResp *http.Response
	httpResp, err = sender.Client().Do(httpReq)
	if err != nil {
		return err
	}
	defer httpResp.Body.Close()

	// https://go-recipes.dev/downloading-files-with-go-b298c3efb557#:~:text=The%20copy%20function%20is%20around,is%20to%20go%20with%20io.
	decoder := json.NewDecoder(httpResp.Body)
	if http.StatusOK > httpResp.StatusCode || httpResp.StatusCode >= http.StatusMultipleChoices {
		// Get error
		errResp := sender.ErrorDTO()
		err = decoder.Decode(errResp)
		if err != nil {
			return err
		}
		return errResp
	}

	err = decoder.Decode(resp)
	if err != nil {
		return err
	}
	return nil
}

// MakeRequest construct the http request for api call
func MakeRequest(ctx context.Context, baseURL url.URL, method, path string, params interface{}, options ...Option) (req *http.Request, err error) {
	apiURL, err := url.Parse(baseURL.String() + path)
	if err != nil {
		return nil, err
	}

	var js []byte
	switch method {
	case http.MethodGet, http.MethodDelete:
		var queryValues url.Values
		queryValues, err = query.Values(params)
		if err != nil {
			return nil, err
		}

		for k, values := range apiURL.Query() {
			for _, v := range values {
				queryValues.Add(k, v)
			}
		}
		apiURL.RawQuery = queryValues.Encode()
	default:
		js, err = json.Marshal(params)
		if err != nil {
			return nil, err
		}
	}

	req, err = http.NewRequestWithContext(ctx, method, apiURL.String(), bytes.NewBuffer(js))
	if err != nil {
		return nil, err
	}

	req.Header.Add(HeaderContentType, ApplicationJSON)
	req.Header.Add(HeaderAccept, ApplicationJSON)

	// Append custom header
	for _, opt := range options {
		opt(req)
	}

	return req, nil
}
