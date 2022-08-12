package report

import (
	"context"
	"fmt"
	"net/http"

	dto "github.com/lengzuo/goshopify/adminapi/dto"

	"github.com/lengzuo/goshopify/common"
)

func setupUpdateEndpoint() common.Endpoint {
	return common.Endpoint{
		Path:   "/reports/%s.json",
		Method: http.MethodPut,
	}
}

func (c impl) Update(ctx context.Context, req dto.UpdateReportRequest) (*dto.ReportCollection, error) {
	endpoint := c.updateEndpoint
	path := fmt.Sprintf(endpoint.Path, common.Int64Str(req.Report.ID))
	resp := new(dto.ReportCollection)
	err := c.call(ctx, endpoint.Method, path, req, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
