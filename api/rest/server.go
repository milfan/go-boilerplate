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
	config_postgres "github.com/milfan/go-boilerplate/configs/postgres"
	api_controllers "github.com/milfan/go-boilerplate/internal/api/controllers"
	pkg_response "github.com/milfan/go-boilerplate/pkg/response"
	"github.com/sirupsen/logrus"
)

func NewServer(
	server *gin.Engine,
	httpConf config.HttpConfig,
	postgresConn config_postgres.Postgres,
	logger *logrus.Logger,
) *http.Server {
	pkgResponse := pkg_response.New()

	apiControllers := api_controllers.LoadControllers(pkgResponse, postgresConn)

	rest_routes.DefaultRoute(server)
	rest_routes.WebRouteV1(server, apiControllers)

	return &http.Server{
		Addr:    fmt.Sprintf(":%s", httpConf.Port()),
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
