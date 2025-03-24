package repositories

import (
	"context"
	"errors"

	config_postgres "github.com/milfan/go-boilerplate/configs/postgres"
	"github.com/milfan/go-boilerplate/internal/api/entities"
	api_helpers "github.com/milfan/go-boilerplate/internal/api/helpers"
	"github.com/milfan/go-boilerplate/internal/api/models"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type (
	IOrdersRepository interface {
		Add(ctx context.Context, entity entities.Orders) error
		List(ctx context.Context, page, perPage int) ([]entities.Orders, *int64, error)
	}
	orderRepository struct {
		conn   config_postgres.Postgres
		logger *logrus.Logger
	}
)

// List implements IOrderRepository.
func (r *orderRepository) List(ctx context.Context, page int, perPage int) ([]entities.Orders, *int64, error) {

	var (
		models   []models.Order
		entities []entities.Orders
		count    int64
	)

	gormSession := &gorm.Session{PrepareStmt: true, QueryFields: true}
	query := r.conn.Conn.Session(gormSession).WithContext(ctx)

	err := query.Model(&models).Count(&count).Error
	if err != nil {
		return nil, nil, err
	}

	query = query.Scopes(api_helpers.Paginate(page, perPage))
	query = query.Preload("OrderDetails.Product")

	if err := query.Order("created_at DESC").Find(&models).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil, nil
		}
		return nil, nil, err
	}

	for _, model := range models {
		temp := model.Entity()
		entities = append(entities, *temp)
	}

	return entities, &count, nil
}

// Add implements IOrderRepository.
func (r *orderRepository) Add(ctx context.Context, entity entities.Orders) error {

	model := models.TransformOrderModel(entity)

	gormSession := &gorm.Session{PrepareStmt: true, FullSaveAssociations: true}
	query := r.conn.Conn.Session(gormSession).WithContext(ctx).Begin()

	if model.ID == 0 {
		var seq uint64
		if err := query.Raw(`SELECT nextval(pg_get_serial_sequence('orders', 'id'));`).Scan(&seq).Error; err != nil {
			query.Rollback()
			return err
		}
		if seq == 0 {
			query.Rollback()
			return errors.New("error when get order sequence")
		}
		model.OrderCode = entity.GenerateOrderCode(uint64(seq) + 1)
	}

	if err := query.Save(&model).Error; err != nil {
		query.Rollback()
		return err
	}
	query.Commit()

	return nil
}

func newOrdersRepository(
	conn config_postgres.Postgres,
	logger *logrus.Logger,
) IOrdersRepository {
	return &orderRepository{
		conn:   conn,
		logger: logger,
	}
}
