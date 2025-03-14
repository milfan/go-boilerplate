package api_web_usecases

import "github.com/milfan/go-boilerplate/internal/api/repositories"

type (
	WebUsecases struct {
		EmployeeUsecases IEmployeeUsecase
	}
)

func RegisterWebUsecases(
	repo repositories.Repositories,
) WebUsecases {
	return WebUsecases{
		EmployeeUsecases: newEmployeeUsecase(repo.EmployeeRepositories),
	}
}
