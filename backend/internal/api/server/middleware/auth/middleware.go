package auth

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/ksusonic/niti/backend/pgk/genapi"
)

func AuthMw(auth AuthDeps) genapi.MiddlewareFunc {
	return func(c *gin.Context) {
		// check only auth-enabled endpoints
		if _, ok := c.Get(genapi.BearerAuthScopes); ok {
			header := c.GetHeader("Authorization")
			if header == "" {
				c.AbortWithStatusJSON(http.StatusUnauthorized, genapi.Error{Message: "unauthorized"})
				return
			}

			parts := strings.SplitN(header, " ", 2)
			if len(parts) != 2 || parts[0] != "Bearer" || parts[1] == "" {
				c.AbortWithStatusJSON(http.StatusUnauthorized, genapi.Error{Message: "unauthorized"})
				return
			}

			userID, err := auth.ValidateAccessToken(parts[1])
			if err != nil {
				c.AbortWithStatusJSON(http.StatusUnauthorized, genapi.Error{Message: "invalid token"})
				return
			}

			SetUserIDInContext(c, userID)
		}
	}
}
