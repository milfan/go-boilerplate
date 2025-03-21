package api_usecases

import (
	"github.com/milfan/go-boilerplate/internal/api/repositories"
	api_web_usecases "github.com/milfan/go-boilerplate/internal/api/usecases/web"
	"github.com/sirupsen/logrus"
)

type (
	Usecases struct {
		WebUsecases api_web_usecases.WebUsecases
	}
)

func LoadUsecases(
	repo repositories.Repositories,
	logger *logrus.Logger,
) Usecases {
	return Usecases{
		WebUsecases: api_web_usecases.RegisterWebUsecases(repo, logger),
	}
}
