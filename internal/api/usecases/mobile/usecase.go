package api_mobile_usecases

import (
	"github.com/milfan/go-boilerplate/internal/api/repositories"
	"github.com/sirupsen/logrus"
)

type (
	MobileUsecases struct {
		OrderUsecase IOrderUsecase
	}
)

func RegisterMobileUsecases(
	repo repositories.Repositories,
	logger *logrus.Logger,
) MobileUsecases {
	return MobileUsecases{
		OrderUsecase: newOrderUsecase(
			logger,
			repo.OrderRepositories,
			repo.ProductRepositories,
		),
	}
}
