package routes

import (
	"github.com/gin-gonic/gin"
	"fmt"
	"muskooters/user"
	"net/http"
	"muskooters/user/middleware"
)

func login(c *gin.Context) {
	var u user.User
	if err := c.Bind(&u); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	u, err := user.GetByName(u.Username)
	if err != nil {
		c.JSON(http.StatusNotFound, nil)
		return
	}

	t := middleware.GenToken(string(u.Role))
	c.JSON(http.StatusOK, struct {
		Token string
	}{t})
}
