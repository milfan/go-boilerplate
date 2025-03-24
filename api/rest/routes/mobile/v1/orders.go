package mobile_routes_v1

import (
	"github.com/gin-gonic/gin"
	api_mobile_controller "github.com/milfan/go-boilerplate/internal/api/controllers/mobile"
)

func OrdersMobileV1Route(routeGroup *gin.RouterGroup, ctrl api_mobile_controller.MobileControllers) {
	routes := routeGroup.Group("/orders")
	routes.POST("add", ctrl.OrderController.Add)
	routes.GET("list", ctrl.OrderController.List)
}
