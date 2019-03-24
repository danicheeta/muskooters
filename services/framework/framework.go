package framework

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"muskooters/services/config"
	"muskooters/services/initializer"
)

var (
	all []Routes
)

type Routes interface {
	Routes(*gin.Engine)
}

type initer struct {
}

func (i *initer) Initialize() func() {
	port := config.MustString("PORT")
	e := gin.New()
	e.Use(Recovery)

	for i := range all {
		all[i].Routes(e)
	}

	go func() {
		err := e.Run(port)
		logrus.Errorln("[framework]", err)
	}()

	return nil
}

func Error(c *gin.Context, status int, err string) {
	c.JSON(status, struct {
		Error string
	}{err})
}

// Register a new controller class
func Register(c ...Routes) {
	all = append(all, c...)
}

func init() {
	initializer.Register(&initer{})
}
