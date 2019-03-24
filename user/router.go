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
	e.POST("login/user", login)
	e.GET("login/scooter", scooterLogin)
	e.POST("hunter", FetchToken, Auth(Zeus), register)
}

func scooterLogin(c *gin.Context) {
	t := jwt.GenToken(string(Scooter))
	c.JSON(http.StatusOK, struct {
		Token string
	}{t})
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
		framework.Error(c, http.StatusUnauthorized, "user not found")
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(dbuser.Password), []byte(u.Password)); err != nil {
		logrus.Errorln("login route:", err)
		framework.Error(c, http.StatusUnauthorized, "invalid password")
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

	if !assert.String(u.Username, u.Password) || !ensureRole(u.Role) {
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

func ensureRole(r Role) bool {
	for _, role := range allRoles {
		if role == r {
			return true
		}
	}

	return false
}

type Route struct{}

func init() {
	framework.Register(Route{})
}
