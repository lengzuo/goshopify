package discountcode

import (
	"context"
	"fmt"
	"net/http"

	"github.com/lengzuo/goshopify/adminapi/dto"

	"github.com/lengzuo/goshopify/common"
)

func setupBatchEndpoint() common.Endpoint {
	return common.Endpoint{
		Path:   "/price_rules/%s/batch.json",
		Method: http.MethodPost,
	}
}

func (c impl) Batch(ctx context.Context, priceRuleID int64, req dto.PriceRuleDiscountCodeCollection) (*dto.CreatePriceRuleDiscountCodeResponse, error) {
	endpoint := c.batchEndpoint
	path := fmt.Sprintf(endpoint.Path, common.Int64Str(priceRuleID))
	resp := new(dto.CreatePriceRuleDiscountCodeResponse)
	err := c.call(ctx, endpoint.Method, path, req, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
