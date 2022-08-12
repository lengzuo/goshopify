package discountcode

import (
	"context"
	"net/http"

	"github.com/lengzuo/goshopify/adminapi/dto"

	"github.com/lengzuo/goshopify/common"
)

func setupLookupEndpoint() common.Endpoint {
	return common.Endpoint{
		Path:   "/discount_codes/lookup.json",
		Method: http.MethodGet,
	}
}

func (c impl) Lookup(ctx context.Context, req dto.DiscountLookCountRequest) error {
	endpoint := c.lookupEndpoint
	err := c.call(ctx, endpoint.Method, endpoint.Path, req, nil)
	if err != nil {
		return err
	}
	return nil
}
