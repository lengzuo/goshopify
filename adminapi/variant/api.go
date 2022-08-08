package variant

import (
	"context"

	"github.com/lengzuo/goshopify/adminapi/dto"
	"github.com/lengzuo/goshopify/common"
)

type API interface {
	Count(ctx context.Context) (*dto.CountResponse, error)
	Create(ctx context.Context, req dto.VariantCollection) (*dto.VariantCollection, error)
	Delete(ctx context.Context, productID, variantID int64) error
	Get(ctx context.Context, req dto.GetVariantRequest) (*dto.GetVariantResponse, error)
	GetOne(ctx context.Context, variantID int64, req dto.GetOneRequest) (*dto.VariantCollection, error)
	Update(ctx context.Context, req dto.VariantCollection) (*dto.VariantCollection, error)
}

// impl call variant APIs
type impl struct {
	call common.Call
	// Setup all APIs
	countEndpoint  common.Endpoint
	createEndpoint common.Endpoint
	deleteEndpoint common.Endpoint
	getEndpoint    common.Endpoint
	getOneEndpoint common.Endpoint
	updateEndpoint common.Endpoint
}

func New(call common.Call) *impl {
	return &impl{
		call:           call,
		countEndpoint:  setupCountEndpoint(),
		createEndpoint: setupCreateEndpoint(),
		deleteEndpoint: setupDeleteEndpoint(),
		getEndpoint:    setupGetEndpoint(),
		getOneEndpoint: setupGetOneEndpoint(),
		updateEndpoint: setupUpdateEndpoint(),
	}
}
