package api

import (
	"context"

	"github.com/ksusonic/niti/backend/pgk/genapi"
)

func (a *API) AuthRefreshToken(ctx context.Context, request genapi.AuthRefreshTokenRequestObject) (genapi.AuthRefreshTokenResponseObject, error) {
	return genapi.AuthRefreshToken200JSONResponse{}, nil
}
