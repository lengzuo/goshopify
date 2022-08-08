package variant

import (
	"context"
	"net/http"

	"github.com/lengzuo/goshopify/adminapi/dto"

	"github.com/lengzuo/goshopify/common"
)

func setupGetEndpoint() common.Endpoint {
	return common.Endpoint{
		Path:   "/variants.json",
		Method: http.MethodGet,
	}
}

func (c impl) Get(ctx context.Context, req dto.GetVariantRequest) (*dto.GetVariantResponse, error) {
	endpoint := c.getEndpoint
	resp := new(dto.GetVariantResponse)
	err := c.call(ctx, endpoint.Method, endpoint.Path, req, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
