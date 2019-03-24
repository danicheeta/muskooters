package station

import (
	"muskooters/services/framework"
	"muskooters/user"
	"muskooters/user/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
	"errors"
	"muskooters/services/assert"
	"github.com/sirupsen/logrus"
)

type Route struct{}

func (Route) Routes(e *gin.Engine) {
	g := e.Group("scooter")
	g.Use(middleware.FetchToken)
	g.GET(":id/state", middleware.Auth(user.Zeus), getScooterState)
	g.POST(":id/state", setScooterState)
	g.POST("", middleware.Auth(user.Zeus), registerScooter)
}

func registerScooter(c *gin.Context) {
	s := NewScooter()
	c.JSON(http.StatusOK, s)
}

func getScooterState(c *gin.Context) {
	id := c.Param("id")
	scooter, err := GetScooter(id)
	if err != nil {
		logrus.Errorln("getScooterState route:", err)
		framework.Error(c, http.StatusNotFound, "scooter id not found")
		return
	}

	c.JSON(http.StatusOK, scooter)
}

// todo validation on id
func setScooterState(c *gin.Context) {
	var payload struct{
		State string
	}
	err := c.Bind(&payload)
	if err != nil {
		logrus.Errorln("setScooterState route:", err)
		framework.Error(c, http.StatusBadRequest, "invalid payload")
		return
	}

	id := c.Param("id")
	scooter, err := GetScooter(id)
	if err != nil {
		logrus.Errorln("setScooterState route:", err)
		framework.Error(c, http.StatusNotFound, "scooter id not found")
		return
	}

	userRole, ok := c.Get("role")
	assert.True(ok)

	state, ok := stringToState[payload.State]
	if !ok {
		msg := "invalid state name"
		logrus.Errorln("setScooterState route:", errors.New(msg))
		framework.Error(c, http.StatusNotFound, msg)
		return
	}

	if err = scooter.Transit(state, userRole.(user.Role)); err != nil {
		logrus.Errorln("setScooterState route:", err)
		framework.Error(c, http.StatusNotFound, err.Error())
		return
	}

	c.JSON(http.StatusOK, scooter)
}

func init() {
	framework.Register(Route{})
}
