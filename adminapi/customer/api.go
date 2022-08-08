package customer

import (
	"context"

	"github.com/lengzuo/goshopify/adminapi/dto"
	"github.com/lengzuo/goshopify/common"
)

type API interface {
	Count(ctx context.Context) (*dto.CountResponse, error)
	Create(ctx context.Context, req dto.CustomerCollection) (*dto.CustomerCollection, error)
	CreateActivationURL(ctx context.Context, customerID int64) (*dto.CreateActivationURLResponse, error)
	Delete(ctx context.Context, customerID int64) error
	Get(ctx context.Context, req dto.GetCustomerRequest) (*dto.GetCustomerResponse, error)
	GetOne(ctx context.Context, customerID int64) (*dto.CustomerCollection, error)
	GetOrder(ctx context.Context, customerID int64, req dto.GetCustomerOrderRequest) (*dto.CustomerCollection, error)
	Update(ctx context.Context, req dto.CustomerCollection) (*dto.CustomerCollection, error)
}

// impl call product APIs
type impl struct {
	call common.Call
	// Setup all APIs
	countEndpoint               common.Endpoint
	createEndpoint              common.Endpoint
	createActivationURLEndpoint common.Endpoint
	deleteEndpoint              common.Endpoint
	getEndpoint                 common.Endpoint
	getOneEndpoint              common.Endpoint
	getOrderEndpoint            common.Endpoint
	searchEndpoint              common.Endpoint
	updateEndpoint              common.Endpoint
}

func New(call common.Call) *impl {
	return &impl{
		call:                        call,
		countEndpoint:               setupCountEndpoint(),
		createEndpoint:              setupCreateEndpoint(),
		createActivationURLEndpoint: setupCreateActivationURLEndpoint(),
		deleteEndpoint:              setupDeleteEndpoint(),
		getEndpoint:                 setupGetEndpoint(),
		getOneEndpoint:              setupGetOneEndpoint(),
		getOrderEndpoint:            setupGetOrderEndpoint(),
		searchEndpoint:              setupSearchEndpoint(),
		updateEndpoint:              setupUpdateEndpoint(),
	}
}
