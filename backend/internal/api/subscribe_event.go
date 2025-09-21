package api

import (
	"context"

	"github.com/ksusonic/niti/backend/internal/models"
	"github.com/ksusonic/niti/backend/pgk/genapi"
	"go.uber.org/zap"
)

func (a *API) SubscribeEvent(ctx context.Context, request genapi.SubscribeEventRequestObject) (genapi.SubscribeEventResponseObject, error) {
	tgUserID := models.MustTGUserID(ctx)

	err := a.subscriptionsRepo.CreateSubscription(ctx, tgUserID, request.Id)
	if err != nil {
		a.logger.Error("create subscription", zap.Error(err))

		return genapi.SubscribeEvent500JSONResponse{Message: err.Error()}, nil
	}

	return genapi.SubscribeEvent200JSONResponse{}, nil
}
