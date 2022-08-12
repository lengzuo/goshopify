package event

import (
	"context"
	"net/http"

	"github.com/lengzuo/goshopify/adminapi/dto"
	"github.com/lengzuo/goshopify/common"
)

func setupCountEndpoint() common.Endpoint {
	return common.Endpoint{
		Path:   "/events/count.json",
		Method: http.MethodGet,
	}
}

func (c impl) Count(ctx context.Context, req dto.CountEventRequest) (*dto.CountResponse, error) {
	endpoint := c.countEndpoint
	resp := new(dto.CountResponse)
	err := c.call(ctx, endpoint.Method, endpoint.Path, req, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
