package metafield

import (
	"context"

	adminapidto "github.com/lengzuo/goshopify/adminapi/common"
	"github.com/lengzuo/goshopify/adminapi/dto"
	"github.com/lengzuo/goshopify/common"
)

type API interface {
	Count(ctx context.Context) (*dto.CountResponse, error)
	Create(ctx context.Context, collection adminapidto.Collection, collectionID int64, req dto.MetafieldCollection) (*dto.MetafieldCollection, error)
	Delete(ctx context.Context, metafieldID int64) error
	Get(ctx context.Context, req dto.GetMetafieldRequest) (*dto.GetMetafieldResponse, error)
	GetOne(ctx context.Context, metafieldID int64) (*dto.MetafieldCollection, error)
	Update(ctx context.Context, req dto.MetafieldCollection) (*dto.MetafieldCollection, error)
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
	queryEndpoint  common.Endpoint
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
