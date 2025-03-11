package web_routes_v1

import (
	"github.com/gin-gonic/gin"
	api_controllers "github.com/milfan/go-boilerplate/internal/api/controllers"
	web_controller "github.com/milfan/go-boilerplate/internal/api/controllers/web"
)

func WebRouteV1(
	routeGroup *gin.Engine,
	ctrl api_controllers.Controllers,
) {
	platformRoute := routeGroup.Group("v1/web")

	employeeMobileV1Route(platformRoute, ctrl.WebControllers)
}

func employeeMobileV1Route(routeGroup *gin.RouterGroup, ctrl web_controller.WebControllers) {
	routes := routeGroup.Group("/employee")
	routes.POST("register", ctrl.EmployeeController.Register)
}
