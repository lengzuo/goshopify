package storefrontacesstoken

import (
	"context"
	"net/http"

	"github.com/lengzuo/goshopify/adminapi/dto"

	"github.com/lengzuo/goshopify/common"
)

func setupGetEndpoint() common.Endpoint {
	return common.Endpoint{
		Path:   "/storefront_access_tokens.json",
		Method: http.MethodGet,
	}
}

func (c impl) Get(ctx context.Context) (*dto.GetStorefrontAccessTokenResponse, error) {
	endpoint := c.getEndpoint
	resp := new(dto.GetStorefrontAccessTokenResponse)
	err := c.call(ctx, endpoint.Method, endpoint.Path, nil, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
