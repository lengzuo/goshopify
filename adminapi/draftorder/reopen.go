package draftorder

import (
	"context"
	"fmt"
	"net/http"

	"github.com/lengzuo/goshopify/adminapi/dto"

	"github.com/lengzuo/goshopify/common"
)

func setupSendInvoiceEndpoint() common.Endpoint {
	return common.Endpoint{
		Path:   "/draft_orders/%s/send_invoice.json",
		Method: http.MethodPost,
	}
}

func (c impl) SendInvoice(ctx context.Context, draftOrderID int64, req dto.DraftOrderInvoice) (*dto.DraftOrderInvoice, error) {
	endpoint := c.sendInvoiceEndpoint
	path := fmt.Sprint(endpoint.Path, common.Int64Str(draftOrderID))
	resp := new(dto.DraftOrderInvoice)
	err := c.call(ctx, endpoint.Method, path, req, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
