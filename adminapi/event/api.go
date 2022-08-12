package event

import (
	"context"

	dto "github.com/lengzuo/goshopify/adminapi/dto"

	"github.com/lengzuo/goshopify/common"
)

type API interface {
	Count(ctx context.Context, req dto.CountEventRequest) (*dto.CountResponse, error)
	Get(ctx context.Context, req dto.GetEventRequest) (*dto.GetEventResponse, error)
	GetOne(ctx context.Context, eventID int64, req dto.GetOneRequest) (*dto.EventCollection, error)
}

// impl call product APIs
type impl struct {
	call common.Call
	// Setup all APIs
	countEndpoint  common.Endpoint
	getEndpoint    common.Endpoint
	getOneEndpoint common.Endpoint
}

func New(call common.Call) *impl {
	return &impl{
		call:           call,
		countEndpoint:  setupCountEndpoint(),
		getEndpoint:    setupGetEndpoint(),
		getOneEndpoint: setupGetOneEndpoint(),
	}
}
