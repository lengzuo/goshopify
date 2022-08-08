package image

import (
	"context"
	"fmt"
	"net/http"

	"github.com/lengzuo/goshopify/adminapi/dto"
	"github.com/lengzuo/goshopify/common"
)

func setupGetOneEndpoint() common.Endpoint {
	return common.Endpoint{
		Path:   "/products/%s/images/%s.json",
		Method: http.MethodGet,
	}
}

func (c impl) GetOne(ctx context.Context, productID, imageID int64, req dto.GetOneRequest) (*dto.ImageCollection, error) {
	endpoint := c.getOneEndpoint
	path := fmt.Sprintf(endpoint.Path, common.Int64Str(productID), common.Int64Str(imageID))
	resp := new(dto.ImageCollection)
	err := c.call(ctx, endpoint.Method, path, req, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
