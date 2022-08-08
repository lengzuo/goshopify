package address

import (
	"context"
	"fmt"
	"net/http"

	"github.com/lengzuo/goshopify/adminapi/dto"

	"github.com/lengzuo/goshopify/common"
)

func setupCreateEndpoint() common.Endpoint {
	return common.Endpoint{
		Path:   "/customers/%s/addresses.json",
		Method: http.MethodPost,
	}
}

func (c impl) Create(ctx context.Context, customerID int64, req dto.AddressCollection) (*dto.AddressCollection, error) {
	endpoint := c.createEndpoint
	path := fmt.Sprintf(endpoint.Path, common.Int64Str(customerID))
	resp := new(dto.AddressCollection)
	err := c.call(ctx, endpoint.Method, path, req, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
