package api_usecases

import (
	"github.com/milfan/go-boilerplate/internal/api/repositories"
	api_mobile_usecases "github.com/milfan/go-boilerplate/internal/api/usecases/mobile"
	api_web_usecases "github.com/milfan/go-boilerplate/internal/api/usecases/web"
	"github.com/sirupsen/logrus"
)

type (
	Usecases struct {
		WebUsecases    api_web_usecases.WebUsecases
		MobileUsecases api_mobile_usecases.MobileUsecases
	}
)

func LoadUsecases(
	repo repositories.Repositories,
	logger *logrus.Logger,
) Usecases {
	return Usecases{
		WebUsecases:    api_web_usecases.RegisterWebUsecases(repo, logger),
		MobileUsecases: api_mobile_usecases.RegisterMobileUsecases(repo, logger),
	}
}
