package report

import (
	"context"

	dto "github.com/lengzuo/goshopify/adminapi/dto"

	"github.com/lengzuo/goshopify/common"
)

type API interface {
	Create(ctx context.Context, req dto.ReportCollection) (*dto.ReportCollection, error)
	Get(ctx context.Context, req dto.GetReportRequest) (*dto.GetReportResponse, error)
	GetOne(ctx context.Context, reportID int64) (*dto.ReportCollection, error)
	Update(ctx context.Context, req dto.UpdateReportRequest) (*dto.ReportCollection, error)
	Delete(ctx context.Context, reportID int64) error
}

// impl call product APIs
type impl struct {
	call common.Call
	// Setup all APIs
	createEndpoint common.Endpoint
	deleteEndpoint common.Endpoint
	getEndpoint    common.Endpoint
	getOneEndpoint common.Endpoint
	updateEndpoint common.Endpoint
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
