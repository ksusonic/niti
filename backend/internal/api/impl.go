package api

import (
	"github.com/ksusonic/niti/backend/pgk/genapi"
	"go.uber.org/zap"
)

// compile-time implementation check
var _ genapi.StrictServerInterface = (*Server)(nil)

type Server struct {
	auth   auth
	logger *zap.Logger
}

func NewServer(
	auth auth,
	logger *zap.Logger,
) *Server {
	return &Server{
		auth:   auth,
		logger: logger,
	}
}
