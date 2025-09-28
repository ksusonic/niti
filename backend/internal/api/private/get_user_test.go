package private_test

import (
	"context"
	"testing"

	"github.com/ksusonic/niti/backend/internal/api/private"
	"github.com/ksusonic/niti/backend/internal/api/private/mocks"
	"github.com/ksusonic/niti/backend/internal/models"
	"github.com/ksusonic/niti/backend/pkg/privateapi"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"go.uber.org/zap"
)

func TestGetUserByTelegramId(t *testing.T) {
	type fields struct {
		usersRepo func(ctrl *gomock.Controller) *mocks.MockusersRepo
	}

	tests := []struct {
		name        string
		fields      fields
		request     privateapi.GetUserByTelegramIdRequestObject
		expected    privateapi.GetUserByTelegramIdResponseObject
		expectedErr assert.ErrorAssertionFunc
	}{
		{
			name: "user not found",
			fields: fields{
				usersRepo: func(ctrl *gomock.Controller) *mocks.MockusersRepo {
					mock := mocks.NewMockusersRepo(ctrl)
					mock.EXPECT().
						Get(gomock.Any(), int64(123)).
						Return(nil, models.ErrNotFound)
					return mock
				},
			},
			request:     privateapi.GetUserByTelegramIdRequestObject{TelegramId: 123},
			expected:    privateapi.GetUserByTelegramId404JSONResponse{Message: "user not found"},
			expectedErr: assert.NoError,
		},
		{
			name: "repository error",
			fields: fields{
				usersRepo: func(ctrl *gomock.Controller) *mocks.MockusersRepo {
					mock := mocks.NewMockusersRepo(ctrl)
					mock.EXPECT().
						Get(gomock.Any(), int64(456)).
						Return(nil, assert.AnError)
					return mock
				},
			},
			request:     privateapi.GetUserByTelegramIdRequestObject{TelegramId: 456},
			expected:    privateapi.GetUserByTelegramId500JSONResponse{Message: "Internal server error"},
			expectedErr: assert.NoError,
		},
		{
			name: "success with minimal user data",
			fields: fields{
				usersRepo: func(ctrl *gomock.Controller) *mocks.MockusersRepo {
					mock := mocks.NewMockusersRepo(ctrl)
					user := &models.User{
						TelegramID: 789,
						Username:   "testuser",
						FirstName:  "Test",
						LastName:   nil,
						AvatarURL:  nil,
						IsDJ:       false,
					}
					mock.EXPECT().
						Get(gomock.Any(), int64(789)).
						Return(user, nil)
					return mock
				},
			},
			request: privateapi.GetUserByTelegramIdRequestObject{TelegramId: 789},
			expected: privateapi.GetUserByTelegramId200JSONResponse{
				TelegramId: 789,
				Username:   "testuser",
				FirstName:  "Test",
				LastName:   nil,
				AvatarUrl:  nil,
				IsDj:       false,
			},
			expectedErr: assert.NoError,
		},
		{
			name: "success with full user data",
			fields: fields{
				usersRepo: func(ctrl *gomock.Controller) *mocks.MockusersRepo {
					mock := mocks.NewMockusersRepo(ctrl)
					lastName := "User"
					avatarURL := "https://example.com/avatar.jpg"
					user := &models.User{
						TelegramID: 999,
						Username:   "fulltestuser",
						FirstName:  "Full",
						LastName:   &lastName,
						AvatarURL:  &avatarURL,
						IsDJ:       true,
					}
					mock.EXPECT().
						Get(gomock.Any(), int64(999)).
						Return(user, nil)
					return mock
				},
			},
			request: privateapi.GetUserByTelegramIdRequestObject{TelegramId: 999},
			expected: privateapi.GetUserByTelegramId200JSONResponse{
				TelegramId: 999,
				Username:   "fulltestuser",
				FirstName:  "Full",
				LastName:   func() *string { s := "User"; return &s }(),
				AvatarUrl:  func() *string { s := "https://example.com/avatar.jpg"; return &s }(),
				IsDj:       true,
			},
			expectedErr: assert.NoError,
		},
		{
			name: "success with DJ user",
			fields: fields{
				usersRepo: func(ctrl *gomock.Controller) *mocks.MockusersRepo {
					mock := mocks.NewMockusersRepo(ctrl)
					user := &models.User{
						TelegramID: 555,
						Username:   "djuser",
						FirstName:  "DJ",
						LastName:   nil,
						AvatarURL:  nil,
						IsDJ:       true,
					}
					mock.EXPECT().
						Get(gomock.Any(), int64(555)).
						Return(user, nil)
					return mock
				},
			},
			request: privateapi.GetUserByTelegramIdRequestObject{TelegramId: 555},
			expected: privateapi.GetUserByTelegramId200JSONResponse{
				TelegramId: 555,
				Username:   "djuser",
				FirstName:  "DJ",
				LastName:   nil,
				AvatarUrl:  nil,
				IsDj:       true,
			},
			expectedErr: assert.NoError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			srv := private.NewAPI(
				tt.fields.usersRepo(ctrl),
				zap.NewNop(),
			)

			result, err := srv.GetUserByTelegramId(context.Background(), tt.request)
			assert.Equal(t, tt.expected, result)
			tt.expectedErr(t, err)
		})
	}
}
