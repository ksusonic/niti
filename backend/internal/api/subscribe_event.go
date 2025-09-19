package api

import (
	"context"

	"github.com/ksusonic/niti/backend/pgk/genapi"
)

func (s *Server) SubscribeEvent(ctx context.Context, request genapi.SubscribeEventRequestObject) (genapi.SubscribeEventResponseObject, error) {
	return genapi.SubscribeEvent200JSONResponse{}, nil
}
