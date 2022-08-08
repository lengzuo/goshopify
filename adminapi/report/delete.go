package report

import (
	"context"
	"fmt"
	"net/http"

	"github.com/lengzuo/goshopify/adminapi/dto"

	"github.com/lengzuo/goshopify/common"
)

func setupDeleteEndpoint() common.Endpoint {
	return common.Endpoint{
		Path:   "/reports/%s.json",
		Method: http.MethodDelete,
	}
}

func (c impl) Delete(ctx context.Context, reportID int64) error {
	endpoint := c.deleteEndpoint
	path := fmt.Sprintf(endpoint.Path, common.Int64Str(reportID))
	resp := new(dto.StorefrontAccessToken)
	err := c.call(ctx, endpoint.Method, path, nil, resp)
	if err != nil {
		return err
	}
	return nil
}
