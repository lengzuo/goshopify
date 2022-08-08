package image

import (
	"context"
	"fmt"
	"net/http"

	adminapidto "github.com/lengzuo/goshopify/adminapi/dto"

	"github.com/lengzuo/goshopify/common"
)

func setupCountEndpoint() common.Endpoint {
	return common.Endpoint{
		Path:   "/products/%s/images/count.json",
		Method: http.MethodGet,
	}
}

func (c impl) Count(ctx context.Context, productID int64) (*adminapidto.CountResponse, error) {
	endpoint := c.countEndpoint
	path := fmt.Sprintf(endpoint.Path, common.Int64Str(productID))
	resp := new(adminapidto.CountResponse)
	err := c.call(ctx, endpoint.Method, path, nil, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
