package adminoauth

import (
	"context"
	"net/http"

	"github.com/lengzuo/goshopify/adminoauth/dto"
	"github.com/lengzuo/goshopify/common"
)

func setUpAccessScopesEndpoint() common.Endpoint {
	return common.Endpoint{
		Path:   "/access_scopes.json",
		Method: http.MethodGet,
	}
}

func (c *impl) GetAccessScopes(ctx context.Context) (*dto.GetAccessScopeResponse, error) {
	endpoint := c.accessScopesEndpoint
	resp := new(dto.GetAccessScopeResponse)
	err := c.Call(ctx, endpoint.Method, endpoint.Path, nil, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
