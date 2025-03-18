package repositories

import (
	"context"
	"errors"

	config_postgres "github.com/milfan/go-boilerplate/configs/postgres"
	"github.com/milfan/go-boilerplate/internal/api/entities"
	"github.com/milfan/go-boilerplate/internal/api/models"
	pkg_log "github.com/milfan/go-boilerplate/pkg/log"
)

type (
	IEmployeeRepository interface {
		Create(ctx context.Context, entity entities.Employee) error
	}
	employeeRepository struct {
		conn      config_postgres.Postgres
		appLogger *pkg_log.AppLogger
	}
)

// Create implements IEmployeeRepository.
func (r *employeeRepository) Create(ctx context.Context, entity entities.Employee) error {

	model := models.ToEmployeeModel(entity)

	query := r.conn.Conn.WithContext(ctx).Begin()

	if model.ID == 0 {
		var seq uint64
		if err := query.Raw(`SELECT nextval(pg_get_serial_sequence('employees', 'id'));`).Scan(&seq).Error; err != nil {
			query.Rollback()
			return err
		}
		if seq == 0 {
			query.Rollback()
			return errors.New("error when get order_code_seq")
		}
		model.EmpCode = entities.GenerateEmpCode(uint64(seq) + 1)
	}

	if err := query.Save(&model).Error; err != nil {
		query.Rollback()
		return err
	}
	query.Commit()

	return nil
}

func newEmployeeRepository(
	conn config_postgres.Postgres,
	appLogger *pkg_log.AppLogger,
) IEmployeeRepository {
	return &employeeRepository{
		conn:      conn,
		appLogger: appLogger,
	}
}
