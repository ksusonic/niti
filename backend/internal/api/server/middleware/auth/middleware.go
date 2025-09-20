package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ksusonic/niti/backend/pgk/genapi"
)

type AuthDeps interface {
	ValidateAccessToken(tokenStr string) (int64, error)
}

func AuthMw(auth AuthDeps) genapi.MiddlewareFunc {
	return func(c *gin.Context) {
		// check only auth-enabled endpoints
		if _, ok := c.Get(genapi.BearerAuthScopes); ok {
			token := c.GetHeader("Authorization")
			if token == "" {
				c.AbortWithStatusJSON(http.StatusUnauthorized, genapi.Error{Message: "unauthorized"})
				return
			}

			userID, err := auth.ValidateAccessToken(token)
			if err != nil {
				c.AbortWithStatusJSON(http.StatusUnauthorized, genapi.Error{Message: "invalid token"})
				return
			}

			SetUserIDInContext(c, userID)
		}
	}
}
