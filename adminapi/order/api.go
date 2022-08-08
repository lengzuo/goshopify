package order

import (
	"context"

	"github.com/lengzuo/goshopify/adminapi/dto"
	adminapidto "github.com/lengzuo/goshopify/adminapi/dto"
	"github.com/lengzuo/goshopify/common"
)

type API interface {
	Cancel(ctx context.Context, orderID int64, req dto.CancelOrderRequest) (*dto.OrderCollection, error)
	Close(ctx context.Context, orderID int64) (*dto.OrderCollection, error)
	Count(ctx context.Context, req dto.CountOrderRequest) (*adminapidto.CountResponse, error)
	Create(ctx context.Context, req dto.OrderCollection) (*dto.OrderCollection, error)
	Delete(ctx context.Context, orderID int64) error
	Get(ctx context.Context, req dto.GetOrderRequest) (*dto.GetOrderResponse, error)
	GetOne(ctx context.Context, orderID int64, req dto.GetOneRequest) (*dto.OrderCollection, error)
	Reopen(ctx context.Context, orderID int64) (*dto.OrderCollection, error)
	Update(ctx context.Context, req dto.OrderCollection) (*dto.OrderCollection, error)
}

// impl call product APIs
type impl struct {
	call common.Call
	// Setup all APIs
	cancelEndpoint common.Endpoint
	closeEndpoint  common.Endpoint
	countEndpoint  common.Endpoint
	createEndpoint common.Endpoint
	deleteEndpoint common.Endpoint
	getEndpoint    common.Endpoint
	getOneEndpoint common.Endpoint
	reopenEndpoint common.Endpoint
	updateEndpoint common.Endpoint
}

func New(call common.Call) *impl {
	return &impl{
		call:           call,
		cancelEndpoint: setupCancelEndpoint(),
		closeEndpoint:  setupCloseEndpoint(),
		countEndpoint:  setupCountEndpoint(),
		createEndpoint: setupCreateEndpoint(),
		deleteEndpoint: setupDeleteEndpoint(),
		getEndpoint:    setupGetEndpoint(),
		getOneEndpoint: setupGetOneEndpoint(),
		reopenEndpoint: setupReopenEndpoint(),
		updateEndpoint: setupUpdateEndpoint(),
	}
}
