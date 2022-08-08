package scripttag

import (
	"context"

	dto "github.com/lengzuo/goshopify/adminapi/dto"

	"github.com/lengzuo/goshopify/common"
)

type API interface {
	Create(ctx context.Context, req dto.ScriptTagCollection) (*dto.ScriptTagCollection, error)
	Delete(ctx context.Context, scriptTagID int64) error
	Get(ctx context.Context, req dto.GetScriptTagRequest) (*dto.GetScriptTagResponse, error)
	GetOne(ctx context.Context, scriptTagID int64, req dto.GetScriptTagRequest) (*dto.ScriptTagCollection, error)
	Update(ctx context.Context, req dto.ScriptTagCollection) (*dto.ScriptTagCollection, error)
}

// impl call product APIs
type impl struct {
	call common.Call
	// Setup all APIs
	createEndpoint     common.Endpoint
	deleteEndpoint     common.Endpoint
	getEndpoint        common.Endpoint
	getOneEndpoint     common.Endpoint
	setDefaultEndpoint common.Endpoint
	updateEndpoint     common.Endpoint
}

func New(call common.Call) *impl {
	return &impl{
		call:           call,
		createEndpoint: setupCreateEndpoint(),
		deleteEndpoint: setupDeleteEndpoint(),
		getEndpoint:    setupGetEndpoint(),
		getOneEndpoint: setupGetOneEndpoint(),
		updateEndpoint: setupUpdateEndpoint(),
	}
}
