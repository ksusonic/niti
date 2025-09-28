package public

import (
	"context"

	"github.com/ksusonic/niti/backend/internal/models"
	"github.com/ksusonic/niti/backend/pgk/publicapi"
	"go.uber.org/zap"
)

// under auth middleware
func (a *API) SubscribeEvent(ctx context.Context, request publicapi.SubscribeEventRequestObject) (publicapi.SubscribeEventResponseObject, error) {
	tgUserID := models.MustTGUserID(ctx)

	err := a.subscriptionsRepo.CreateSubscription(ctx, tgUserID, request.Id)
	if err != nil {
		a.logger.Error("create subscription", zap.Error(err))

		return publicapi.SubscribeEvent500JSONResponse{Message: err.Error()}, nil
	}

	return publicapi.SubscribeEvent200JSONResponse{}, nil
}
