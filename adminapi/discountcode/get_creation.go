package discountcode

import (
	"context"
	"fmt"
	"github.com/lengzuo/goshopify/adminapi/dto"
	"github.com/lengzuo/goshopify/common"
	"net/http"
)

func setupGetCreationEndpoint() common.Endpoint {
	return common.Endpoint{
		Path:   "/price_rules/%s/batch/%s.json",
		Method: http.MethodGet,
	}
}

func (c impl) GetCreation(ctx context.Context, priceRuleID, batchID int64) (*dto.CreatePriceRuleDiscountCodeResponse, error) {
	endpoint := c.getCreationEndpoint
	path := fmt.Sprintf(endpoint.Path,
		common.Int64Str(priceRuleID),
		common.Int64Str(batchID),
	)
	resp := new(dto.CreatePriceRuleDiscountCodeResponse)
	err := c.call(ctx, endpoint.Method, path, nil, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
