package recurringapplicationcharge

import (
	"context"

	dto "github.com/lengzuo/goshopify/adminapi/dto"

	"github.com/lengzuo/goshopify/common"
)

type API interface {
	Create(ctx context.Context, req dto.ApplicationChargeCollection) (*dto.ApplicationChargeCollection, error)
	Delete(ctx context.Context, recurringApplicationChargeID int64) error
	Get(ctx context.Context, req dto.GetApplicationChargeRequest) (*dto.GetApplicationChargeResponse, error)
	GetOne(ctx context.Context, recurringApplicationChargeID int64) (*dto.ApplicationChargeCollection, error)
}

// impl call product APIs
type impl struct {
	call common.Call
	// Setup all APIs
	createEndpoint common.Endpoint
	deleteEndpoint common.Endpoint
	getEndpoint    common.Endpoint
	getOneEndpoint common.Endpoint
}

func New(call common.Call) *impl {
	return &impl{
		call:           call,
		createEndpoint: setupCreateEndpoint(),
		deleteEndpoint: setupDeleteEndpoint(),
		getOneEndpoint: setupGetOneEndpoint(),
		getEndpoint:    setupGetEndpoint(),
	}
}
