package api_web_usecases

import (
	"context"

	"github.com/milfan/go-boilerplate/internal/api/entities"
	"github.com/milfan/go-boilerplate/internal/api/presenters/requests"
	"github.com/milfan/go-boilerplate/internal/api/repositories"
)

type (
	IEmployeeUsecase interface {
		Register(ctx context.Context, req requests.NewEmployeeRequest) error
	}
	employeeUsecase struct {
		repo repositories.IEmployeeRepository
	}
)

// Register implements IEmployeeUsecase.
func (u *employeeUsecase) Register(ctx context.Context, req requests.NewEmployeeRequest) error {

	entity := entities.NewEmployee(
		req.EmpName,
		"app",
	)
	if err := u.repo.Create(ctx, *entity); err != nil {
		return err
	}

	return nil
}

func newEmployeeUsecase(
	repo repositories.IEmployeeRepository,
) IEmployeeUsecase {
	return &employeeUsecase{
		repo: repo,
	}
}
