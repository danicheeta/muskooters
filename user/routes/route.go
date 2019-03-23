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
	g.Use(middleware.FetchToken)
	g.POST("register", middleware.Auth(user.Zeus), register)
}

func init(){
	framework.Register(Route{})
}