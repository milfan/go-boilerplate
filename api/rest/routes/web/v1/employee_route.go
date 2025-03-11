package web_routes_v1

import (
	"github.com/gin-gonic/gin"
	web_controller "github.com/milfan/go-boilerplate/internal/api/controllers/web"
)

func EmployeeMobileV1Route(routeGroup *gin.RouterGroup, ctrl web_controller.WebControllers) {
	routes := routeGroup.Group("/employee")
	routes.POST("register", ctrl.EmployeeController.Register)
}
