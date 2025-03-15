package api_web_usecases

import (
	"github.com/milfan/go-boilerplate/internal/api/repositories"
	"github.com/sirupsen/logrus"
)

type (
	WebUsecases struct {
		EmployeeUsecases IEmployeeUsecase
	}
)

func RegisterWebUsecases(
	repo repositories.Repositories,
	logger *logrus.Logger,
) WebUsecases {
	return WebUsecases{
		EmployeeUsecases: newEmployeeUsecase(logger, repo.EmployeeRepositories),
	}
}
