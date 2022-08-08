package order

import (
	"context"
	"fmt"
	"net/http"

	"github.com/lengzuo/goshopify/adminapi/dto"
	"github.com/lengzuo/goshopify/common"
)

func setupCloseEndpoint() common.Endpoint {
	return common.Endpoint{
		Path:   "/orders/%s/close.json",
		Method: http.MethodPost,
	}
}

func (c impl) Close(ctx context.Context, orderID int64) (*dto.OrderCollection, error) {
	endpoint := c.closeEndpoint
	path := fmt.Sprint(endpoint.Path, common.Int64Str(orderID))
	resp := new(dto.OrderCollection)
	err := c.call(ctx, endpoint.Method, path, nil, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
