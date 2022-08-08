package image

import (
	"context"
	"fmt"
	"net/http"

	"github.com/lengzuo/goshopify/adminapi/dto"

	"github.com/lengzuo/goshopify/common"
)

func setupUpdateEndpoint() common.Endpoint {
	return common.Endpoint{
		Path:   "/products/%s/images/%s.json",
		Method: http.MethodPut,
	}
}

func (c impl) Update(ctx context.Context, req dto.ImageCollection) (*dto.ImageCollection, error) {
	endpoint := c.updateEndpoint
	path := fmt.Sprintf(endpoint.Path, common.Int64Str(req.Image.ProductID), common.Int64Str(req.Image.ID))
	resp := new(dto.ImageCollection)
	err := c.call(ctx, endpoint.Method, path, req, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
