package public_test

import (
	"context"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/ksusonic/niti/backend/internal/api/public"
	"github.com/ksusonic/niti/backend/internal/api/public/mocks"
	"github.com/ksusonic/niti/backend/internal/models"
	"github.com/ksusonic/niti/backend/pgk/publicapi"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"go.uber.org/zap"
)

func TestGetProfile(t *testing.T) {
	type fields struct {
		auth       func(ctrl *gomock.Controller) *mocks.Mockauth
		usersRepo  func(ctrl *gomock.Controller) *mocks.MockusersRepo
		eventsRepo func(ctrl *gomock.Controller) *mocks.MockeventsRepo
	}

	tests := []struct {
		name        string
		fields      fields
		setupCtx    func() context.Context
		request     publicapi.GetProfileRequestObject
		expected    publicapi.GetProfileResponseObject
		expectedErr assert.ErrorAssertionFunc
	}{
		{
			name: "user not found",
			fields: fields{
				auth: mocks.NewMockauth,
				usersRepo: func(ctrl *gomock.Controller) *mocks.MockusersRepo {
					mock := mocks.NewMockusersRepo(ctrl)
					mock.EXPECT().
						Get(gomock.Any(), int64(123)).
						Return(nil, models.ErrNotFound)
					return mock
				},
				eventsRepo: func(ctrl *gomock.Controller) *mocks.MockeventsRepo {
					mock := mocks.NewMockeventsRepo(ctrl)
					mock.EXPECT().
						GetUserEvents(gomock.Any(), int64(123)).
						Return(nil, models.ErrNotFound)
					return mock
				},
			},
			setupCtx: func() context.Context {
				ginCtx := &gin.Context{}
				ginCtx.Set(models.ContextKeyTGUserID, int64(123))
				return ginCtx
			},
			request:     publicapi.GetProfileRequestObject{},
			expected:    publicapi.GetProfile404JSONResponse{Message: "profile not found"},
			expectedErr: assert.NoError,
		},
		{
			name: "repository error",
			fields: fields{
				auth: mocks.NewMockauth,
				usersRepo: func(ctrl *gomock.Controller) *mocks.MockusersRepo {
					mock := mocks.NewMockusersRepo(ctrl)
					mock.EXPECT().
						Get(gomock.Any(), int64(456)).
						Return(nil, assert.AnError)
					return mock
				},
				eventsRepo: func(ctrl *gomock.Controller) *mocks.MockeventsRepo {
					mock := mocks.NewMockeventsRepo(ctrl)
					mock.EXPECT().
						GetUserEvents(gomock.Any(), int64(456)).
						Return(nil, assert.AnError)
					return mock
				},
			},
			setupCtx: func() context.Context {
				ginCtx := &gin.Context{}
				ginCtx.Set(models.ContextKeyTGUserID, int64(456))
				return ginCtx
			},
			request:     publicapi.GetProfileRequestObject{},
			expected:    nil,
			expectedErr: assert.Error,
		},
		{
			name: "success with minimal user data",
			fields: fields{
				auth: mocks.NewMockauth,
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
				eventsRepo: func(ctrl *gomock.Controller) *mocks.MockeventsRepo {
					mock := mocks.NewMockeventsRepo(ctrl)
					mock.EXPECT().
						GetUserEvents(gomock.Any(), int64(789)).
						Return([]models.EventEnriched{}, nil)
					return mock
				},
			},
			setupCtx: func() context.Context {
				ginCtx := &gin.Context{}
				ginCtx.Set(models.ContextKeyTGUserID, int64(789))
				return ginCtx
			},
			request: publicapi.GetProfileRequestObject{},
			expected: publicapi.GetProfile200JSONResponse{
				TelegramId:    789,
				Username:      "testuser",
				FirstName:     "Test",
				LastName:      nil,
				AvatarUrl:     nil,
				IsDj:          false,
				Subscriptions: []publicapi.Event{},
			},
			expectedErr: assert.NoError,
		},
		{
			name: "success with full user data",
			fields: fields{
				auth: mocks.NewMockauth,
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
				eventsRepo: func(ctrl *gomock.Controller) *mocks.MockeventsRepo {
					mock := mocks.NewMockeventsRepo(ctrl)
					mock.EXPECT().
						GetUserEvents(gomock.Any(), int64(999)).
						Return([]models.EventEnriched{}, nil)
					return mock
				},
			},
			setupCtx: func() context.Context {
				ginCtx := &gin.Context{}
				ginCtx.Set(models.ContextKeyTGUserID, int64(999))
				return ginCtx
			},
			request: publicapi.GetProfileRequestObject{},
			expected: publicapi.GetProfile200JSONResponse{
				TelegramId:    999,
				Username:      "fulltestuser",
				FirstName:     "Full",
				LastName:      func() *string { s := "User"; return &s }(),
				AvatarUrl:     func() *string { s := "https://example.com/avatar.jpg"; return &s }(),
				IsDj:          true,
				Subscriptions: []publicapi.Event{},
			},
			expectedErr: assert.NoError,
		},
		{
			name: "success with DJ user",
			fields: fields{
				auth: mocks.NewMockauth,
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
				eventsRepo: func(ctrl *gomock.Controller) *mocks.MockeventsRepo {
					mock := mocks.NewMockeventsRepo(ctrl)
					mock.EXPECT().
						GetUserEvents(gomock.Any(), int64(555)).
						Return([]models.EventEnriched{}, nil)
					return mock
				},
			},
			setupCtx: func() context.Context {
				ginCtx := &gin.Context{}
				ginCtx.Set(models.ContextKeyTGUserID, int64(555))
				return ginCtx
			},
			request: publicapi.GetProfileRequestObject{},
			expected: publicapi.GetProfile200JSONResponse{
				TelegramId:    555,
				Username:      "djuser",
				FirstName:     "DJ",
				LastName:      nil,
				AvatarUrl:     nil,
				IsDj:          true,
				Subscriptions: []publicapi.Event{},
			},
			expectedErr: assert.NoError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			ctx := tt.setupCtx()

			srv := public.NewAPI(
				tt.fields.auth(ctrl),
				tt.fields.usersRepo(ctrl),
				nil,
				tt.fields.eventsRepo(ctrl),
				zap.NewNop(),
			)

			result, err := srv.GetProfile(ctx, tt.request)
			assert.Equal(t, tt.expected, result)
			tt.expectedErr(t, err)
		})
	}
}

func TestGetProfile_PanicOnMissingUserID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	srv := public.NewAPI(
		mocks.NewMockauth(ctrl),
		mocks.NewMockusersRepo(ctrl),
		nil,
		mocks.NewMockeventsRepo(ctrl),
		zap.NewNop(),
	)

	// Context without user ID should cause panic in MustTGUserID
	ctx := &gin.Context{}

	assert.Panics(t, func() {
		_, _ = srv.GetProfile(ctx, publicapi.GetProfileRequestObject{})
	})
}
