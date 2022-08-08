package adminoauth

import (
	"context"

	"github.com/lengzuo/goshopify/adminoauth/dto"
)

type API interface {
	GetAccessScopes(ctx context.Context) (*dto.GetAccessScopeResponse, error)
	GetAndSetAccessToken(ctx context.Context, code string) (*dto.GetAccessTokenResponse, error)
}
