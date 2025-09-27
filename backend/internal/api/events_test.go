package api_test

import (
	"context"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/ksusonic/niti/backend/internal/api"
	"github.com/ksusonic/niti/backend/internal/api/mocks"
	"github.com/ksusonic/niti/backend/internal/models"
	"github.com/ksusonic/niti/backend/pgk/publicapi"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"go.uber.org/zap"
)

func TestEvents(t *testing.T) {
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
		request     publicapi.EventsRequestObject
		expected    publicapi.EventsResponseObject
		expectedErr assert.ErrorAssertionFunc
	}{
		{
			name: "success with default parameters",
			fields: fields{
				auth:              mocks.NewMockauth,
				usersRepo:         mocks.NewMockusersRepo,
				subscriptionsRepo: mocks.NewMocksubscriptionsRepo,
				eventsRepo: func(ctrl *gomock.Controller) *mocks.MockeventsRepo {
					mock := mocks.NewMockeventsRepo(ctrl)
					mock.EXPECT().
						ListEvents(gomock.Any(), int64(123), 30, 0).
						Return([]models.EventEnriched{}, nil)
					return mock
				},
			},
			setupCtx: func() context.Context {
				ginCtx := &gin.Context{}
				ginCtx.Set(models.ContextKeyTGUserID, int64(123))
				return ginCtx
			},
			request: publicapi.EventsRequestObject{
				Params: publicapi.EventsParams{},
			},
			expected:    publicapi.Events200JSONResponse{},
			expectedErr: assert.NoError,
		},
		{
			name: "success with custom limit and offset",
			fields: fields{
				auth:              mocks.NewMockauth,
				usersRepo:         mocks.NewMockusersRepo,
				subscriptionsRepo: mocks.NewMocksubscriptionsRepo,
				eventsRepo: func(ctrl *gomock.Controller) *mocks.MockeventsRepo {
					mock := mocks.NewMockeventsRepo(ctrl)
					mock.EXPECT().
						ListEvents(gomock.Any(), int64(456), 10, 20).
						Return([]models.EventEnriched{}, nil)
					return mock
				},
			},
			setupCtx: func() context.Context {
				ginCtx := &gin.Context{}
				ginCtx.Set(models.ContextKeyTGUserID, int64(456))
				return ginCtx
			},
			request: publicapi.EventsRequestObject{
				Params: publicapi.EventsParams{
					Limit:  func() *int { l := 10; return &l }(),
					Offset: func() *int { o := 20; return &o }(),
				},
			},
			expected:    publicapi.Events200JSONResponse{},
			expectedErr: assert.NoError,
		},
		{
			name: "success with zero limit uses default",
			fields: fields{
				auth:              mocks.NewMockauth,
				usersRepo:         mocks.NewMockusersRepo,
				subscriptionsRepo: mocks.NewMocksubscriptionsRepo,
				eventsRepo: func(ctrl *gomock.Controller) *mocks.MockeventsRepo {
					mock := mocks.NewMockeventsRepo(ctrl)
					mock.EXPECT().
						ListEvents(gomock.Any(), int64(789), 30, 0).
						Return([]models.EventEnriched{}, nil)
					return mock
				},
			},
			setupCtx: func() context.Context {
				ginCtx := &gin.Context{}
				ginCtx.Set(models.ContextKeyTGUserID, int64(789))
				return ginCtx
			},
			request: publicapi.EventsRequestObject{
				Params: publicapi.EventsParams{
					Limit: func() *int { l := 0; return &l }(),
				},
			},
			expected:    publicapi.Events200JSONResponse{},
			expectedErr: assert.NoError,
		},
		{
			name: "success with events data",
			fields: fields{
				auth:              mocks.NewMockauth,
				usersRepo:         mocks.NewMockusersRepo,
				subscriptionsRepo: mocks.NewMocksubscriptionsRepo,
				eventsRepo: func(ctrl *gomock.Controller) *mocks.MockeventsRepo {
					mock := mocks.NewMockeventsRepo(ctrl)

					startsAt := time.Date(2024, 1, 15, 20, 0, 0, 0, time.UTC)
					location := "Berlin"
					videoURL := "https://example.com/stream"
					participantsCount := 42
					avatarURL := "https://example.com/avatar.jpg"
					socialIcon := "instagram"

					events := []models.EventEnriched{
						{
							Event: models.Event{
								ID:          1,
								Title:       "Test Event",
								Description: "A test event description",
								Location:    &location,
								VideoURL:    &videoURL,
								StartsAt:    &startsAt,
							},
							ParticipantsCount: &participantsCount,
							IsSubscribed:      true,
							DJs: []models.DJ{
								{
									ID:        1,
									StageName: "DJ Test",
									AvatarURL: &avatarURL,
									Socials: []models.Social{
										{
											Name: "Instagram",
											URL:  "https://instagram.com/djtest",
											Icon: &socialIcon,
										},
									},
								},
							},
						},
					}

					mock.EXPECT().
						ListEvents(gomock.Any(), int64(999), 30, 0).
						Return(events, nil)
					return mock
				},
			},
			setupCtx: func() context.Context {
				ginCtx := &gin.Context{}
				ginCtx.Set(models.ContextKeyTGUserID, int64(999))
				return ginCtx
			},
			request: publicapi.EventsRequestObject{
				Params: publicapi.EventsParams{},
			},
			expected: publicapi.Events200JSONResponse{
				{
					Id:                1,
					Title:             "Test Event",
					Description:       "A test event description",
					Location:          func() *string { s := "Berlin"; return &s }(),
					VideoUrl:          func() *string { s := "https://example.com/stream"; return &s }(),
					StartsAt:          func() *time.Time { t := time.Date(2024, 1, 15, 20, 0, 0, 0, time.UTC); return &t }(),
					IsSubscribed:      true,
					ParticipantsCount: func() *int { c := 42; return &c }(),
					Djs: []publicapi.DJ{
						{
							StageName: "DJ Test",
							AvatarUrl: func() *string { s := "https://example.com/avatar.jpg"; return &s }(),
							Socials: []publicapi.Social{
								{
									Name: "Instagram",
									Url:  "https://instagram.com/djtest",
									Icon: func() *string { s := "instagram"; return &s }(),
								},
							},
						},
					},
				},
			},
			expectedErr: assert.NoError,
		},
		{
			name: "success with multiple events",
			fields: fields{
				auth:              mocks.NewMockauth,
				usersRepo:         mocks.NewMockusersRepo,
				subscriptionsRepo: mocks.NewMocksubscriptionsRepo,
				eventsRepo: func(ctrl *gomock.Controller) *mocks.MockeventsRepo {
					mock := mocks.NewMockeventsRepo(ctrl)

					startsAt1 := time.Date(2024, 1, 15, 20, 0, 0, 0, time.UTC)
					startsAt2 := time.Date(2024, 1, 20, 22, 0, 0, 0, time.UTC)

					events := []models.EventEnriched{
						{
							Event: models.Event{
								ID:          1,
								Title:       "First Event",
								Description: "First event description",
								StartsAt:    &startsAt1,
							},
							IsSubscribed: true,
							DJs:          []models.DJ{},
						},
						{
							Event: models.Event{
								ID:          2,
								Title:       "Second Event",
								Description: "Second event description",
								StartsAt:    &startsAt2,
							},
							IsSubscribed: false,
							DJs:          []models.DJ{},
						},
					}

					mock.EXPECT().
						ListEvents(gomock.Any(), int64(555), 30, 0).
						Return(events, nil)
					return mock
				},
			},
			setupCtx: func() context.Context {
				ginCtx := &gin.Context{}
				ginCtx.Set(models.ContextKeyTGUserID, int64(555))
				return ginCtx
			},
			request: publicapi.EventsRequestObject{
				Params: publicapi.EventsParams{},
			},
			expected: publicapi.Events200JSONResponse{
				{
					Id:                1,
					Title:             "First Event",
					Description:       "First event description",
					Location:          nil,
					VideoUrl:          nil,
					StartsAt:          func() *time.Time { t := time.Date(2024, 1, 15, 20, 0, 0, 0, time.UTC); return &t }(),
					IsSubscribed:      true,
					ParticipantsCount: nil,
					Djs:               []publicapi.DJ{},
				},
				{
					Id:                2,
					Title:             "Second Event",
					Description:       "Second event description",
					Location:          nil,
					VideoUrl:          nil,
					StartsAt:          func() *time.Time { t := time.Date(2024, 1, 20, 22, 0, 0, 0, time.UTC); return &t }(),
					IsSubscribed:      false,
					ParticipantsCount: nil,
					Djs:               []publicapi.DJ{},
				},
			},
			expectedErr: assert.NoError,
		},
		{
			name: "repository error",
			fields: fields{
				auth:              mocks.NewMockauth,
				usersRepo:         mocks.NewMockusersRepo,
				subscriptionsRepo: mocks.NewMocksubscriptionsRepo,
				eventsRepo: func(ctrl *gomock.Controller) *mocks.MockeventsRepo {
					mock := mocks.NewMockeventsRepo(ctrl)
					mock.EXPECT().
						ListEvents(gomock.Any(), int64(777), 30, 0).
						Return(nil, assert.AnError)
					return mock
				},
			},
			setupCtx: func() context.Context {
				ginCtx := &gin.Context{}
				ginCtx.Set(models.ContextKeyTGUserID, int64(777))
				return ginCtx
			},
			request: publicapi.EventsRequestObject{
				Params: publicapi.EventsParams{},
			},
			expected:    nil,
			expectedErr: assert.Error,
		},
		{
			name: "success with large limit",
			fields: fields{
				auth:              mocks.NewMockauth,
				usersRepo:         mocks.NewMockusersRepo,
				subscriptionsRepo: mocks.NewMocksubscriptionsRepo,
				eventsRepo: func(ctrl *gomock.Controller) *mocks.MockeventsRepo {
					mock := mocks.NewMockeventsRepo(ctrl)
					mock.EXPECT().
						ListEvents(gomock.Any(), int64(888), 100, 50).
						Return([]models.EventEnriched{}, nil)
					return mock
				},
			},
			setupCtx: func() context.Context {
				ginCtx := &gin.Context{}
				ginCtx.Set(models.ContextKeyTGUserID, int64(888))
				return ginCtx
			},
			request: publicapi.EventsRequestObject{
				Params: publicapi.EventsParams{
					Limit:  func() *int { l := 100; return &l }(),
					Offset: func() *int { o := 50; return &o }(),
				},
			},
			expected:    publicapi.Events200JSONResponse{},
			expectedErr: assert.NoError,
		},
		{
			name: "success with event having multiple DJs",
			fields: fields{
				auth:              mocks.NewMockauth,
				usersRepo:         mocks.NewMockusersRepo,
				subscriptionsRepo: mocks.NewMocksubscriptionsRepo,
				eventsRepo: func(ctrl *gomock.Controller) *mocks.MockeventsRepo {
					mock := mocks.NewMockeventsRepo(ctrl)

					startsAt := time.Date(2024, 2, 10, 21, 0, 0, 0, time.UTC)
					avatar1 := "https://example.com/dj1.jpg"
					avatar2 := "https://example.com/dj2.jpg"
					icon1 := "soundcloud"
					icon2 := "spotify"

					events := []models.EventEnriched{
						{
							Event: models.Event{
								ID:          5,
								Title:       "Multi DJ Event",
								Description: "Event with multiple DJs",
								StartsAt:    &startsAt,
							},
							IsSubscribed: false,
							DJs: []models.DJ{
								{
									ID:        1,
									StageName: "DJ Alpha",
									AvatarURL: &avatar1,
									Socials: []models.Social{
										{
											Name: "SoundCloud",
											URL:  "https://soundcloud.com/djalpha",
											Icon: &icon1,
										},
									},
								},
								{
									ID:        2,
									StageName: "DJ Beta",
									AvatarURL: &avatar2,
									Socials: []models.Social{
										{
											Name: "Spotify",
											URL:  "https://spotify.com/djbeta",
											Icon: &icon2,
										},
									},
								},
							},
						},
					}

					mock.EXPECT().
						ListEvents(gomock.Any(), int64(333), 30, 0).
						Return(events, nil)
					return mock
				},
			},
			setupCtx: func() context.Context {
				ginCtx := &gin.Context{}
				ginCtx.Set(models.ContextKeyTGUserID, int64(333))
				return ginCtx
			},
			request: publicapi.EventsRequestObject{
				Params: publicapi.EventsParams{},
			},
			expected: publicapi.Events200JSONResponse{
				{
					Id:                5,
					Title:             "Multi DJ Event",
					Description:       "Event with multiple DJs",
					Location:          nil,
					VideoUrl:          nil,
					StartsAt:          func() *time.Time { t := time.Date(2024, 2, 10, 21, 0, 0, 0, time.UTC); return &t }(),
					IsSubscribed:      false,
					ParticipantsCount: nil,
					Djs: []publicapi.DJ{
						{
							StageName: "DJ Alpha",
							AvatarUrl: func() *string { s := "https://example.com/dj1.jpg"; return &s }(),
							Socials: []publicapi.Social{
								{
									Name: "SoundCloud",
									Url:  "https://soundcloud.com/djalpha",
									Icon: func() *string { s := "soundcloud"; return &s }(),
								},
							},
						},
						{
							StageName: "DJ Beta",
							AvatarUrl: func() *string { s := "https://example.com/dj2.jpg"; return &s }(),
							Socials: []publicapi.Social{
								{
									Name: "Spotify",
									Url:  "https://spotify.com/djbeta",
									Icon: func() *string { s := "spotify"; return &s }(),
								},
							},
						},
					},
				},
			},
			expectedErr: assert.NoError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			ctx := tt.setupCtx()

			srv := api.NewAPI(
				tt.fields.auth(ctrl),
				tt.fields.usersRepo(ctrl),
				tt.fields.subscriptionsRepo(ctrl),
				tt.fields.eventsRepo(ctrl),
				zap.NewNop(),
			)

			result, err := srv.Events(ctx, tt.request)
			assert.Equal(t, tt.expected, result)
			tt.expectedErr(t, err)
		})
	}
}

func TestEvents_PanicOnMissingUserID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	srv := api.NewAPI(
		mocks.NewMockauth(ctrl),
		mocks.NewMockusersRepo(ctrl),
		mocks.NewMocksubscriptionsRepo(ctrl),
		mocks.NewMockeventsRepo(ctrl),
		zap.NewNop(),
	)

	// Context without user ID should cause panic in MustTGUserID
	ctx := &gin.Context{}

	assert.Panics(t, func() {
		_, _ = srv.Events(ctx, publicapi.EventsRequestObject{})
	})
}
