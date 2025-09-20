package api_test

import (
	"context"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/ksusonic/niti/backend/internal/api"
	"github.com/ksusonic/niti/backend/internal/api/mocks"
	"github.com/ksusonic/niti/backend/internal/models"
	"github.com/ksusonic/niti/backend/pgk/genapi"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"go.uber.org/zap"
)

func TestAuthRefreshToken(t *testing.T) {
	ctx := context.Background()

	type fields struct {
		auth      func(ctrl *gomock.Controller) *mocks.Mockauth
		usersRepo func(ctrl *gomock.Controller) *mocks.MockusersRepo
	}

	tests := []struct {
		name        string
		fields      fields
		request     genapi.AuthRefreshTokenRequestObject
		expected    genapi.AuthRefreshTokenResponseObject
		expectedErr assert.ErrorAssertionFunc
	}{
		{
			name: "nil body",
			request: genapi.AuthRefreshTokenRequestObject{
				Body: nil,
			},
			fields: fields{
				auth:      mocks.NewMockauth,
				usersRepo: mocks.NewMockusersRepo,
			},
			expected:    genapi.AuthRefreshToken400JSONResponse{Message: "invalid request"},
			expectedErr: assert.NoError,
		},
		{
			name: "empty refresh token",
			request: genapi.AuthRefreshTokenRequestObject{
				Body: &genapi.AuthRefreshTokenJSONRequestBody{RefreshToken: ""},
			},
			fields: fields{
				auth:      mocks.NewMockauth,
				usersRepo: mocks.NewMockusersRepo,
			},
			expected:    genapi.AuthRefreshToken400JSONResponse{Message: "invalid request"},
			expectedErr: assert.NoError,
		},
		{
			name: "validate refresh token error",
			request: genapi.AuthRefreshTokenRequestObject{
				Body: &genapi.AuthRefreshTokenJSONRequestBody{RefreshToken: "invalid_token"},
			},
			fields: fields{
				auth: func(ctrl *gomock.Controller) *mocks.Mockauth {
					mock := mocks.NewMockauth(ctrl)
					mock.EXPECT().
						ValidateRefreshToken(ctx, "invalid_token").
						Return(nil, assert.AnError)
					return mock
				},
				usersRepo: mocks.NewMockusersRepo,
			},
			expected:    genapi.AuthRefreshToken400JSONResponse{Message: assert.AnError.Error()},
			expectedErr: assert.NoError,
		},
		{
			name: "roll tokens error",
			request: genapi.AuthRefreshTokenRequestObject{
				Body: &genapi.AuthRefreshTokenJSONRequestBody{RefreshToken: "valid_token"},
			},
			fields: fields{
				auth: func(ctrl *gomock.Controller) *mocks.Mockauth {
					mock := mocks.NewMockauth(ctrl)

					oldRefreshToken := &models.RefreshToken{
						JTI:       uuid.New(),
						UserID:    123,
						ExpiresAt: time.Now().Add(time.Hour),
						Revoked:   false,
						CreatedAt: time.Now(),
					}

					mock.EXPECT().
						ValidateRefreshToken(ctx, "valid_token").
						Return(oldRefreshToken, nil)

					mock.EXPECT().
						RollTokens(ctx, oldRefreshToken).
						Return(nil, assert.AnError)

					return mock
				},
				usersRepo: mocks.NewMockusersRepo,
			},
			expected:    genapi.AuthRefreshToken400JSONResponse{Message: assert.AnError.Error()},
			expectedErr: assert.NoError,
		},
		{
			name: "success",
			request: genapi.AuthRefreshTokenRequestObject{
				Body: &genapi.AuthRefreshTokenJSONRequestBody{RefreshToken: "valid_refresh_token"},
			},
			fields: fields{
				auth: func(ctrl *gomock.Controller) *mocks.Mockauth {
					mock := mocks.NewMockauth(ctrl)

					oldRefreshToken := &models.RefreshToken{
						JTI:       uuid.New(),
						UserID:    456,
						ExpiresAt: time.Now().Add(time.Hour),
						Revoked:   false,
						CreatedAt: time.Now(),
					}

					newTokens := &models.JWTokens{
						AccessToken:  "new_access_token",
						RefreshToken: "new_refresh_token",
						JTI:          uuid.New(),
						ExpiresIn:    30 * time.Minute,
					}

					mock.EXPECT().
						ValidateRefreshToken(ctx, "valid_refresh_token").
						Return(oldRefreshToken, nil)

					mock.EXPECT().
						RollTokens(ctx, oldRefreshToken).
						Return(newTokens, nil)

					return mock
				},
				usersRepo: mocks.NewMockusersRepo,
			},
			expected: genapi.AuthRefreshToken200JSONResponse{
				AccessToken:  "new_access_token",
				RefreshToken: "new_refresh_token",
				ExpiresIn:    1800, // 30 minutes in seconds
			},
			expectedErr: assert.NoError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			srv := api.NewAPI(
				tt.fields.auth(ctrl),
				tt.fields.usersRepo(ctrl),
				zap.NewNop(),
			)

			result, err := srv.AuthRefreshToken(ctx, tt.request)
			assert.Equal(t, tt.expected, result)
			tt.expectedErr(t, err)
		})
	}
}
