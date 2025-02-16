package route_v1

import (
	"github.com/gin-gonic/gin"
	"github.com/milfan/golang-gin/internal/application/controllers"
)

func WebV1Route(
	routeGroup *gin.RouterGroup,
	ctrl controllers.Controllers,
) {
	platformRoute := routeGroup.Group("v1//web")

	EmployeeWebV1Route(platformRoute, ctrl.V1Controller.MobileControllers)
}

func EmployeeWebV1Route(routeGroup *gin.RouterGroup, ctrl controllers.MobileV1Controllers) {
	routes := routeGroup.Group("/employee")
	routes.POST("register", ctrl.EmployeeController.Register)
}
