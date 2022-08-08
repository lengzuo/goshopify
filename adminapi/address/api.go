package address

import (
	"context"

	"github.com/lengzuo/goshopify/adminapi/dto"
	"github.com/lengzuo/goshopify/common"
)

type API interface {
	Bulk(ctx context.Context, customerID int64, req dto.BulkUpdateAddressRequest) error
	Create(ctx context.Context, customerID int64, req dto.AddressCollection) (*dto.AddressCollection, error)
	Delete(ctx context.Context, customerID int64, addressID int64) error
	Get(ctx context.Context, customerID int64, req dto.GetAddressRequest) (*dto.GetAddressResponse, error)
	GetOne(ctx context.Context, customerID int64, addressID int64) (*dto.AddressCollection, error)
	SetDefault(ctx context.Context, customerID int64, addressID int64) (*dto.AddressCollection, error)
	Update(ctx context.Context, req dto.UpdateAddressRequest) (*dto.AddressCollection, error)
}

// impl call product APIs
type impl struct {
	call common.Call
	// Setup all APIs
	bulkEndpoint       common.Endpoint
	createEndpoint     common.Endpoint
	deleteEndpoint     common.Endpoint
	getEndpoint        common.Endpoint
	getOneEndpoint     common.Endpoint
	setDefaultEndpoint common.Endpoint
	updateEndpoint     common.Endpoint
}

func New(call common.Call) *impl {
	return &impl{
		call:               call,
		bulkEndpoint:       setupBulkEndpoint(),
		createEndpoint:     setupCreateEndpoint(),
		deleteEndpoint:     setupDeleteEndpoint(),
		getEndpoint:        setupGetEndpoint(),
		getOneEndpoint:     setupGetOneEndpoint(),
		setDefaultEndpoint: setupSetDefaultEndpoint(),
		updateEndpoint:     setupUpdateEndpoint(),
	}
}
