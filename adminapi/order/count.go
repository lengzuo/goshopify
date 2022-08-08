package order

import (
	"context"
	"net/http"

	"github.com/lengzuo/goshopify/adminapi/dto"
	adminapidto "github.com/lengzuo/goshopify/adminapi/dto"

	"github.com/lengzuo/goshopify/common"
)

func setupCountEndpoint() common.Endpoint {
	return common.Endpoint{
		Path:   "/orders/count.json",
		Method: http.MethodGet,
	}
}

func (c impl) Count(ctx context.Context, req dto.CountOrderRequest) (*adminapidto.CountResponse, error) {
	endpoint := c.countEndpoint
	resp := new(adminapidto.CountResponse)
	err := c.call(ctx, endpoint.Method, endpoint.Path, req, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
