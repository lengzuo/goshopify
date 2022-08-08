package customer

import (
	"context"
	"fmt"
	"net/http"

	dto "github.com/lengzuo/goshopify/adminapi/dto"

	"github.com/lengzuo/goshopify/common"
)

func setupGetOrderEndpoint() common.Endpoint {
	return common.Endpoint{
		Path:   "/customers/%s/orders.json",
		Method: http.MethodGet,
	}
}

func (c impl) GetOrder(ctx context.Context, customerID int64, req dto.GetCustomerOrderRequest) (*dto.CustomerCollection, error) {
	endpoint := c.getOrderEndpoint
	path := fmt.Sprintf(endpoint.Path, common.Int64Str(customerID))
	resp := new(dto.CustomerCollection)
	err := c.call(ctx, endpoint.Method, path, req, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
