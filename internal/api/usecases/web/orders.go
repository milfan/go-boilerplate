package api_web_usecases

import (
	"context"

	"github.com/milfan/go-boilerplate/internal/api/entities"
	api_error "github.com/milfan/go-boilerplate/internal/api/errors"
	"github.com/milfan/go-boilerplate/internal/api/presenters/requests"
	"github.com/milfan/go-boilerplate/internal/api/repositories"
	pkg_errors "github.com/milfan/go-boilerplate/pkg/errors"
	"github.com/sirupsen/logrus"
)

type (
	IProductUsecase interface {
		Add(ctx context.Context, req requests.NewProductRequest) error
		List(ctx context.Context, req requests.ProductListRequest) ([]entities.Products, *int64, error)
	}
	productUsecase struct {
		logger *logrus.Logger
		repo   repositories.IProductRepository
	}
)

// List implements IProductUsecase.
func (u *productUsecase) List(ctx context.Context, req requests.ProductListRequest) ([]entities.Products, *int64, error) {

	results, count, err := u.repo.List(ctx, req.Page, req.PerPage)
	if err != nil {
		return nil, nil, pkg_errors.New().Error(api_error.REPOSITORY_GET_ERROR, err)
	}

	return results, count, nil
}

// Add implements IProductUsecase.
func (u *productUsecase) Add(ctx context.Context, req requests.NewProductRequest) error {

	entity := entities.NewProduct(
		req.ProductName,
		req.ProductPrice,
		"user",
	)
	if err := u.repo.Add(ctx, *entity); err != nil {
		return err
	}

	return nil
}

func newProductUsecase(
	logger *logrus.Logger,
	repo repositories.IProductRepository,
) IProductUsecase {
	return &productUsecase{
		logger: logger,
		repo:   repo,
	}
}
