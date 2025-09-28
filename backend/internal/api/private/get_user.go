package private

import (
	"context"
	"errors"

	"github.com/ksusonic/niti/backend/internal/models"
	"github.com/ksusonic/niti/backend/pkg/privateapi"
	"go.uber.org/zap"
)

func (h *API) GetUserByTelegramId(ctx context.Context, request privateapi.GetUserByTelegramIdRequestObject) (privateapi.GetUserByTelegramIdResponseObject, error) {
	user, err := h.usersRepo.Get(ctx, request.TelegramId)
	if err != nil {
		if errors.Is(err, models.ErrNotFound) {
			return privateapi.GetUserByTelegramId404JSONResponse{
				Message: "user not found",
			}, nil
		}
		h.logger.Error("Failed to get user", zap.Error(err), zap.Int64("telegram_id", request.TelegramId))
		return privateapi.GetUserByTelegramId500JSONResponse{
			Message: "Internal server error",
		}, nil
	}

	return privateapi.GetUserByTelegramId200JSONResponse{
		TelegramId: user.TelegramID,
		Username:   user.Username,
		FirstName:  user.FirstName,
		LastName:   user.LastName,
		AvatarUrl:  user.AvatarURL,
		IsDj:       user.IsDJ,
	}, nil
}
