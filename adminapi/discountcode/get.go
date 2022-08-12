package discountcode

import (
	"context"
	"fmt"
	"net/http"

	"github.com/lengzuo/goshopify/adminapi/dto"

	"github.com/lengzuo/goshopify/common"
)

func setupGetEndpoint() common.Endpoint {
	return common.Endpoint{
		Path:   "/price_rules/%s/discount_codes.json",
		Method: http.MethodGet,
	}
}

func (c impl) Get(ctx context.Context, priceRuleID int64) (*dto.GetDiscountCodeResponse, error) {
	endpoint := c.getEndpoint
	path := fmt.Sprintf(endpoint.Path, common.Int64Str(priceRuleID))
	resp := new(dto.GetDiscountCodeResponse)
	err := c.call(ctx, endpoint.Method, path, nil, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
