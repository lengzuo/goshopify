package image

import (
	"context"
	"fmt"
	"net/http"

	"github.com/lengzuo/goshopify/adminapi/dto"

	"github.com/lengzuo/goshopify/common"
)

func setupGetEndpoint() common.Endpoint {
	return common.Endpoint{
		Path:   "/products/%s/images.json",
		Method: http.MethodGet,
	}
}

func (c impl) Get(ctx context.Context, productID int64, req dto.GetImageRequest) (*dto.GetImageResponse, error) {
	endpoint := c.getEndpoint
	path := fmt.Sprintf(endpoint.Path, common.Int64Str(productID))
	resp := new(dto.GetImageResponse)
	err := c.call(ctx, endpoint.Method, path, req, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
