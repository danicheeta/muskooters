package station

import (
	"muskooters/services/framework"
	"muskooters/user"
	"net/http"

	"errors"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"muskooters/services/assert"
)

type Route struct{}

func (Route) Routes(e *gin.Engine) {
	g := e.Group("scooter")
	g.Use(user.FetchToken)
	g.GET(":id/state", user.Auth(user.Zeus), getScooterState)
	g.POST(":id/state", setScooterState)
	g.POST("", user.Auth(user.Zeus), registerScooter)
}

// create new scooter only with admin permission
func registerScooter(c *gin.Context) {
	panic("YAAA ALI")
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

func setScooterState(c *gin.Context) {
	var payload struct {
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
