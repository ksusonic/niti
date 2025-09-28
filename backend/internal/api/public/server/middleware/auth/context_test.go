package auth_test

import (
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/ksusonic/niti/backend/internal/api/public/server/middleware/auth"
)

func TestSetUserIDInContext(t *testing.T) {
	gin.SetMode(gin.TestMode)

	c, _ := gin.CreateTestContext(nil)
	userID := int64(123)

	auth.SetUserIDInContext(c, userID)

	retrievedUserID := auth.UserIDFromContext(c)
	if retrievedUserID != userID {
		t.Errorf("expected user ID %d, got %d", userID, retrievedUserID)
	}
}

func TestUserIDFromContext_Panic(t *testing.T) {
	gin.SetMode(gin.TestMode)

	c, _ := gin.CreateTestContext(nil)

	defer func() {
		if r := recover(); r == nil {
			t.Error("expected panic when user ID not in context")
		}
	}()

	auth.UserIDFromContext(c)
}
