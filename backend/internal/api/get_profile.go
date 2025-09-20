package api

import (
	"context"

	"github.com/ksusonic/niti/backend/pgk/genapi"
)

func (a *API) GetProfile(ctx context.Context, request genapi.GetProfileRequestObject) (genapi.GetProfileResponseObject, error) {
	return genapi.GetProfile200JSONResponse{}, nil
}
