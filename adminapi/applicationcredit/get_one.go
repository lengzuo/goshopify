package applicationcredit

import (
	"context"
	"fmt"
	"net/http"

	"github.com/lengzuo/goshopify/adminapi/dto"

	"github.com/lengzuo/goshopify/common"
)

func setupGetOneEndpoint() common.Endpoint {
	return common.Endpoint{
		Path:   "/application_credits/%s.json",
		Method: http.MethodGet,
	}
}

func (c impl) GetOne(ctx context.Context, applicationChargeID int64) (*dto.ApplicationCreditCollection, error) {
	endpoint := c.getOneEndpoint
	path := fmt.Sprintf(endpoint.Path, common.Int64Str(applicationChargeID))
	resp := new(dto.ApplicationCreditCollection)
	err := c.call(ctx, endpoint.Method, path, nil, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
