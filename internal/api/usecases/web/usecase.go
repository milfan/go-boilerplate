package api_web_usecases

import (
	"github.com/milfan/go-boilerplate/internal/api/repositories"
	pkg_log "github.com/milfan/go-boilerplate/pkg/log"
)

type (
	WebUsecases struct {
		EmployeeUsecases IEmployeeUsecase
	}
)

func RegisterWebUsecases(
	repo repositories.Repositories,
	appLogger *pkg_log.AppLogger,
) WebUsecases {
	return WebUsecases{
		EmployeeUsecases: newEmployeeUsecase(appLogger, repo.EmployeeRepositories),
	}
}
