package public

import (
	"context"

	"github.com/ksusonic/niti/backend/pgk/publicapi"
)

func (a *API) Healthcheck(_ context.Context, _ publicapi.HealthcheckRequestObject) (publicapi.HealthcheckResponseObject, error) {
	return publicapi.Healthcheck200JSONResponse{Status: "ok"}, nil
}
