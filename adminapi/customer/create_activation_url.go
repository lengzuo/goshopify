package customer

import (
	"context"
	"fmt"
	"net/http"

	"github.com/lengzuo/goshopify/adminapi/dto"
	"github.com/lengzuo/goshopify/common"
)

func setupCreateActivationURLEndpoint() common.Endpoint {
	return common.Endpoint{
		Path:   "/customers/%s/account_activation_url.json",
		Method: http.MethodPost,
	}
}

func (c impl) CreateActivationURL(ctx context.Context, customerID int64) (*dto.CreateActivationURLResponse, error) {
	endpoint := c.createActivationURLEndpoint
	path := fmt.Sprintf(endpoint.Path, common.Int64Str(customerID))
	resp := new(dto.CreateActivationURLResponse)
	err := c.call(ctx, endpoint.Method, path, nil, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
