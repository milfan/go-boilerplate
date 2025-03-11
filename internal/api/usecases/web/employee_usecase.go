package api_web_usecases

import (
	"context"
	"fmt"
)

type (
	IEmployeeUsecase interface {
		Register(ctx context.Context) error
	}
	employeeUsecase struct{}
)

// Register implements IEmployeeUsecase.
func (e *employeeUsecase) Register(ctx context.Context) error {
	fmt.Println("OK")

	return nil
}

func newEmployeeUsecase() IEmployeeUsecase {
	return &employeeUsecase{}
}
