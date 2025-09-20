package api

import (
	"context"
	"errors"

	"github.com/ksusonic/niti/backend/internal/models"
	"github.com/ksusonic/niti/backend/pgk/genapi"
	"go.uber.org/zap"
)

// under auth middleware
func (a *API) GetProfile(ctx context.Context, request genapi.GetProfileRequestObject) (genapi.GetProfileResponseObject, error) {
	tgUserID := models.MustTGUserID(ctx)

	user, err := a.usersRepo.Get(ctx, tgUserID)
	if err != nil {
		if errors.Is(err, models.ErrNotFound) {
			a.logger.Error("not found user", zap.Int64("user_id", tgUserID))

			return genapi.GetProfile404JSONResponse{Message: "profile not found"}, nil
		}

		a.logger.Error("get user", zap.Int64("user_id", tgUserID))

		return nil, err
	}

	return genapi.GetProfile200JSONResponse{
		TelegramId:    user.TelegramID,
		Username:      user.Username,
		FirstName:     user.FirstName,
		LastName:      user.LastName,
		AvatarUrl:     user.AvatarURL,
		IsDj:          user.IsDJ,
		Subscriptions: nil, // TODO
	}, nil
}
