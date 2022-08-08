package draftorder

import (
	"context"

	"github.com/lengzuo/goshopify/adminapi/dto"
	"github.com/lengzuo/goshopify/common"
)

type API interface {
	Complete(ctx context.Context, draftOrderID int64, req dto.CompleteDraftOrderRequest) (*dto.DraftOrderCollection, error)
	Count(ctx context.Context, req dto.CountDraftOrderRequest) (*dto.CountResponse, error)
	Create(ctx context.Context, req dto.DraftOrderCollection) (*dto.DraftOrderCollection, error)
	Delete(ctx context.Context, draftOrderID int64) error
	Get(ctx context.Context, req dto.GetDraftOrderRequest) (*dto.GetDraftOrderResponse, error)
	GetOne(ctx context.Context, draftOrderID int64, req dto.GetOneRequest) (*dto.DraftOrderCollection, error)
	SendInvoice(ctx context.Context, draftOrderID int64, req dto.DraftOrderInvoice) (*dto.DraftOrderInvoice, error)
	Update(ctx context.Context, req dto.DraftOrderCollection) (*dto.DraftOrderCollection, error)
}

// impl call product APIs
type impl struct {
	call common.Call
	// Setup all APIs
	completeEndpoint    common.Endpoint
	countEndpoint       common.Endpoint
	createEndpoint      common.Endpoint
	deleteEndpoint      common.Endpoint
	getEndpoint         common.Endpoint
	getOneEndpoint      common.Endpoint
	sendInvoiceEndpoint common.Endpoint
	updateEndpoint      common.Endpoint
}

func New(call common.Call) *impl {
	return &impl{
		call:                call,
		completeEndpoint:    setupCompleteEndpoint(),
		countEndpoint:       setupCountEndpoint(),
		createEndpoint:      setupCreateEndpoint(),
		deleteEndpoint:      setupDeleteEndpoint(),
		getEndpoint:         setupGetEndpoint(),
		getOneEndpoint:      setupGetOneEndpoint(),
		sendInvoiceEndpoint: setupSendInvoiceEndpoint(),
		updateEndpoint:      setupUpdateEndpoint(),
	}
}
