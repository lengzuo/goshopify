package customer

import (
	"context"
	"net/http"

	dto "github.com/lengzuo/goshopify/adminapi/dto"

	"github.com/lengzuo/goshopify/common"
)

func setupSearchEndpoint() common.Endpoint {
	return common.Endpoint{
		Path:   "/customers/search.json",
		Method: http.MethodGet,
	}
}

func (c impl) Search(ctx context.Context, req dto.SearchCustomerRequest) (*dto.GetCustomerResponse, error) {
	endpoint := c.getEndpoint
	resp := new(dto.GetCustomerResponse)
	err := c.call(ctx, endpoint.Method, endpoint.Path, req, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
