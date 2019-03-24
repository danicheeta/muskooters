package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/kataras/iris/core/errors"
	"muskooters/user"
	"net/http"
)

const contextRole = "role"

// TODO recovery
// get role from token and set it in context
// tokens are based on jwt on Authorization header
func FetchToken(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		c.AbortWithError(http.StatusBadRequest, errors.New("invalid token"))
		return
	}

	if len(authHeader) < 7 || authHeader[:6] != "bearer" {
		c.AbortWithError(http.StatusBadRequest, errors.New("invalid token"))
		return
	}

	role, err := getRoleFromToken(authHeader[7:])
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, errors.New("invalid token"))
		return
	}

	c.Set(contextRole, role)
	c.Next()
}

// validates given role with current user's role
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
