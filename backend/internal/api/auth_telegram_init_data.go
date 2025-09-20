package api

import (
	"context"

	"github.com/ksusonic/niti/backend/internal/models"
	"github.com/ksusonic/niti/backend/pgk/genapi"
	"go.uber.org/zap"
	"golang.org/x/sync/errgroup"
)

func (a *API) AuthTelegramInitData(ctx context.Context, request genapi.AuthTelegramInitDataRequestObject) (genapi.AuthTelegramInitDataResponseObject, error) {
	if request.Body == nil || request.Body.InitData == nil || *request.Body.InitData == "" {
		return genapi.AuthTelegramInitData400JSONResponse{Message: "invalid request"}, nil
	}

	userData, err := a.auth.ParseInitData(*request.Body.InitData)
	if err != nil {
		a.logger.Debug("validate request", zap.Error(err), zap.String("init_data", *request.Body.InitData))
		return genapi.AuthTelegramInitData400JSONResponse{Message: "invalid token"}, nil
	}

	var (
		tokens models.JWTokens
		user   *models.User
	)

	eg, gCtx := errgroup.WithContext(ctx)

	eg.Go(func() (err error) {
		tokens, err = a.auth.GenerateTokens(gCtx, userData.TelegramID)
		if err != nil {
			a.logger.Error("generate token", zap.Error(err), zap.Int64("user_id", user.TelegramID))
		}

		return err
	})

	eg.Go(func() (err error) {
		user, err = a.usersRepo.Create(gCtx, userData)
		if err != nil {
			a.logger.Error("create user", zap.Error(err), zap.Int64("user_id", userData.TelegramID))
		}

		return err
	})

	if err := eg.Wait(); err != nil {
		return genapi.AuthTelegramInitData500JSONResponse{Message: "internal server error"}, nil
	}

	return genapi.AuthTelegramInitData200JSONResponse{
		AccessToken:  tokens.AccessToken,
		RefreshToken: tokens.RefreshToken,
		User: genapi.User{
			TelegramId: user.TelegramID,
			FirstName:  user.FirstName,
			LastName:   user.LastName,
			Username:   user.Username,
			AvatarUrl:  user.AvatarURL,
			IsDj:       user.IsDJ,
		},
	}, nil
}
