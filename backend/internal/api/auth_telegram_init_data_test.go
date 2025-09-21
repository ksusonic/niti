package api_test

import (
	"context"
	"testing"

	"github.com/ksusonic/niti/backend/internal/api"
	"github.com/ksusonic/niti/backend/internal/api/mocks"
	"github.com/ksusonic/niti/backend/internal/models"
	"github.com/ksusonic/niti/backend/internal/utils"
	"github.com/ksusonic/niti/backend/pgk/genapi"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"go.uber.org/zap"
)

func TestAuthTelegramInitData(t *testing.T) {
	ctx := context.Background()

	type fields struct {
		auth      func(ctrl *gomock.Controller) *mocks.Mockauth
		usersRepo func(ctrl *gomock.Controller) *mocks.MockusersRepo
	}

	tests := []struct {
		name        string
		fields      fields
		request     genapi.AuthTelegramInitDataRequestObject
		expected    genapi.AuthTelegramInitDataResponseObject
		expectedErr assert.ErrorAssertionFunc
	}{
		{
			name: "nil body",
			request: genapi.AuthTelegramInitDataRequestObject{
				Body: nil,
			},
			fields: fields{
				auth:      mocks.NewMockauth,
				usersRepo: mocks.NewMockusersRepo,
			},
			expected:    genapi.AuthTelegramInitData400JSONResponse{Message: "invalid request"},
			expectedErr: assert.NoError,
		},
		{
			name: "empty init data",
			request: genapi.AuthTelegramInitDataRequestObject{
				Body: &genapi.AuthTelegramInitDataJSONRequestBody{InitData: utils.Ptr("")},
			},
			fields: fields{
				auth:      mocks.NewMockauth,
				usersRepo: mocks.NewMockusersRepo,
			},
			expected:    genapi.AuthTelegramInitData400JSONResponse{Message: "invalid request"},
			expectedErr: assert.NoError,
		},
		{
			name: "parse error",
			request: genapi.AuthTelegramInitDataRequestObject{
				Body: &genapi.AuthTelegramInitDataJSONRequestBody{InitData: utils.Ptr("bad_data")},
			},
			fields: fields{
				auth: func(ctrl *gomock.Controller) *mocks.Mockauth {
					mock := mocks.NewMockauth(ctrl)
					mock.EXPECT().
						ParseInitData("bad_data").
						Return(nil, assert.AnError)

					return mock
				},
				usersRepo: mocks.NewMockusersRepo,
			},
			expected:    genapi.AuthTelegramInitData400JSONResponse{Message: "invalid token"},
			expectedErr: assert.NoError,
		},
		{
			name: "generate token error",
			request: genapi.AuthTelegramInitDataRequestObject{
				Body: &genapi.AuthTelegramInitDataJSONRequestBody{InitData: utils.Ptr("good_data")},
			},
			fields: fields{
				auth: func(ctrl *gomock.Controller) *mocks.Mockauth {
					mock := mocks.NewMockauth(ctrl)
					mock.EXPECT().
						ParseInitData("good_data").
						Return(&models.User{TelegramID: 123}, nil)

					mock.EXPECT().
						GenerateTokens(gomock.Any(), int64(123)).
						Return(models.JWTokens{}, assert.AnError)

					return mock
				},
				usersRepo: func(ctrl *gomock.Controller) *mocks.MockusersRepo {
					mock := mocks.NewMockusersRepo(ctrl)
					mock.EXPECT().
						Create(gomock.Any(), &models.User{TelegramID: 123}).
						Return(&models.User{TelegramID: 123}, nil)

					return mock
				},
			},
			expected:    genapi.AuthTelegramInitData500JSONResponse{Message: "internal server error"},
			expectedErr: assert.NoError,
		},
		{
			name: "create user error",
			request: genapi.AuthTelegramInitDataRequestObject{
				Body: &genapi.AuthTelegramInitDataJSONRequestBody{InitData: utils.Ptr("good_data")},
			},
			fields: fields{
				auth: func(ctrl *gomock.Controller) *mocks.Mockauth {
					mock := mocks.NewMockauth(ctrl)
					mock.EXPECT().
						ParseInitData("good_data").
						Return(&models.User{TelegramID: 123, FirstName: "John", Username: "john123"}, nil)

					mock.EXPECT().
						GenerateTokens(gomock.Any(), int64(123)).
						Return(models.JWTokens{AccessToken: "access", RefreshToken: "refresh"}, nil)

					return mock
				},
				usersRepo: func(ctrl *gomock.Controller) *mocks.MockusersRepo {
					mock := mocks.NewMockusersRepo(ctrl)
					mock.EXPECT().
						Create(gomock.Any(), &models.User{TelegramID: 123, FirstName: "John", Username: "john123"}).
						Return(nil, assert.AnError)

					return mock
				},
			},
			expected:    genapi.AuthTelegramInitData500JSONResponse{Message: "internal server error"},
			expectedErr: assert.NoError,
		},
		{
			name: "success",
			request: genapi.AuthTelegramInitDataRequestObject{
				Body: &genapi.AuthTelegramInitDataJSONRequestBody{InitData: utils.Ptr("valid_data")},
			},
			fields: fields{
				auth: func(ctrl *gomock.Controller) *mocks.Mockauth {
					mock := mocks.NewMockauth(ctrl)
					mock.EXPECT().
						ParseInitData("valid_data").
						Return(&models.User{
							TelegramID: 456,
							FirstName:  "Jane",
							Username:   "jane456",
							LastName:   utils.Ptr("Doe"),
							AvatarURL:  utils.Ptr("https://example.com/avatar.jpg"),
							IsDJ:       false,
						}, nil)

					mock.EXPECT().
						GenerateTokens(gomock.Any(), int64(456)).
						Return(models.JWTokens{
							AccessToken:  "access_token_123",
							RefreshToken: "refresh_token_456",
						}, nil)

					return mock
				},
				usersRepo: func(ctrl *gomock.Controller) *mocks.MockusersRepo {
					mock := mocks.NewMockusersRepo(ctrl)
					mock.EXPECT().
						Create(gomock.Any(), &models.User{
							TelegramID: 456,
							FirstName:  "Jane",
							Username:   "jane456",
							LastName:   utils.Ptr("Doe"),
							AvatarURL:  utils.Ptr("https://example.com/avatar.jpg"),
							IsDJ:       false,
						}).
						Return(&models.User{
							TelegramID: 456,
							FirstName:  "Jane",
							Username:   "jane456",
							LastName:   utils.Ptr("Doe"),
							AvatarURL:  utils.Ptr("https://example.com/avatar.jpg"),
							IsDJ:       false,
						}, nil)

					return mock
				},
			},
			expected: genapi.AuthTelegramInitData200JSONResponse{
				AccessToken:  "access_token_123",
				RefreshToken: "refresh_token_456",
				User: genapi.User{
					TelegramId: 456,
					FirstName:  "Jane",
					LastName:   utils.Ptr("Doe"),
					Username:   "jane456",
					AvatarUrl:  utils.Ptr("https://example.com/avatar.jpg"),
					IsDj:       false,
				},
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
				nil,
				nil,
				zap.NewNop(),
			)

			result, err := srv.AuthTelegramInitData(ctx, tt.request)
			assert.Equal(t, tt.expected, result)
			tt.expectedErr(t, err)
		})
	}
}
