package customer

import (
	"context"
	"fmt"
	"net/http"

	"github.com/lengzuo/goshopify/common"
)

func setupDeleteEndpoint() common.Endpoint {
	return common.Endpoint{
		Path:   "/customers/%s.json",
		Method: http.MethodDelete,
	}
}

func (c impl) Delete(ctx context.Context, customerID int64) error {
	endpoint := c.deleteEndpoint
	path := fmt.Sprintf(endpoint.Path, common.Int64Str(customerID))
	err := c.call(ctx, endpoint.Method, path, nil, nil)
	if err != nil {
		return err
	}
	return nil
}
