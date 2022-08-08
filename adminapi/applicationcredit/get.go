package applicationcredit

import (
	"context"
	"net/http"

	"github.com/lengzuo/goshopify/adminapi/dto"

	"github.com/lengzuo/goshopify/common"
)

func setupGetEndpoint() common.Endpoint {
	return common.Endpoint{
		Path:   "/application_credits.json",
		Method: http.MethodGet,
	}
}

func (c impl) Get(ctx context.Context, req dto.GetApplicationCreditRequest) (*dto.GetApplicationCreditResponse, error) {
	endpoint := c.getEndpoint
	resp := new(dto.GetApplicationCreditResponse)
	err := c.call(ctx, endpoint.Method, endpoint.Path, req, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
