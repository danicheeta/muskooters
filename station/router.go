package station

import (
	"muskooters/services/framework"
	"muskooters/user"
	"muskooters/user/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
	"errors"
	"muskooters/services/assert"
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
		c.JSON(http.StatusNotFound, err)
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
		c.JSON(http.StatusBadRequest, err)
		return
	}

	id := c.Param("id")
	scooter, err := GetScooter(id)
	if err != nil {
		c.JSON(http.StatusNotFound, err)
		return
	}

	userRole, ok := c.Get("role")
	assert.True(ok)

	state, ok := stringToState[payload.State]
	if !ok {
		c.JSON(http.StatusBadRequest, errors.New("invalid state name"))
		return
	}

	if err = scooter.Transit(state, userRole.(user.Role)); err != nil {
		c.JSON(http.StatusBadRequest, errors.New("invalid state name"))
		return
	}

	c.JSON(http.StatusOK, scooter)
}

func init() {
	framework.Register(Route{})
}
