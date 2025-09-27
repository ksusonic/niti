package auth_test

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/ksusonic/niti/backend/internal/api/server/middleware/auth"
	"github.com/ksusonic/niti/backend/internal/api/server/middleware/auth/mocks"
	"github.com/ksusonic/niti/backend/pgk/publicapi"
	"go.uber.org/mock/gomock"
)

func TestAuthMw_NoAuthRequired(t *testing.T) {
	gin.SetMode(gin.TestMode)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	authDeps := mocks.NewMockAuthDeps(ctrl)
	// No expectations - ValidateAccessToken should not be called

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = httptest.NewRequest("GET", "/", nil)

	middleware := auth.AuthMw(authDeps)
	middleware(c)

	if c.IsAborted() {
		t.Error("middleware should not abort when auth is not required")
	}
}

func TestAuthMw_MissingAuthHeader(t *testing.T) {
	gin.SetMode(gin.TestMode)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	authDeps := mocks.NewMockAuthDeps(ctrl)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = httptest.NewRequest("GET", "/", nil)
	c.Set(publicapi.BearerAuthScopes, []string{})

	middleware := auth.AuthMw(authDeps)
	middleware(c)

	if !c.IsAborted() {
		t.Error("middleware should abort when auth header is missing")
	}

	if w.Code != http.StatusUnauthorized {
		t.Errorf("expected status %d, got %d", http.StatusUnauthorized, w.Code)
	}
}

func TestAuthMw_InvalidAuthHeader(t *testing.T) {
	gin.SetMode(gin.TestMode)

	tests := []struct {
		name   string
		header string
	}{
		{"no bearer prefix", "token123"},
		{"wrong prefix", "Basic token123"},
		{"empty token", "Bearer "},
		{"no token", "Bearer"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			authDeps := mocks.NewMockAuthDeps(ctrl)

			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)
			c.Request = httptest.NewRequest("GET", "/", nil)
			c.Request.Header.Set("Authorization", tt.header)
			c.Set(publicapi.BearerAuthScopes, []string{})

			middleware := auth.AuthMw(authDeps)
			middleware(c)

			if !c.IsAborted() {
				t.Error("middleware should abort with invalid auth header")
			}

			if w.Code != http.StatusUnauthorized {
				t.Errorf("expected status %d, got %d", http.StatusUnauthorized, w.Code)
			}
		})
	}
}

func TestAuthMw_InvalidToken(t *testing.T) {
	gin.SetMode(gin.TestMode)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	authDeps := mocks.NewMockAuthDeps(ctrl)
	authDeps.EXPECT().ValidateAccessToken("invalid-token").Return(int64(0), errors.New("invalid token"))

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = httptest.NewRequest("GET", "/", nil)
	c.Request.Header.Set("Authorization", "Bearer invalid-token")
	c.Set(publicapi.BearerAuthScopes, []string{})

	middleware := auth.AuthMw(authDeps)
	middleware(c)

	if !c.IsAborted() {
		t.Error("middleware should abort with invalid token")
	}

	if w.Code != http.StatusUnauthorized {
		t.Errorf("expected status %d, got %d", http.StatusUnauthorized, w.Code)
	}
}

func TestAuthMw_ValidToken(t *testing.T) {
	gin.SetMode(gin.TestMode)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	expectedUserID := int64(123)
	authDeps := mocks.NewMockAuthDeps(ctrl)
	authDeps.EXPECT().ValidateAccessToken("valid-token").Return(expectedUserID, nil)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = httptest.NewRequest("GET", "/", nil)
	c.Request.Header.Set("Authorization", "Bearer valid-token")
	c.Set(publicapi.BearerAuthScopes, []string{})

	middleware := auth.AuthMw(authDeps)
	middleware(c)

	if c.IsAborted() {
		t.Error("middleware should not abort with valid token")
	}

	userID := auth.UserIDFromContext(c)
	if userID != expectedUserID {
		t.Errorf("expected user ID %d, got %d", expectedUserID, userID)
	}
}
