package pricerule

import (
	"context"
	"net/http"

	"github.com/lengzuo/goshopify/adminapi/dto"
	"github.com/lengzuo/goshopify/common"
)

func setupCountEndpoint() common.Endpoint {
	return common.Endpoint{
		Path:   "/price_rules/count.json",
		Method: http.MethodGet,
	}
}

func (c impl) Count(ctx context.Context) (*dto.CountResponse, error) {
	endpoint := c.countEndpoint
	resp := new(dto.CountResponse)
	err := c.call(ctx, endpoint.Method, endpoint.Path, nil, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
