package theme

import (
	"context"
	"fmt"
	"net/http"

	dto "github.com/lengzuo/goshopify/adminapi/dto"

	"github.com/lengzuo/goshopify/common"
)

func setupGetOneEndpoint() common.Endpoint {
	return common.Endpoint{
		Path:   "/themes/%s.json",
		Method: http.MethodGet,
	}
}

func (c impl) GetOne(ctx context.Context, themeID int64, req dto.GetThemeRequest) (*dto.ThemeCollection, error) {
	endpoint := c.getOneEndpoint
	path := fmt.Sprintf(endpoint.Path, common.Int64Str(themeID))
	resp := new(dto.ThemeCollection)
	err := c.call(ctx, endpoint.Method, path, req, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
