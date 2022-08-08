package order

import (
	"context"
	"fmt"
	"net/http"

	"github.com/lengzuo/goshopify/adminapi/dto"

	"github.com/lengzuo/goshopify/common"
)

func setupReopenEndpoint() common.Endpoint {
	return common.Endpoint{
		Path:   "/orders/%s/open.json",
		Method: http.MethodPost,
	}
}

func (c impl) Reopen(ctx context.Context, orderID int64) (*dto.OrderCollection, error) {
	endpoint := c.updateEndpoint
	path := fmt.Sprint(endpoint.Path, common.Int64Str(orderID))
	resp := new(dto.OrderCollection)
	err := c.call(ctx, endpoint.Method, path, nil, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
