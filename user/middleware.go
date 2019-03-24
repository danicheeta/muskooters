package user

import (
	"github.com/gin-gonic/gin"
	"github.com/kataras/iris/core/errors"
	"muskooters/services/framework"
	"muskooters/user/middleware"
	"net/http"
)

const contextRole = "role"

// get role from token and set it in context
// tokens are based on jwt on Authorization header
func FetchToken(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" || len(authHeader) < 8 || authHeader[:6] != "bearer" {
		framework.Error(c, http.StatusUnauthorized, "invalid token")
		return
	}

	role, err := middleware.GetRoleFromToken(authHeader[7:])
	if err != nil {
		framework.Error(c, http.StatusUnauthorized, "invalid jwt token")
		return
	}

	c.Set(contextRole, Role(role.(string)))
	c.Next()
}

// validates given role with current user's role
func Auth(role Role) gin.HandlerFunc {
	return func(c *gin.Context) {
		r, ok := c.Get(contextRole)
		if !ok {
			c.AbortWithError(http.StatusBadRequest, errors.New("invalid token"))
			return
		}

		if r.(Role) != role {
			c.AbortWithError(http.StatusBadRequest, errors.New("you don't have permission"))
			return
		}
	}
}
