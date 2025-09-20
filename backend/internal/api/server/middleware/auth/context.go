package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/ksusonic/niti/backend/internal/models"
)

func SetUserIDInContext(c *gin.Context, userID int64) {
	c.Set(models.ContextKeyTGUserID, userID)
}

func UserIDFromContext(c *gin.Context) int64 {
	return c.MustGet(models.ContextKeyTGUserID).(int64)
}
