package draftorder

import (
	"context"
	"fmt"
	"net/http"

	"github.com/lengzuo/goshopify/common"
)

func setupDeleteEndpoint() common.Endpoint {
	return common.Endpoint{
		Path:   "/draft_orders/%s.json",
		Method: http.MethodDelete,
	}
}

func (c impl) Delete(ctx context.Context, draftOrderID int64) error {
	endpoint := c.deleteEndpoint
	path := fmt.Sprint(endpoint.Path, common.Int64Str(draftOrderID))
	err := c.call(ctx, endpoint.Method, path, nil, nil)
	if err != nil {
		return err
	}
	return nil
}
