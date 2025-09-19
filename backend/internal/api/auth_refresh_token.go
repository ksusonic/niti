package api

import (
	"context"

	"github.com/ksusonic/niti/backend/pgk/genapi"
)

func (s *Server) AuthRefreshToken(ctx context.Context, request genapi.AuthRefreshTokenRequestObject) (genapi.AuthRefreshTokenResponseObject, error) {
	return genapi.AuthRefreshToken200JSONResponse{}, nil
}
