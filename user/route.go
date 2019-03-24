package user

import (
	"github.com/gin-gonic/gin"
	"muskooters/services/framework"
	"muskooters/user/middleware"
	"golang.org/x/crypto/bcrypt"
	"muskooters/services/assert"
	"net/http"
)

func (Route) Routes(e *gin.Engine) {
	g := e.Group("user")
	g.POST("login", login)
	g.Use(middleware.FetchToken)
	g.POST("register", middleware.Auth(Zeus), register)
}

func login(c *gin.Context) {
	var u User
	if err := c.Bind(&u); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	dbuser, err := GetByName(u.Username)
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

	t := middleware.GenToken(string(dbuser.Role))
	c.JSON(http.StatusOK, struct {
		Token string
	}{t})
}

func register(c *gin.Context) {
	var u User
	if err := c.Bind(&u); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	pwd, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	assert.Nil(err)

	err = Add(u.Username, string(pwd), u.Role)
	assert.Nil(err)

	t := middleware.GenToken(string(u.Role))
	c.JSON(http.StatusOK, struct {
		Token string
	}{t})
}

type Route struct{}

func init() {
	framework.Register(Route{})
}
