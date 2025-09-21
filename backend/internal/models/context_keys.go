package models

import (
	"context"

	"github.com/gin-gonic/gin"
)

const (
	ContextKeyTGUserID = "user_id"
)

func MustTGUserID(ctx context.Context) int64 {
	// hack: oapi-codegen casts *gin.Context to context.Context
	userID, ok := ctx.(*gin.Context).Get(ContextKeyTGUserID)
	if !ok {
		panic("no user_id in context")
	}

	return userID.(int64)
}
