package event

import (
	"context"
	"fmt"
	"net/http"

	"github.com/lengzuo/goshopify/adminapi/dto"
	"github.com/lengzuo/goshopify/common"
)

func setupGetOneEndpoint() common.Endpoint {
	return common.Endpoint{
		Path:   "/events/%s.json",
		Method: http.MethodGet,
	}
}

func (c impl) GetOne(ctx context.Context, eventID int64, req dto.GetOneRequest) (*dto.EventCollection, error) {
	endpoint := c.getOneEndpoint
	path := fmt.Sprintf(endpoint.Path, common.Int64Str(eventID))
	resp := new(dto.EventCollection)
	err := c.call(ctx, endpoint.Method, path, req, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
