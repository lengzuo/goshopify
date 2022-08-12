package pricerule

import (
	"context"
	"fmt"
	"net/http"

	"github.com/lengzuo/goshopify/adminapi/dto"

	"github.com/lengzuo/goshopify/common"
)

func setupGetOneEndpoint() common.Endpoint {
	return common.Endpoint{
		Path:   "/price_rules/%s.json",
		Method: http.MethodGet,
	}
}

func (c impl) GetOne(ctx context.Context, priceRuleID int64) (*dto.PriceRuleCollection, error) {
	endpoint := c.getOneEndpoint
	path := fmt.Sprintf(endpoint.Path, common.Int64Str(priceRuleID))
	resp := new(dto.PriceRuleCollection)
	err := c.call(ctx, endpoint.Method, path, nil, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
