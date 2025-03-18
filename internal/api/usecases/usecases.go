package api_usecases

import (
	"github.com/milfan/go-boilerplate/internal/api/repositories"
	api_web_usecases "github.com/milfan/go-boilerplate/internal/api/usecases/web"
	pkg_log "github.com/milfan/go-boilerplate/pkg/log"
)

type (
	Usecases struct {
		WebUsecases api_web_usecases.WebUsecases
	}
)

func LoadUsecases(
	repo repositories.Repositories,
	appLogger *pkg_log.AppLogger,
) Usecases {
	return Usecases{
		WebUsecases: api_web_usecases.RegisterWebUsecases(repo, appLogger),
	}
}
