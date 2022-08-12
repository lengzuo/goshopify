package address

import (
	"context"
	"fmt"
	"net/http"

	"github.com/lengzuo/goshopify/adminapi/dto"

	"github.com/lengzuo/goshopify/common"
)

func setupSetDefaultEndpoint() common.Endpoint {
	return common.Endpoint{
		Path:   "/customers/%s/addresses/%s/default.json",
		Method: http.MethodPut,
	}
}

func (c impl) SetDefault(ctx context.Context, customerID, addressID int64) (*dto.AddressCollection, error) {
	endpoint := c.updateEndpoint
	path := fmt.Sprintf(endpoint.Path,
		common.Int64Str(customerID),
		common.Int64Str(addressID),
	)
	resp := new(dto.AddressCollection)
	err := c.call(ctx, endpoint.Method, path, nil, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
