package address

import (
	"context"
	"fmt"
	"net/http"

	dto "github.com/lengzuo/goshopify/adminapi/dto"

	"github.com/lengzuo/goshopify/common"
)

func setupBulkEndpoint() common.Endpoint {
	return common.Endpoint{
		Path:   "/customers/%s/addresses/set.json",
		Method: http.MethodPut,
	}
}

func (c impl) Bulk(ctx context.Context, customerID int64, req dto.BulkUpdateAddressRequest) error {
	endpoint := c.bulkEndpoint
	path := fmt.Sprint(endpoint.Path, common.Int64Str(customerID))
	resp := new(dto.AddressCollection)
	err := c.call(ctx, endpoint.Method, path, req, resp)
	if err != nil {
		return err
	}
	return nil
}
