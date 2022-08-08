package customer

import (
	"context"
	"net/http"

	adminapidto "github.com/lengzuo/goshopify/adminapi/dto"

	"github.com/lengzuo/goshopify/common"
)

func setupCountEndpoint() common.Endpoint {
	return common.Endpoint{
		Path:   "/customers/count.json",
		Method: http.MethodGet,
	}
}

func (c impl) Count(ctx context.Context) (*adminapidto.CountResponse, error) {
	endpoint := c.countEndpoint
	resp := new(adminapidto.CountResponse)
	err := c.call(ctx, endpoint.Method, endpoint.Path, nil, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
