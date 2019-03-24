package framework

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Recovery(c *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			Error(c, http.StatusInternalServerError, "ITS NOT OK!")
		}
	}()
	c.Next()
}
