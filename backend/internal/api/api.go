package api

import (
	"github.com/ksusonic/niti/backend/pgk/genapi"
	"go.uber.org/zap"
)

// compile-time implementation check
var _ genapi.StrictServerInterface = (*API)(nil)

type API struct {
	auth              auth
	usersRepo         usersRepo
	subscriptionsRepo subscriptionsRepo
	eventsRepo        eventsRepo
	logger            *zap.Logger
}

func NewAPI(
	auth auth,
	usersRepo usersRepo,
	subscriptionsRepo subscriptionsRepo,
	eventsRepo eventsRepo,
	logger *zap.Logger,
) *API {
	return &API{
		auth:              auth,
		usersRepo:         usersRepo,
		subscriptionsRepo: subscriptionsRepo,
		eventsRepo:        eventsRepo,
		logger:            logger,
	}
}
