package discountcode

import (
	"context"
	"fmt"
	"net/http"

	"github.com/lengzuo/goshopify/adminapi/dto"

	"github.com/lengzuo/goshopify/common"
)

func setupUpdateEndpoint() common.Endpoint {
	return common.Endpoint{
		Path:   "/price_rules/%s/discount_codes/%s.json",
		Method: http.MethodPut,
	}
}

func (c impl) Update(ctx context.Context, priceRuleID int64, req dto.DiscountCodeCollection) (*dto.DiscountCodeCollection, error) {
	endpoint := c.updateEndpoint
	path := fmt.Sprintf(endpoint.Path,
		common.Int64Str(priceRuleID),
		common.Int64Str(req.DiscountCode.ID),
	)
	resp := new(dto.DiscountCodeCollection)
	err := c.call(ctx, endpoint.Method, path, req, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
