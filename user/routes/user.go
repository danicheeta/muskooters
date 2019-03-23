package routes

import (
	"github.com/gin-gonic/gin"
	"muskooters/user"
	"net/http"
	"muskooters/user/middleware"
	"golang.org/x/crypto/bcrypt"
	"muskooters/services/assert"
)

func login(c *gin.Context) {
	var u user.User
	if err := c.Bind(&u); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	dbuser, err := user.GetByName(u.Username)
	if err != nil {
		c.JSON(http.StatusNotFound, nil)
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(dbuser.Password), []byte(u.Password)); err != nil {
		c.JSON(http.StatusBadRequest, struct {
			Error string
		}{"invalid password"})
		return
	}

	t := middleware.GenToken(string(u.Role))
	c.JSON(http.StatusOK, struct {
		Token string
	}{t})
}

func register(c *gin.Context) {
	var u user.User
	if err := c.Bind(&u); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	pwd, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	assert.Nil(err)

	err = user.Add(u.Username, string(pwd), u.Role)
	assert.Nil(err)

	t := middleware.GenToken(string(u.Role))
	c.JSON(http.StatusOK, struct {
		Token string
	}{t})
}
