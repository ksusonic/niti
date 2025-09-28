package api

import (
	"context"

	"github.com/ksusonic/niti/backend/pkg/openapi"
	"go.uber.org/zap"
)

func (a *API) AuthRefreshToken(ctx context.Context, request openapi.AuthRefreshTokenRequestObject) (openapi.AuthRefreshTokenResponseObject, error) {
	if request.Body == nil || request.Body.RefreshToken == "" {
		return openapi.AuthRefreshToken400JSONResponse{Message: "invalid request"}, nil
	}

	oldRefreshToken, err := a.auth.ValidateRefreshToken(ctx, request.Body.RefreshToken)
	if err != nil {
		return openapi.AuthRefreshToken400JSONResponse{Message: err.Error()}, nil
	}

	jwTokens, err := a.auth.RollTokens(ctx, oldRefreshToken)
	if err != nil {
		a.logger.Error("roll tokens", zap.Error(err), zap.Any("old_refresh_token", *oldRefreshToken))

		return openapi.AuthRefreshToken400JSONResponse{Message: err.Error()}, nil
	}

	return openapi.AuthRefreshToken200JSONResponse{
		AccessToken:  jwTokens.AccessToken,
		RefreshToken: jwTokens.RefreshToken,
		ExpiresIn:    int(jwTokens.ExpiresIn.Seconds()),
	}, nil
}
