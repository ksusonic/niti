package public

import (
	"context"
	"fmt"

	"github.com/ksusonic/niti/backend/internal/models"
	"github.com/ksusonic/niti/backend/pgk/publicapi"
	"go.uber.org/zap"
	"golang.org/x/sync/errgroup"
)

func (a *API) AuthTelegramInitData(ctx context.Context, request publicapi.AuthTelegramInitDataRequestObject) (publicapi.AuthTelegramInitDataResponseObject, error) {
	if request.Body == nil || request.Body.InitData == nil || *request.Body.InitData == "" {
		return publicapi.AuthTelegramInitData400JSONResponse{Message: "invalid request"}, nil
	}

	userData, err := a.auth.ParseInitData(*request.Body.InitData)
	if err != nil {
		a.logger.Debug("validate request", zap.Error(err), zap.String("init_data", *request.Body.InitData))
		return publicapi.AuthTelegramInitData400JSONResponse{Message: "invalid token"}, nil
	}

	var (
		tokens models.JWTokens
		user   *models.User
	)

	eg, gCtx := errgroup.WithContext(ctx)

	eg.Go(func() (err error) {
		tokens, err = a.auth.GenerateTokens(gCtx, userData.TelegramID)
		if err != nil {
			return fmt.Errorf("generate tokens: %w", err)
		}

		return err
	})

	eg.Go(func() (err error) {
		user, err = a.usersRepo.Create(gCtx, userData)
		if err != nil {
			return fmt.Errorf("create user: %w", err)
		}

		return err
	})

	if err := eg.Wait(); err != nil {
		a.logger.Error("process user save and gen token", zap.Error(err), zap.Any("user_data", *userData))

		return publicapi.AuthTelegramInitData500JSONResponse{Message: "internal server error"}, nil
	}

	return publicapi.AuthTelegramInitData200JSONResponse{
		AccessToken:  tokens.AccessToken,
		RefreshToken: tokens.RefreshToken,
		User: publicapi.User{
			TelegramId: user.TelegramID,
			FirstName:  user.FirstName,
			LastName:   user.LastName,
			Username:   user.Username,
			AvatarUrl:  user.AvatarURL,
			IsDj:       user.IsDJ,
		},
	}, nil
}
