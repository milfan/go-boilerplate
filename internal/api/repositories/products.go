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
	IProductRepository interface {
		Add(ctx context.Context, entity entities.Products) error
		List(ctx context.Context, page, perPage int) ([]entities.Products, *int64, error)
		FindByIds(ctx context.Context, ids []uint64) ([]entities.Products, error)
	}
	productRepository struct {
		conn   config_postgres.Postgres
		logger *logrus.Logger
	}
)

// FindByIds implements IProductRepository.
func (r *productRepository) FindByIds(ctx context.Context, ids []uint64) ([]entities.Products, error) {
	var (
		models   []models.Product
		entities []entities.Products
	)

	gormSession := &gorm.Session{PrepareStmt: true, QueryFields: true}
	query := r.conn.Conn.Session(gormSession).WithContext(ctx)

	if ids != nil {
		query = query.Where("id IN (?)", ids)
	}

	if err := query.Find(&models).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}

	for _, model := range models {
		temp := model.Entity()
		entities = append(entities, *temp)
	}

	return entities, nil
}

// List implements IProductRepository.
func (r *productRepository) List(ctx context.Context, page int, perPage int) ([]entities.Products, *int64, error) {

	var (
		models   []models.Product
		entities []entities.Products
		count    int64
	)

	gormSession := &gorm.Session{PrepareStmt: true, QueryFields: true}
	query := r.conn.Conn.Session(gormSession).WithContext(ctx)

	err := query.Model(&models).Count(&count).Error
	if err != nil {
		return nil, nil, err
	}

	query = query.Scopes(api_helpers.Paginate(page, perPage))
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

// Add implements IProductRepository.
func (r *productRepository) Add(ctx context.Context, entity entities.Products) error {

	model := models.TransformProductModel(entity)

	query := r.conn.Conn.WithContext(ctx).Begin()

	if model.ID == 0 {
		var seq uint64
		if err := query.Raw(`SELECT nextval(pg_get_serial_sequence('products', 'id'));`).Scan(&seq).Error; err != nil {
			query.Rollback()
			return err
		}
		if seq == 0 {
			query.Rollback()
			return errors.New("error when get product sequence")
		}
		model.ProductCode = entity.GenerateProductCode(uint64(seq) + 1)
	}

	if err := query.Save(&model).Error; err != nil {
		query.Rollback()
		return err
	}
	query.Commit()

	return nil
}

func newProductsRepository(
	conn config_postgres.Postgres,
	logger *logrus.Logger,
) IProductRepository {
	return &productRepository{
		conn:   conn,
		logger: logger,
	}
}
