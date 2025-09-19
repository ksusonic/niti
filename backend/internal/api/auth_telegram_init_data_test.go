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
	initdata "github.com/telegram-mini-apps/init-data-golang"
	"go.uber.org/mock/gomock"
	"go.uber.org/zap"
)

func TestAuthTelegramInitData(t *testing.T) {
	ctx := context.Background()

	type fields struct {
		auth func(ctrl *gomock.Controller) *mocks.Mockauth
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
				auth: mocks.NewMockauth,
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
				auth: mocks.NewMockauth,
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
						Return(&initdata.InitData{User: initdata.User{ID: 123}}, nil)

					mock.EXPECT().
						GenerateToken(int64(123)).
						Return(models.JWTAuth{}, assert.AnError)

					return mock
				},
			},
			expected:    nil,
			expectedErr: assert.Error,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			srv := api.NewServer(
				tt.fields.auth(ctrl),
				zap.NewNop(),
			)

			result, err := srv.AuthTelegramInitData(ctx, tt.request)
			assert.Equal(t, tt.expected, result)
			tt.expectedErr(t, err)
		})
	}
}
