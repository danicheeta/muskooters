package user

import (
	"muskooters/services/assert"
	"muskooters/services/framework"
	"muskooters/user/jwt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

func (Route) Routes(e *gin.Engine) {
	g := e.Group("user")
	g.POST("login", login)
	g.Use(FetchToken)
	g.POST("register", Auth(Zeus), register)
}

func login(c *gin.Context) {
	var u User
	if err := c.Bind(&u); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	dbuser, err := GetByName(u.Username)
	if err != nil {
		logrus.Errorln("login route:", err)
		framework.Error(c, http.StatusNotFound, "user not found")
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(dbuser.Password), []byte(u.Password)); err != nil {
		logrus.Errorln("login route:", err)
		framework.Error(c, http.StatusNotFound, "invalid password")
		return
	}

	t := jwt.GenToken(string(dbuser.Role))
	c.JSON(http.StatusOK, struct {
		Token string
	}{t})
}

func register(c *gin.Context) {
	var u User
	if err := c.Bind(&u); err != nil {
		logrus.Errorln("user register route:", err)
		framework.Error(c, http.StatusBadRequest, "invalid payload")
		return
	}

	pwd, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	assert.Nil(err)

	err = Add(u.Username, string(pwd), u.Role)
	assert.Nil(err)

	t := jwt.GenToken(string(u.Role))
	c.JSON(http.StatusOK, struct {
		Token string
	}{t})
}

type Route struct{}

func init() {
	framework.Register(Route{})
}
