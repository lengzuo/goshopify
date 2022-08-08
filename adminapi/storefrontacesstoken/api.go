package storefrontacesstoken

import (
	"context"

	dto "github.com/lengzuo/goshopify/adminapi/dto"

	"github.com/lengzuo/goshopify/common"
)

type API interface {
	Create(ctx context.Context, req dto.StorefrontAccessTokenCollection) (*dto.StorefrontAccessTokenCollection, error)
	Delete(ctx context.Context, storefrontAccessTokenID int64) error
	Get(ctx context.Context) (*dto.GetStorefrontAccessTokenResponse, error)
}

// impl call product APIs
type impl struct {
	call common.Call
	// Setup all APIs
	createEndpoint common.Endpoint
	deleteEndpoint common.Endpoint
	getEndpoint    common.Endpoint
}

func New(call common.Call) *impl {
	return &impl{
		call:           call,
		createEndpoint: setupCreateEndpoint(),
		deleteEndpoint: setupDeleteEndpoint(),
		getEndpoint:    setupGetEndpoint(),
	}
}
