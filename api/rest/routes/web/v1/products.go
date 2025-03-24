package web_routes_v1

import (
	"github.com/gin-gonic/gin"
	api_web_controller "github.com/milfan/go-boilerplate/internal/api/controllers/web"
)

func ProductsWebV1Route(routeGroup *gin.RouterGroup, ctrl api_web_controller.WebControllers) {
	routes := routeGroup.Group("/products")
	routes.POST("add", ctrl.ProductController.Add)
	routes.GET("list", ctrl.ProductController.List)
}
