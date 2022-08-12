package pricerule

import (
	"context"
	dto "github.com/lengzuo/goshopify/adminapi/dto"

	"github.com/lengzuo/goshopify/common"
)

type API interface {
	Count(ctx context.Context) (*dto.CountResponse, error)
	Create(ctx context.Context, req dto.PriceRuleCollection) (*dto.PriceRuleCollection, error)
	Delete(ctx context.Context, priceRuleID int64) error
	Get(ctx context.Context, req dto.GetPriceRuleRequest) (*dto.GetPriceRuleResponse, error)
	GetOne(ctx context.Context, priceRuleID int64) (*dto.PriceRuleCollection, error)
	Update(ctx context.Context, req dto.PriceRuleCollection) (*dto.PriceRuleCollection, error)
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
