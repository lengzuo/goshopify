package discountcode

import (
	"context"
	"fmt"
	"net/http"

	"github.com/lengzuo/goshopify/adminapi/dto"

	"github.com/lengzuo/goshopify/common"
)

func setupCreateEndpoint() common.Endpoint {
	return common.Endpoint{
		Path:   "/price_rules/%s/discount_codes.json",
		Method: http.MethodPost,
	}
}

func (c impl) Create(ctx context.Context, priceRuleID int64, req dto.CreateDiscountCodeRequest) (*dto.DiscountCodeCollection, error) {
	endpoint := c.createEndpoint
	path := fmt.Sprintf(endpoint.Path, common.Int64Str(priceRuleID))
	resp := new(dto.DiscountCodeCollection)
	err := c.call(ctx, endpoint.Method, path, req, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
