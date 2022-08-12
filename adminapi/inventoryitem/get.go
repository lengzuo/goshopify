package inventoryitem

import (
	"context"
	"net/http"

	"github.com/lengzuo/goshopify/adminapi/dto"

	"github.com/lengzuo/goshopify/common"
)

func setupGetEndpoint() common.Endpoint {
	return common.Endpoint{
		Path:   "/inventory_items.json",
		Method: http.MethodGet,
	}
}

func (c impl) Get(ctx context.Context, req dto.GetInventoryItemRequest) (*dto.GetInventoryItemResponse, error) {
	endpoint := c.getEndpoint
	resp := new(dto.GetInventoryItemResponse)
	err := c.call(ctx, endpoint.Method, endpoint.Path, req, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
