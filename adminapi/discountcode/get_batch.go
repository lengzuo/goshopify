package discountcode

import (
	"context"
	"fmt"
	"github.com/lengzuo/goshopify/adminapi/dto"
	"github.com/lengzuo/goshopify/common"
	"net/http"
)

func setupGetBatchEndpoint() common.Endpoint {
	return common.Endpoint{
		Path:   "/price_rules/%s/batch/%s/discount_codes.json",
		Method: http.MethodGet,
	}
}

func (c impl) GetBatch(ctx context.Context, priceRuleID, batchID int64) (*dto.PriceRuleDiscountCodeCollection, error) {
	endpoint := c.getBatchEndpoint
	path := fmt.Sprintf(endpoint.Path,
		common.Int64Str(priceRuleID),
		common.Int64Str(batchID),
	)
	resp := new(dto.PriceRuleDiscountCodeCollection)
	err := c.call(ctx, endpoint.Method, path, nil, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
