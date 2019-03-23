package routes

import (
	"github.com/gin-gonic/gin"
	"muskooters/services/framework"
	"muskooters/user/middleware"
	"muskooters/user"
)

type Route struct {}

func (Route) Routes(e *gin.Engine) {
	g := e.Group("user")
	g.POST("login", login)
}

func init(){
	framework.Register(Route{})
}