package api_controllers

import (
	config_postgres "github.com/milfan/go-boilerplate/configs/postgres"
	web_controller "github.com/milfan/go-boilerplate/internal/api/controllers/web"
	"github.com/milfan/go-boilerplate/internal/api/repositories"
	api_usecases "github.com/milfan/go-boilerplate/internal/api/usecases"
	pkg_log "github.com/milfan/go-boilerplate/pkg/log"
	pkg_response "github.com/milfan/go-boilerplate/pkg/response"
)

type (
	Controllers struct {
		WebControllers web_controller.WebControllers
	}
)

func LoadControllers(
	pkgResponse pkg_response.IResponse,
	conn config_postgres.Postgres,
	appLogger *pkg_log.AppLogger,
) Controllers {

	loadRepositories := repositories.LoadRepositories(conn, appLogger)
	loadUsecases := api_usecases.LoadUsecases(loadRepositories, appLogger)

	return Controllers{
		WebControllers: web_controller.RegisterWebController(pkgResponse, loadUsecases),
	}
}
