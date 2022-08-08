package variant

import (
	"context"
	"fmt"
	"net/http"

	"github.com/lengzuo/goshopify/common"
)

func setupDeleteEndpoint() common.Endpoint {
	return common.Endpoint{
		Path:   "/products/%s/variants/%s.json",
		Method: http.MethodDelete,
	}
}

func (c impl) Delete(ctx context.Context, productID, variantID int64) error {
	endpoint := c.deleteEndpoint
	path := fmt.Sprintf(endpoint.Path,
		common.Int64Str(productID),
		common.Int64Str(variantID),
	)
	err := c.call(ctx, endpoint.Method, path, nil, nil)
	if err != nil {
		return err
	}
	return nil
}
