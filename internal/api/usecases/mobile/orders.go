package api_mobile_usecases

import (
	"context"
	"errors"

	"github.com/milfan/go-boilerplate/internal/api/entities"
	api_error "github.com/milfan/go-boilerplate/internal/api/errors"
	"github.com/milfan/go-boilerplate/internal/api/presenters/requests"
	"github.com/milfan/go-boilerplate/internal/api/repositories"
	pkg_errors "github.com/milfan/go-boilerplate/pkg/errors"
	"github.com/sirupsen/logrus"
)

type (
	IOrderUsecase interface {
		Add(ctx context.Context, req requests.NewOrderRequest) error
		List(ctx context.Context, req requests.OrderListRequest) ([]entities.Orders, *int64, error)
	}
	orderUsecase struct {
		logger      *logrus.Logger
		repo        repositories.IOrdersRepository
		repoProduct repositories.IProductRepository
	}
)

// List implements IOrderUsecase.
func (u *orderUsecase) List(ctx context.Context, req requests.OrderListRequest) ([]entities.Orders, *int64, error) {

	results, count, err := u.repo.List(ctx, req.Page, req.PerPage)
	if err != nil {
		return nil, nil, pkg_errors.New().Error(api_error.REPOSITORY_GET_ERROR, err)
	}

	return results, count, nil
}

// Add implements IOrderUsecase.
func (u *orderUsecase) Add(ctx context.Context, req requests.NewOrderRequest) error {

	productIds := make([]uint64, 0)
	orderDetails := make([]entities.OrderDetails, 0)

	for _, item := range req.OrderDetails {
		productIds = append(productIds, item.ProductID)
		orderDetail := entities.NewOrderDetail(
			item.ProductID,
			item.ProductQty,
			item.ProductPrice,
			"user",
		)
		orderDetails = append(orderDetails, *orderDetail)
	}
	order := entities.NewOrder(
		req.OrderDate,
		"user",
		orderDetails,
	)

	checkProducts, err := u.repoProduct.FindByIds(ctx, productIds)
	if err != nil {
		return pkg_errors.New().Error(api_error.REPOSITORY_GET_ERROR, err)
	}

	if len(orderDetails) != len(checkProducts) {
		return pkg_errors.New().Error(api_error.MISSMATCH_PRODUCT_ID, errors.New("miss match product"))
	}

	if err := u.repo.Add(ctx, *order); err != nil {
		return pkg_errors.New().Error(api_error.REPOSITORY_SAVE_ERROR, err)
	}

	return nil
}

func newOrderUsecase(
	logger *logrus.Logger,
	repo repositories.IOrdersRepository,
	repoProduct repositories.IProductRepository,
) IOrderUsecase {
	return &orderUsecase{
		logger:      logger,
		repo:        repo,
		repoProduct: repoProduct,
	}
}
