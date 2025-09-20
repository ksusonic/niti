package api

import (
	"github.com/ksusonic/niti/backend/pgk/genapi"
	"go.uber.org/zap"
)

// compile-time implementation check
var _ genapi.StrictServerInterface = (*API)(nil)

type API struct {
	auth   auth
	logger *zap.Logger
}

func NewAPI(
	auth auth,
	logger *zap.Logger,
) *API {
	return &API{
		auth:   auth,
		logger: logger,
	}
}
