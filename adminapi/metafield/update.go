package metafield

import (
	"context"
	"fmt"
	"net/http"

	"github.com/lengzuo/goshopify/adminapi/dto"

	"github.com/lengzuo/goshopify/common"
)

func setupUpdateEndpoint() common.Endpoint {
	return common.Endpoint{
		Path:   "/metafields/%s.json",
		Method: http.MethodPut,
	}
}

func (c impl) Update(ctx context.Context, req dto.MetafieldCollection) (*dto.MetafieldCollection, error) {
	endpoint := c.updateEndpoint
	path := fmt.Sprint(endpoint.Path, common.Int64Str(req.Metafield.ID))
	resp := new(dto.MetafieldCollection)
	err := c.call(ctx, endpoint.Method, path, req, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
