package api

import (
	"github.com/ksusonic/niti/backend/pkg/openapi"
	"go.uber.org/zap"
)

// compile-time implementation check
var _ openapi.StrictServerInterface = (*API)(nil)

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
