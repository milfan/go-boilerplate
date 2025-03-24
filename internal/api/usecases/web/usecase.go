package api_web_usecases

import (
	"github.com/milfan/go-boilerplate/internal/api/repositories"
	"github.com/sirupsen/logrus"
)

type (
	WebUsecases struct {
		ProductUsecases IProductUsecase
	}
)

func RegisterWebUsecases(
	repo repositories.Repositories,
	logger *logrus.Logger,
) WebUsecases {
	return WebUsecases{
		ProductUsecases: newProductUsecase(logger, repo.ProductRepositories),
	}
}
