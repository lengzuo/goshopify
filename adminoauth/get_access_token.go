package adminoauth

import (
	"context"
	"net/http"

	"github.com/lengzuo/goshopify/adminoauth/dto"
	"github.com/lengzuo/goshopify/common"
)

func setUpAccessTokenEndpoint() common.Endpoint {
	return common.Endpoint{
		Path:   "/access_token",
		Method: http.MethodPost,
	}
}

func (c *impl) GetAndSetAccessToken(ctx context.Context, code string) (*dto.GetAccessTokenResponse, error) {
	endpoint := c.accessTokenEndpoint
	resp := new(dto.GetAccessTokenResponse)
	err := c.Call(ctx,
		endpoint.Method,
		endpoint.Path,
		dto.GetAccessTokenRequest{
			ClientID:     c.credentials.APIKey,
			ClientSecret: c.credentials.APISecret,
			Code:         code,
		},
		resp,
	)
	if err != nil {
		return nil, err
	}

	c.token = resp.AccessToken

	return resp, nil
}
