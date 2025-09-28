package api

import (
	"context"

	"github.com/ksusonic/niti/backend/pkg/openapi"
)

func (a *API) Healthcheck(_ context.Context, _ openapi.HealthcheckRequestObject) (openapi.HealthcheckResponseObject, error) {
	return openapi.Healthcheck200JSONResponse{Status: "ok"}, nil
}
