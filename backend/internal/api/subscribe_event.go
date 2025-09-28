package api

import (
	"context"

	"github.com/ksusonic/niti/backend/internal/models"
	"github.com/ksusonic/niti/backend/pkg/openapi"
	"go.uber.org/zap"
)

func (a *API) SubscribeEvent(ctx context.Context, request openapi.SubscribeEventRequestObject) (openapi.SubscribeEventResponseObject, error) {
	tgUserID := models.MustTGUserID(ctx)

	err := a.subscriptionsRepo.CreateSubscription(ctx, tgUserID, request.Id)
	if err != nil {
		a.logger.Error("create subscription", zap.Error(err))

		return openapi.SubscribeEvent500JSONResponse{Message: err.Error()}, nil
	}

	return openapi.SubscribeEvent200JSONResponse{}, nil
}
