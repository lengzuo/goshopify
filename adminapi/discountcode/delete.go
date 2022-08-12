package discountcode

import (
	"context"
	"fmt"
	"net/http"

	"github.com/lengzuo/goshopify/common"
)

func setupDeleteEndpoint() common.Endpoint {
	return common.Endpoint{
		Path:   "/price_rules/%s/discount_codes/%s.json",
		Method: http.MethodDelete,
	}
}

func (c impl) Delete(ctx context.Context, priceRuleID, discountCodeID int64) error {
	endpoint := c.deleteEndpoint
	path := fmt.Sprintf(endpoint.Path,
		common.Int64Str(priceRuleID),
		common.Int64Str(discountCodeID),
	)
	err := c.call(ctx, endpoint.Method, path, nil, nil)
	if err != nil {
		return err
	}
	return nil
}
