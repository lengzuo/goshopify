package draftorder

import (
	"context"
	"fmt"
	"net/http"

	"github.com/lengzuo/goshopify/adminapi/dto"
	"github.com/lengzuo/goshopify/common"
)

func setupCompleteEndpoint() common.Endpoint {
	return common.Endpoint{
		Path:   "/draft_orders/%s/complete.json",
		Method: http.MethodPut,
	}
}

func (c impl) Complete(ctx context.Context, draftOrderID int64, req dto.CompleteDraftOrderRequest) (*dto.DraftOrderCollection, error) {
	endpoint := c.completeEndpoint
	path := fmt.Sprintf(endpoint.Path, common.Int64Str(draftOrderID))
	resp := new(dto.DraftOrderCollection)
	err := c.call(ctx, endpoint.Method, path, req, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
