package inventoryitem

import (
	"context"

	dto "github.com/lengzuo/goshopify/adminapi/dto"

	"github.com/lengzuo/goshopify/common"
)

type API interface {
	Get(ctx context.Context, req dto.GetInventoryItemRequest) (*dto.GetInventoryItemResponse, error)
	GetOne(ctx context.Context, inventoryItemID int64) (*dto.InventoryItemCollection, error)
	Update(ctx context.Context, req dto.InventoryItemCollection) (*dto.InventoryItemCollection, error)
}

// impl call product APIs
type impl struct {
	call common.Call
	// Setup all APIs
	getEndpoint    common.Endpoint
	getOneEndpoint common.Endpoint
	updateEndpoint common.Endpoint
}

func New(call common.Call) *impl {
	return &impl{
		call:           call,
		getEndpoint:    setupGetEndpoint(),
		getOneEndpoint: setupGetOneEndpoint(),
		updateEndpoint: setupUpdateEndpoint(),
	}
}
