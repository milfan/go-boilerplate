package route_v1

import (
	"github.com/gin-gonic/gin"
	"github.com/milfan/golang-gin/internal/application/controllers"
)

func MobileV1Route(
	routeGroup *gin.RouterGroup,
	ctrl controllers.Controllers,
) {
	platformRoute := routeGroup.Group("v1//mobile")

	EmployeeMobileV1Route(platformRoute, ctrl.V1Controller.MobileControllers)
}

func EmployeeMobileV1Route(routeGroup *gin.RouterGroup, ctrl controllers.MobileV1Controllers) {
	routes := routeGroup.Group("/employee")
	routes.POST("register", ctrl.EmployeeController.Register)
}
