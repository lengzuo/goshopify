package pricerule

import (
	"context"
	"net/http"

	"github.com/lengzuo/goshopify/adminapi/dto"

	"github.com/lengzuo/goshopify/common"
)

func setupCreateEndpoint() common.Endpoint {
	return common.Endpoint{
		Path:   "/price_rules.json",
		Method: http.MethodPost,
	}
}

func (c impl) Create(ctx context.Context, req dto.PriceRuleCollection) (*dto.PriceRuleCollection, error) {
	endpoint := c.createEndpoint
	resp := new(dto.PriceRuleCollection)
	err := c.call(ctx, endpoint.Method, endpoint.Path, req, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
