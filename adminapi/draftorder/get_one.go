package draftorder

import (
	"context"
	"fmt"
	"net/http"

	"github.com/lengzuo/goshopify/adminapi/dto"

	"github.com/lengzuo/goshopify/common"
)

func setupGetOneEndpoint() common.Endpoint {
	return common.Endpoint{
		Path:   "/draft_orders/%s.json",
		Method: http.MethodGet,
	}
}

func (c impl) GetOne(ctx context.Context, draftOrderID int64, req dto.GetOneRequest) (*dto.DraftOrderCollection, error) {
	endpoint := c.getOneEndpoint
	path := fmt.Sprintf(endpoint.Path, common.Int64Str(draftOrderID))
	resp := new(dto.DraftOrderCollection)
	err := c.call(ctx, endpoint.Method, path, req, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
