package metafield

import (
	"context"
	"fmt"
	"net/http"

	"github.com/lengzuo/goshopify/common"
)

func setupDeleteEndpoint() common.Endpoint {
	return common.Endpoint{
		Path:   "/metafields/%s.json",
		Method: http.MethodDelete,
	}
}

func (c impl) Delete(ctx context.Context, metafieldID int64) error {
	endpoint := c.deleteEndpoint
	path := fmt.Sprint(endpoint.Path, common.Int64Str(metafieldID))
	err := c.call(ctx, endpoint.Method, path, nil, nil)
	if err != nil {
		return err
	}
	return nil
}
