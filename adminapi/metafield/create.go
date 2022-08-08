package metafield

import (
	"context"
	"fmt"
	"net/http"

	"github.com/lengzuo/goshopify/adminapi/dto"

	adminapidto "github.com/lengzuo/goshopify/adminapi/common"

	"github.com/lengzuo/goshopify/common"
)

func setupCreateEndpoint() common.Endpoint {
	return common.Endpoint{
		Path:   "/%s/%s/metafields.json",
		Method: http.MethodPost,
	}
}

func (c impl) Create(ctx context.Context, collection adminapidto.Collection, collectionID int64, req dto.MetafieldCollection) (*dto.MetafieldCollection, error) {
	endpoint := c.createEndpoint
	path := fmt.Sprintf(
		endpoint.Path,
		collection.String(),
		common.Int64Str(collectionID),
	)
	resp := new(dto.MetafieldCollection)
	err := c.call(ctx, endpoint.Method, path, req, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
