package inventoryitem

import (
	"context"
	"fmt"
	"net/http"

	"github.com/lengzuo/goshopify/adminapi/dto"
	"github.com/lengzuo/goshopify/common"
)

func setupGetOneEndpoint() common.Endpoint {
	return common.Endpoint{
		Path:   "/inventory_items/%s.json",
		Method: http.MethodGet,
	}
}

func (c impl) GetOne(ctx context.Context, inventoryItemID int64) (*dto.InventoryItemCollection, error) {
	endpoint := c.getOneEndpoint
	path := fmt.Sprintf(endpoint.Path, common.Int64Str(inventoryItemID))
	resp := new(dto.InventoryItemCollection)
	err := c.call(ctx, endpoint.Method, path, nil, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
