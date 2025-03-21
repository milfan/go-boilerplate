package api_web_usecases

import (
	"context"

	"github.com/milfan/go-boilerplate/internal/api/entities"
	"github.com/milfan/go-boilerplate/internal/api/presenters/requests"
	"github.com/milfan/go-boilerplate/internal/api/repositories"
	"github.com/sirupsen/logrus"
)

type (
	IEmployeeUsecase interface {
		Register(ctx context.Context, req requests.NewEmployeeRequest) error
	}
	employeeUsecase struct {
		logger *logrus.Logger
		repo   repositories.IEmployeeRepository
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
	logger *logrus.Logger,
	repo repositories.IEmployeeRepository,
) IEmployeeUsecase {
	return &employeeUsecase{
		logger: logger,
		repo:   repo,
	}
}
