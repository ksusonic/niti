package middleware

import "github.com/gin-gonic/gin"

const ContextKey = "user_id"

func SetUserIDInContext(c *gin.Context, userID int64) {
	c.Set(ContextKey, userID)
}

func UserIDFromContext(c *gin.Context) int64 {
	return c.MustGet(ContextKey).(int64)
}
