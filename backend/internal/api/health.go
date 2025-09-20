package api

import (
	"context"

	"github.com/ksusonic/niti/backend/pgk/genapi"
)

func (s *Server) Healthcheck(_ context.Context, request genapi.HealthcheckRequestObject) (genapi.HealthcheckResponseObject, error) {
	return genapi.Healthcheck200JSONResponse{Status: "ok"}, nil
}
