package station

import (
	"muskooters/services/framework"
	"muskooters/user"
	"muskooters/user/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Route struct{}

func (Route) Routes(e *gin.Engine) {
	g := e.Group("scooter")
	g.Use(middleware.FetchToken)
	g.GET(":id", middleware.Auth(user.Zeus), getScooterState)
	g.POST("register", middleware.Auth(user.Zeus), registerScooter)
}

func registerScooter(c *gin.Context) {
	s := NewScooter()
	c.JSON(http.StatusOK, s)
}

func getScooterState(c *gin.Context) {
	id := c.Param("id")
	s, err := GetScooterState(id)
	if err != nil {
		c.JSON(http.StatusNotFound, err)
		return
	}

	c.JSON(http.StatusOK, s)
}

func init() {
	framework.Register(Route{})
}
