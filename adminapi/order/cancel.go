package order

import (
	"context"
	"fmt"
	"net/http"

	"github.com/lengzuo/goshopify/adminapi/dto"
	"github.com/lengzuo/goshopify/common"
)

func setupCancelEndpoint() common.Endpoint {
	return common.Endpoint{
		Path:   "/orders/%s/cancel.json",
		Method: http.MethodPost,
	}
}

func (c impl) Cancel(ctx context.Context, orderID int64, req dto.CancelOrderRequest) (*dto.OrderCollection, error) {
	endpoint := c.cancelEndpoint
	path := fmt.Sprint(endpoint.Path, common.Int64Str(orderID))
	resp := new(dto.OrderCollection)
	err := c.call(ctx, endpoint.Method, path, req, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
