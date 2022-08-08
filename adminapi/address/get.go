package address

import (
	"context"
	"fmt"
	"net/http"

	"github.com/lengzuo/goshopify/adminapi/dto"

	"github.com/lengzuo/goshopify/common"
)

func setupGetEndpoint() common.Endpoint {
	return common.Endpoint{
		Path:   "/customers/%s/addresses.json",
		Method: http.MethodGet,
	}
}

func (c impl) Get(ctx context.Context, customerID int64, req dto.GetAddressRequest) (*dto.GetAddressResponse, error) {
	endpoint := c.getEndpoint
	path := fmt.Sprintf(endpoint.Path, common.Int64Str(customerID))
	resp := new(dto.GetAddressResponse)
	err := c.call(ctx, endpoint.Method, path, req, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
