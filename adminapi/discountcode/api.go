package discountcode

import (
	"context"

	dto "github.com/lengzuo/goshopify/adminapi/dto"

	"github.com/lengzuo/goshopify/common"
)

type API interface {
	Batch(ctx context.Context, priceRuleID int64, req dto.PriceRuleDiscountCodeCollection) (*dto.CreatePriceRuleDiscountCodeResponse, error)
	Count(ctx context.Context, req dto.DiscountCodeCountRequest) (*dto.CountResponse, error)
	Create(ctx context.Context, priceRuleID int64, req dto.CreateDiscountCodeRequest) (*dto.DiscountCodeCollection, error)
	Delete(ctx context.Context, priceRuleID, discountCodeID int64) error
	Get(ctx context.Context, priceRuleID int64) (*dto.GetDiscountCodeResponse, error)
	GetBatch(ctx context.Context, priceRuleID, batchID int64) (*dto.PriceRuleDiscountCodeCollection, error)
	GetCreation(ctx context.Context, priceRuleID, batchID int64) (*dto.CreatePriceRuleDiscountCodeResponse, error)
	GetOne(ctx context.Context, priceRuleID, discountCodeID int64) (*dto.DiscountCodeCollection, error)
	Lookup(ctx context.Context, req dto.DiscountLookCountRequest) error
	Update(ctx context.Context, priceRuleID int64, req dto.DiscountCodeCollection) (*dto.DiscountCodeCollection, error)
}

// impl call product APIs
type impl struct {
	call common.Call
	// Setup all APIs
	batchEndpoint       common.Endpoint
	countEndpoint       common.Endpoint
	createEndpoint      common.Endpoint
	deleteEndpoint      common.Endpoint
	getEndpoint         common.Endpoint
	getBatchEndpoint    common.Endpoint
	getCreationEndpoint common.Endpoint
	getOneEndpoint      common.Endpoint
	lookupEndpoint      common.Endpoint
	updateEndpoint      common.Endpoint
}

func New(call common.Call) *impl {
	return &impl{
		call:                call,
		batchEndpoint:       setupBatchEndpoint(),
		countEndpoint:       setupCountEndpoint(),
		createEndpoint:      setupCreateEndpoint(),
		deleteEndpoint:      setupDeleteEndpoint(),
		getEndpoint:         setupGetEndpoint(),
		getBatchEndpoint:    setupGetBatchEndpoint(),
		getCreationEndpoint: setupGetCreationEndpoint(),
		getOneEndpoint:      setupGetOneEndpoint(),
		lookupEndpoint:      setupLookupEndpoint(),
		updateEndpoint:      setupUpdateEndpoint(),
	}
}
