package api

import (
	"context"

	"github.com/ksusonic/niti/backend/pgk/genapi"
)

func (s *Server) ListEvents(ctx context.Context, request genapi.ListEventsRequestObject) (genapi.ListEventsResponseObject, error) {
	return genapi.ListEvents200JSONResponse{}, nil
}
