package rest_routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	web_routes_v1 "github.com/milfan/go-boilerplate/api/rest/routes/web/v1"
	api_controllers "github.com/milfan/go-boilerplate/internal/api/controllers"
)

func DefaultRoute(route *gin.Engine) {
	route.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, "online!")
	})
}

func WebRouteV1(
	routeGroup *gin.Engine,
	ctrl api_controllers.Controllers,
) {
	webRouteV1 := routeGroup.Group("v1/web")
	web_routes_v1.EmployeeMobileV1Route(webRouteV1, ctrl.WebControllers)
}
