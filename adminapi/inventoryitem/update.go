package inventoryitem

import (
	"context"
	"fmt"
	"net/http"

	"github.com/lengzuo/goshopify/adminapi/dto"
	"github.com/lengzuo/goshopify/common"
)

func setupUpdateEndpoint() common.Endpoint {
	return common.Endpoint{
		Path:   "/inventory_items/%s.json",
		Method: http.MethodPut,
	}
}

func (c impl) Update(ctx context.Context, req dto.InventoryItemCollection) (*dto.InventoryItemCollection, error) {
	endpoint := c.updateEndpoint
	path := fmt.Sprintf(endpoint.Path, common.Int64Str(req.InventoryItem.ID))
	resp := new(dto.InventoryItemCollection)
	err := c.call(ctx, endpoint.Method, path, req, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
