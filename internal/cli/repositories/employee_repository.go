package internal_cli_repositories

import (
	"context"

	"gorm.io/gorm"
)

type (
	IEmployeeRepository interface {
		FindEmployee(ctx context.Context) error
	}
	employeeRepository struct {
		conn *gorm.DB
	}
)

// FindEmployee implements IEmployeeRepository.
func (e *employeeRepository) FindEmployee(ctx context.Context) error {
	panic("unimplemented")
}

func newEmployeeRepository(
	conn *gorm.DB,
) IEmployeeRepository {
	return &employeeRepository{
		conn: conn,
	}
}
