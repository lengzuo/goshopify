package webhook

import (
	"context"
	"fmt"
	"net/http"

	"github.com/lengzuo/goshopify/adminapi/dto"

	"github.com/lengzuo/goshopify/common"
)

func setupGetOneEndpoint() common.Endpoint {
	return common.Endpoint{
		Path:   "/webhooks/%s.json",
		Method: http.MethodGet,
	}
}

func (c impl) GetOne(ctx context.Context, webhookID int64, req dto.GetOneRequest) (*dto.WebhookCollection, error) {
	endpoint := c.getOneEndpoint
	path := fmt.Sprintf(endpoint.Path, common.Int64Str(webhookID))
	resp := new(dto.WebhookCollection)
	err := c.call(ctx, endpoint.Method, path, req, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
