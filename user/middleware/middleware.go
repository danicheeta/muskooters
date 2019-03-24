package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/kataras/iris/core/errors"
	"muskooters/user"
	"net/http"
)

const contextRole = "role"

func FetchToken(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		c.AbortWithError(http.StatusBadRequest, errors.New("invalid token"))
		return
	}

	role, err := getRoleFromToken(authHeader)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, errors.New("invalid token"))
		return
	}

	c.Set(contextRole, role)
	c.Next()
}

func Auth(role user.Role) gin.HandlerFunc {
	return func(c *gin.Context) {
		r, ok := c.Get(contextRole)
		if !ok {
			c.AbortWithError(http.StatusBadRequest, errors.New("invalid token"))
			return
		}

		if r.(user.Role) != role {
			c.AbortWithError(http.StatusBadRequest, errors.New("you don't have permission"))
			return
		}
	}
}
