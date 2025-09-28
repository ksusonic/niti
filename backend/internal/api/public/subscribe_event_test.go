package public_test

import (
	"context"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/ksusonic/niti/backend/internal/api/public"
	"github.com/ksusonic/niti/backend/internal/api/public/mocks"
	"github.com/ksusonic/niti/backend/internal/models"
	"github.com/ksusonic/niti/backend/pkg/publicapi"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"go.uber.org/zap"
)

func TestSubscribeEvent(t *testing.T) {
	type fields struct {
		auth              func(ctrl *gomock.Controller) *mocks.Mockauth
		usersRepo         func(ctrl *gomock.Controller) *mocks.MockusersRepo
		subscriptionsRepo func(ctrl *gomock.Controller) *mocks.MocksubscriptionsRepo
		eventsRepo        func(ctrl *gomock.Controller) *mocks.MockeventsRepo
	}

	tests := []struct {
		name        string
		fields      fields
		setupCtx    func() context.Context
		request     publicapi.SubscribeEventRequestObject
		expected    publicapi.SubscribeEventResponseObject
		expectedErr assert.ErrorAssertionFunc
	}{
		{
			name: "successful subscription",
			fields: fields{
				auth:       mocks.NewMockauth,
				usersRepo:  mocks.NewMockusersRepo,
				eventsRepo: mocks.NewMockeventsRepo,
				subscriptionsRepo: func(ctrl *gomock.Controller) *mocks.MocksubscriptionsRepo {
					mock := mocks.NewMocksubscriptionsRepo(ctrl)
					mock.EXPECT().
						CreateSubscription(gomock.Any(), int64(123), 456).
						Return(nil)
					return mock
				},
			},
			setupCtx: func() context.Context {
				ginCtx := &gin.Context{}
				ginCtx.Set(models.ContextKeyTGUserID, int64(123))
				return ginCtx
			},
			request: publicapi.SubscribeEventRequestObject{
				Id: 456,
			},
			expected:    publicapi.SubscribeEvent200JSONResponse{},
			expectedErr: assert.NoError,
		},
		{
			name: "repository error",
			fields: fields{
				auth:       mocks.NewMockauth,
				usersRepo:  mocks.NewMockusersRepo,
				eventsRepo: mocks.NewMockeventsRepo,
				subscriptionsRepo: func(ctrl *gomock.Controller) *mocks.MocksubscriptionsRepo {
					mock := mocks.NewMocksubscriptionsRepo(ctrl)
					mock.EXPECT().
						CreateSubscription(gomock.Any(), int64(789), 101).
						Return(assert.AnError)
					return mock
				},
			},
			setupCtx: func() context.Context {
				ginCtx := &gin.Context{}
				ginCtx.Set(models.ContextKeyTGUserID, int64(789))
				return ginCtx
			},
			request: publicapi.SubscribeEventRequestObject{
				Id: 101,
			},
			expected: publicapi.SubscribeEvent500JSONResponse{
				Message: assert.AnError.Error(),
			},
			expectedErr: assert.NoError,
		},
		{
			name: "subscription to different event",
			fields: fields{
				auth:       mocks.NewMockauth,
				usersRepo:  mocks.NewMockusersRepo,
				eventsRepo: mocks.NewMockeventsRepo,
				subscriptionsRepo: func(ctrl *gomock.Controller) *mocks.MocksubscriptionsRepo {
					mock := mocks.NewMocksubscriptionsRepo(ctrl)
					mock.EXPECT().
						CreateSubscription(gomock.Any(), int64(555), 999).
						Return(nil)
					return mock
				},
			},
			setupCtx: func() context.Context {
				ginCtx := &gin.Context{}
				ginCtx.Set(models.ContextKeyTGUserID, int64(555))
				return ginCtx
			},
			request: publicapi.SubscribeEventRequestObject{
				Id: 999,
			},
			expected:    publicapi.SubscribeEvent200JSONResponse{},
			expectedErr: assert.NoError,
		},
		{
			name: "subscription with zero event ID",
			fields: fields{
				auth:       mocks.NewMockauth,
				usersRepo:  mocks.NewMockusersRepo,
				eventsRepo: mocks.NewMockeventsRepo,
				subscriptionsRepo: func(ctrl *gomock.Controller) *mocks.MocksubscriptionsRepo {
					mock := mocks.NewMocksubscriptionsRepo(ctrl)
					mock.EXPECT().
						CreateSubscription(gomock.Any(), int64(111), 0).
						Return(nil)
					return mock
				},
			},
			setupCtx: func() context.Context {
				ginCtx := &gin.Context{}
				ginCtx.Set(models.ContextKeyTGUserID, int64(111))
				return ginCtx
			},
			request: publicapi.SubscribeEventRequestObject{
				Id: 0,
			},
			expected:    publicapi.SubscribeEvent200JSONResponse{},
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
				tt.fields.subscriptionsRepo(ctrl),
				tt.fields.eventsRepo(ctrl),
				zap.NewNop(),
			)

			result, err := srv.SubscribeEvent(ctx, tt.request)
			assert.Equal(t, tt.expected, result)
			tt.expectedErr(t, err)
		})
	}
}

func TestSubscribeEvent_PanicOnMissingUserID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	srv := public.NewAPI(
		mocks.NewMockauth(ctrl),
		mocks.NewMockusersRepo(ctrl),
		mocks.NewMocksubscriptionsRepo(ctrl),
		mocks.NewMockeventsRepo(ctrl),
		zap.NewNop(),
	)

	// Context without user ID should cause panic in MustTGUserID
	ctx := &gin.Context{}

	assert.Panics(t, func() {
		_, _ = srv.SubscribeEvent(ctx, publicapi.SubscribeEventRequestObject{Id: 123})
	})
}
