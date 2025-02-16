package rest_api

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	rest_api_routes "github.com/milfan/golang-gin/api/rest/api/routes"
	conf_app "github.com/milfan/golang-gin/configs/app_conf"
	"github.com/milfan/golang-gin/internal/application/controllers"
	"github.com/milfan/golang-gin/internal/pkg/database/postgres"
	"github.com/sirupsen/logrus"
)

func NewServer(
	server *gin.Engine,
	httpConf conf_app.HttpConfig,
	postgresConn postgres.Postgres,
	logger *logrus.Logger,
) *http.Server {

	controllers := controllers.LoadControllers()
	rest_api_routes.DefaultRoute(server)
	rest_api_routes.V1Route(server, logger, httpConf, controllers)

	return &http.Server{
		Addr:    fmt.Sprintf(":%s", httpConf.GetPort()),
		Handler: server.Handler(),
	}
}

func StartServer(srv *http.Server) {
	go func() {
		// service connections
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	gracefullShutdown(srv)
}

func gracefullShutdown(srv *http.Server) {
	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal, 1)
	// kill (no param) default send syscall.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall. SIGKILL but can"t be catch, so don't need add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}

	<-ctx.Done()
	log.Println("timeout of 5 seconds.")
	log.Println("Server exiting")
}
