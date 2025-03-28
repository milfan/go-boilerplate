package api_rest

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
	rest_routes "github.com/milfan/go-boilerplate/api/rest/routes"
	"github.com/milfan/go-boilerplate/configs/config"
	"github.com/milfan/go-boilerplate/configs/middleware"
	config_postgres "github.com/milfan/go-boilerplate/configs/postgres"
	api_controllers "github.com/milfan/go-boilerplate/internal/api/controllers"
	pkg_response "github.com/milfan/go-boilerplate/pkg/response"
	"github.com/sirupsen/logrus"
)

type Server struct {
	httpServer *http.Server
}

func New(
	server *gin.Engine,
	httpConf config.HttpConfig,
	postgresConn config_postgres.Postgres,
	logger *logrus.Logger,
) *Server {

	pkgResponse := pkg_response.New(logger)
	httpTimeout := time.Duration(httpConf.Timeout()) * time.Second
	apiControllers := api_controllers.LoadControllers(pkgResponse, postgresConn, logger)

	server.Use(middleware.CORSMiddleware())
	server.Use(middleware.RequestTimeoutMiddleware(
		httpTimeout,
		pkgResponse,
	))
	server.Use(middleware.GatherRequestData(pkgResponse, logger))

	rest_routes.DefaultRoute(server)
	rest_routes.WebRouteV1(server, apiControllers)
	rest_routes.MobileRouteV1(server, apiControllers)

	serv := Server{
		httpServer: &http.Server{
			Addr:    fmt.Sprintf(":%s", httpConf.Port()),
			Handler: server.Handler(),
		},
	}

	return &serv
}

func (s *Server) Start() {
	go func() {
		// service connections
		if err := s.httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	gracefullShutdown(s.httpServer)
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
