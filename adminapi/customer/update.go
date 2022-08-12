package customer

import (
	"context"
	"fmt"
	"net/http"

	"github.com/lengzuo/goshopify/adminapi/dto"

	"github.com/lengzuo/goshopify/common"
)

func setupUpdateEndpoint() common.Endpoint {
	return common.Endpoint{
		Path:   "/customers/%s.json",
		Method: http.MethodPut,
	}
}

func (c impl) Update(ctx context.Context, req dto.CustomerCollection) (*dto.CustomerCollection, error) {
	endpoint := c.updateEndpoint
	path := fmt.Sprintf(endpoint.Path, common.Int64Str(req.Customer.ID))
	resp := new(dto.CustomerCollection)
	err := c.call(ctx, endpoint.Method, path, req, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
