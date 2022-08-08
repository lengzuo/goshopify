package scripttag

import (
	"context"
	"fmt"
	"net/http"

	dto "github.com/lengzuo/goshopify/adminapi/dto"

	"github.com/lengzuo/goshopify/common"
)

func setupGetOneEndpoint() common.Endpoint {
	return common.Endpoint{
		Path:   "/script_tags/%s.json",
		Method: http.MethodGet,
	}
}

func (c impl) GetOne(ctx context.Context, scriptTagID int64, req dto.GetScriptTagRequest) (*dto.ScriptTagCollection, error) {
	endpoint := c.getOneEndpoint
	path := fmt.Sprintf(endpoint.Path, common.Int64Str(scriptTagID))
	resp := new(dto.ScriptTagCollection)
	err := c.call(ctx, endpoint.Method, path, req, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
