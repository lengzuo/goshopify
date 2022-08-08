package product

import (
	"context"

	dto "github.com/lengzuo/goshopify/adminapi/dto"

	"github.com/lengzuo/goshopify/common"
)

type API interface {
	Count(ctx context.Context, req dto.CountProductRequest) (*dto.CountResponse, error)
	Create(ctx context.Context, req dto.ProductCollection) (*dto.ProductCollection, error)
	Delete(ctx context.Context, productID int64) error
	Get(ctx context.Context, req dto.GetProductRequest) (*dto.GetProductResponse, error)
	GetOne(ctx context.Context, productID int64, req dto.GetOneRequest) (*dto.ProductCollection, error)
	Update(ctx context.Context, req dto.ProductCollection) (*dto.ProductCollection, error)
}

// impl call product APIs
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
