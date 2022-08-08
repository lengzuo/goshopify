package storefrontacesstoken

import (
	"context"
	"fmt"
	"net/http"

	"github.com/lengzuo/goshopify/common"
)

func setupDeleteEndpoint() common.Endpoint {
	return common.Endpoint{
		Path:   "/storefront_access_tokens/%s.json",
		Method: http.MethodDelete,
	}
}

func (c impl) Delete(ctx context.Context, storefrontAccessTokenID int64) error {
	endpoint := c.deleteEndpoint
	path := fmt.Sprintf(endpoint.Path, common.Int64Str(storefrontAccessTokenID))
	err := c.call(ctx, endpoint.Method, path, nil, nil)
	if err != nil {
		return err
	}
	return nil
}
