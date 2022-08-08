package draftorder

import (
	"context"
	"fmt"
	"net/http"

	"github.com/lengzuo/goshopify/adminapi/dto"
	"github.com/lengzuo/goshopify/common"
)

func setupUpdateEndpoint() common.Endpoint {
	return common.Endpoint{
		Path:   "/draft_orders/%s.json",
		Method: http.MethodPut,
	}
}

func (c impl) Update(ctx context.Context, req dto.DraftOrderCollection) (*dto.DraftOrderCollection, error) {
	endpoint := c.updateEndpoint
	path := fmt.Sprint(endpoint.Path, common.Int64Str(req.DraftOrder.ID))
	resp := new(dto.DraftOrderCollection)
	err := c.call(ctx, endpoint.Method, path, req, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
