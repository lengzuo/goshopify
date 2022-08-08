package theme

import (
	"context"

	"github.com/lengzuo/goshopify/adminapi/dto"
	"github.com/lengzuo/goshopify/common"
)

type API interface {
	Create(ctx context.Context, req dto.ThemeCollection) (*dto.ThemeCollection, error)
	Delete(ctx context.Context, themeID int64) error
	Get(ctx context.Context, req dto.GetThemeRequest) (*dto.GetThemeResponse, error)
	GetOne(ctx context.Context, themeID int64, req dto.GetThemeRequest) (*dto.ThemeCollection, error)
	Update(ctx context.Context, req dto.ThemeCollection) (*dto.ThemeCollection, error)
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
