package api

import (
	"context"

	"github.com/ksusonic/niti/backend/internal/utils"
	"github.com/ksusonic/niti/backend/pgk/genapi"
	"go.uber.org/zap"
)

func (a *API) AuthTelegramInitData(ctx context.Context, request genapi.AuthTelegramInitDataRequestObject) (genapi.AuthTelegramInitDataResponseObject, error) {
	if request.Body == nil || request.Body.InitData == nil || *request.Body.InitData == "" {
		return genapi.AuthTelegramInitData400JSONResponse{Message: "invalid request"}, nil
	}

	initData, err := a.auth.ParseInitData(*request.Body.InitData)
	if err != nil {
		a.logger.Debug("validate request", zap.Error(err), zap.String("init_data", *request.Body.InitData))
		return genapi.AuthTelegramInitData400JSONResponse{Message: "invalid token"}, nil
	}

	tokens, err := a.auth.GenerateTokens(ctx, initData.User.ID)
	if err != nil {
		a.logger.Error("generate token", zap.Error(err), zap.Int64("user_id", initData.User.ID))
		return nil, err
	}

	// TODO: database

	return genapi.AuthTelegramInitData200JSONResponse{
		AccessToken:  tokens.AccessToken,
		RefreshToken: tokens.RefreshToken,
		User: genapi.User{
			TelegramId: initData.User.ID,
			FirstName:  initData.User.FirstName,
			LastName:   utils.NilIfEmpty(initData.User.LastName),
			Username:   initData.User.Username,
			AvatarUrl:  initData.User.PhotoURL,
			IsDj:       false, // TODO
		},
	}, nil
}
