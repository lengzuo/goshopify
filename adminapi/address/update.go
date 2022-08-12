package address

import (
	"context"
	"fmt"
	"net/http"

	dto "github.com/lengzuo/goshopify/adminapi/dto"

	"github.com/lengzuo/goshopify/common"
)

func setupUpdateEndpoint() common.Endpoint {
	return common.Endpoint{
		Path:   "/customers/%s/addresses/%s.json",
		Method: http.MethodPut,
	}
}

func (c impl) Update(ctx context.Context, req dto.UpdateAddressRequest) (*dto.AddressCollection, error) {
	endpoint := c.updateEndpoint
	path := fmt.Sprintf(endpoint.Path,
		common.Int64Str(req.Address.CustomerID),
		common.Int64Str(req.Address.ID),
	)
	resp := new(dto.AddressCollection)
	err := c.call(ctx, endpoint.Method, path, req, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
