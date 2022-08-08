package applicationcredit

import (
	"context"

	dto "github.com/lengzuo/goshopify/adminapi/dto"

	"github.com/lengzuo/goshopify/common"
)

type API interface {
	Create(ctx context.Context, req dto.ApplicationCreditCollection) (*dto.ApplicationCreditCollection, error)
	Get(ctx context.Context, req dto.GetApplicationCreditRequest) (*dto.GetApplicationCreditResponse, error)
	GetOne(ctx context.Context, applicationCreditID int64) (*dto.ApplicationCreditCollection, error)
}

// impl call product APIs
type impl struct {
	call common.Call
	// Setup all APIs
	createEndpoint common.Endpoint
	getEndpoint    common.Endpoint
	getOneEndpoint common.Endpoint
}

func New(call common.Call) *impl {
	return &impl{
		call:           call,
		createEndpoint: setupCreateEndpoint(),
		getOneEndpoint: setupGetOneEndpoint(),
		getEndpoint:    setupGetEndpoint(),
	}
}
