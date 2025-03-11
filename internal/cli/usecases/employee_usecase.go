package internal_cli_usecases

import (
	"context"
)

type (
	IEmployeeUsecase interface {
		FindEmployee(ctx context.Context) error
	}
	employeeUsecase struct{}
)

// FindEmployee implements IEmployeeUsecase.
func (e *employeeUsecase) FindEmployee(ctx context.Context) error {
	panic("unimplemented")
}

func newEmployeeUsecase() IEmployeeUsecase {
	return &employeeUsecase{}
}
