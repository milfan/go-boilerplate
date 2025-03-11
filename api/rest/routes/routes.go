package rest_routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func DefaultRoute(route *gin.Engine) {
	route.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, "online!")
	})
}
