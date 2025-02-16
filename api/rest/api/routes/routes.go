package rest_api_routes

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	route_v1 "github.com/milfan/golang-gin/api/rest/api/routes/v1"
	conf_app "github.com/milfan/golang-gin/configs/app_conf"
	conf_middleware "github.com/milfan/golang-gin/configs/middleware"
	"github.com/milfan/golang-gin/internal/application/controllers"
	"github.com/sirupsen/logrus"
)

func DefaultRoute(route *gin.Engine) {
	route.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, "online!")
	})
}

func V1Route(
	route *gin.Engine,
	logger *logrus.Logger,
	httpConf conf_app.HttpConfig,
	controllers controllers.Controllers,
) {
	httpTimeout := time.Duration(httpConf.GetTimeout()) * time.Second

	mainRoute := route.Group("/hris")

	mainRoute.Use(conf_middleware.RequestTimeoutMiddleware(httpTimeout, logger))

	route_v1.MobileV1Route(mainRoute, controllers)
	route_v1.WebV1Route(mainRoute, controllers)
}
