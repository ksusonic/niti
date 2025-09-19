package api

import (
	"context"

	"github.com/ksusonic/niti/backend/internal/utils"
	"github.com/ksusonic/niti/backend/pgk/genapi"
	"go.uber.org/zap"
)

func (s *Server) AuthTelegramInitData(ctx context.Context, request genapi.AuthTelegramInitDataRequestObject) (genapi.AuthTelegramInitDataResponseObject, error) {
	if !validateRequest(request) {
		return genapi.AuthTelegramInitData400JSONResponse{Message: "invalid request"}, nil
	}

	initData, err := s.auth.ParseInitData(*request.Body.InitData)
	if err != nil {
		s.logger.Debug("validate request", zap.Error(err), zap.String("init_data", *request.Body.InitData))
		return genapi.AuthTelegramInitData400JSONResponse{Message: "invalid token"}, nil
	}

	tokens, err := s.auth.GenerateToken(initData.User.ID)
	if err != nil {
		s.logger.Error("generate token", zap.Error(err), zap.Int64("user_id", initData.User.ID))
		return nil, err
	}

	// generate access and refresh tokens
	// accessToken, refreshToken := s.auth.GenerateTokens(initData)

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

func validateRequest(request genapi.AuthTelegramInitDataRequestObject) bool {
	return request.Body != nil && request.Body.InitData != nil && *request.Body.InitData != ""
}
