package image

import (
	"context"
	"fmt"
	"net/http"

	"github.com/lengzuo/goshopify/adminapi/dto"

	"github.com/lengzuo/goshopify/common"
)

func setupCreateEndpoint() common.Endpoint {
	return common.Endpoint{
		Path:   "/products/%s/images.json",
		Method: http.MethodPost,
	}
}

func (c impl) Create(ctx context.Context, productID int64, req dto.ImageCollection) (*dto.ImageCollection, error) {
	endpoint := c.createEndpoint
	path := fmt.Sprintf(endpoint.Path, common.Int64Str(productID))
	resp := new(dto.ImageCollection)
	err := c.call(ctx, endpoint.Method, path, req, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
