package private

import (
	"github.com/ksusonic/niti/backend/pgk/privateapi"
	"go.uber.org/zap"
)

// compile-time implementation check
var _ privateapi.StrictServerInterface = (*API)(nil)

type API struct {
	usersRepo usersRepo
	logger    *zap.Logger
}

func NewAPI(
	usersRepo usersRepo,
	logger *zap.Logger,
) *API {
	return &API{
		usersRepo: usersRepo,
		logger:    logger,
	}
}
