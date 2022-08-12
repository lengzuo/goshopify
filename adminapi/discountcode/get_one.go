package discountcode

import (
	"context"
	"fmt"
	"net/http"

	"github.com/lengzuo/goshopify/adminapi/dto"

	"github.com/lengzuo/goshopify/common"
)

func setupGetOneEndpoint() common.Endpoint {
	return common.Endpoint{
		Path:   "/price_rules/%s/discount_codes/%s.json",
		Method: http.MethodGet,
	}
}

func (c impl) GetOne(ctx context.Context, priceRuleID, discountCodeID int64) (*dto.DiscountCodeCollection, error) {
	endpoint := c.getOneEndpoint
	path := fmt.Sprintf(endpoint.Path,
		common.Int64Str(priceRuleID),
		common.Int64Str(discountCodeID),
	)
	resp := new(dto.DiscountCodeCollection)
	err := c.call(ctx, endpoint.Method, path, nil, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
